package board

import (
	"fmt"
)

var KNIGHT_DELTAS = [8]int{-21, -19, -12, -8, 12, 21, 19, 8}
var BISHOP_DELTAS = [4]int{-11, -9, 9, 11}
var ROOK_DELTAS = [4]int{-10, -1, 1, 10}
var QUEEN_DELTAS = [8]int{-11, -10, -9, -1, 1, 9, 10, 11}
var WHITE_PAWN_ATTACKS = [2]int{-11, -9}
var BLACK_PAWN_ATTACKS = [2]int{11, 9}
var WHITE_PAWN_DELTAS = [2]int{-10, -20}
var BLACK_PAWN_DELTAS = [2]int{10, 20}
var WHITE_PROMOTION_PIECES = [4]int{WHITE_QUEEN, WHITE_ROOK, WHITE_BISHOP, WHITE_KNIGHT}
var BLACK_PROMOTION_PIECES = [4]int{BLACK_QUEEN, BLACK_ROOK, BLACK_BISHOP, BLACK_KNIGHT}

// let's just get on the board
// easy to get trapped in analysis paralysis
// this is obviously suboptimal but we go again
// should go without saying that these methods can be consolidated

type Move struct {
	from            int
	to              int
	capture         bool
	castleKingside  bool
	castleQueenside bool
	promote         bool
	promotionPiece  int
	doublePawnPush  bool
}

func equalMoves(m1, m2 Move) bool {
	if m1.from != m2.from {
		return false
	}
	if m1.to != m2.to {
		return false
	}
	if m1.capture != m2.capture {
		return false
	}
	if m1.castleKingside != m2.castleKingside {
		return false
	}
	if m1.castleQueenside != m2.castleQueenside {
		return false
	}
	if m1.promote != m2.promote {
		return false
	}
	if m1.promotionPiece != m2.promotionPiece {
		return false
	}
	if m1.doublePawnPush != m2.doublePawnPush {
		return false
	}
	return true
}

// these need tests

func (b Board) LegalMoves() []Move {
	moves := make([]Move, 0)
	for _, m := range b.Moves() {
		r := b.MakeMove(m)
		if r == true {
			b.UnmakeMove()
			moves = append(moves, m)
		}
	}
	return moves
}

func (b Board) Moves() []Move {
	m := make([]Move, 0)
	if b.side == WHITE {
		for i := WHITE_PAWN; i <= WHITE_KING; i++ {
			if b.pieceSquares[i] != nil {
				for j := 0; j < len(b.pieceSquares[i]); j++ {
					moves, err := b.MovesForPiece(b.pieceSquares[i][j])
					if err != nil {
						panic(err)
					}
					m = append(m, moves...)
				}
			}
		}
	} else {
		for i := BLACK_PAWN; i <= BLACK_KING; i++ {
			if b.pieceSquares[i] != nil {
				for j := 0; j < len(b.pieceSquares[i]); j++ {
					moves, err := b.MovesForPiece(b.pieceSquares[i][j])
					if err != nil {
						panic(err)
					}
					m = append(m, moves...)
				}
			}
		}
	}
	return m
}

func (b Board) MovesForPiece(sq int) ([]Move, error) {
	piece := b.PieceAt(sq)
	switch piece {
	case WHITE_PAWN:
		return b.WhitePawnMoves(sq), nil
	case WHITE_KNIGHT:
		return b.WhiteKnightMoves(sq), nil
	case WHITE_BISHOP:
		return b.WhiteBishopMoves(sq), nil
	case WHITE_ROOK:
		return b.WhiteRookMoves(sq), nil
	case WHITE_QUEEN:
		return b.WhiteQueenMoves(sq), nil
	case WHITE_KING:
		return b.WhiteKingMoves(sq), nil
	case BLACK_PAWN:
		return b.BlackPawnMoves(sq), nil
	case BLACK_KNIGHT:
		return b.BlackKnightMoves(sq), nil
	case BLACK_BISHOP:
		return b.BlackBishopMoves(sq), nil
	case BLACK_ROOK:
		return b.BlackRookMoves(sq), nil
	case BLACK_QUEEN:
		return b.BlackQueenMoves(sq), nil
	case BLACK_KING:
		return b.BlackKingMoves(sq), nil
	default:
		return nil, fmt.Errorf("invalid square: %d", sq)
	}
}

func (b Board) WhitePawnMoves(sq int) []Move {
	moves := make([]Move, 14, 14)
	mi := 0
	for _, d := range WHITE_PAWN_ATTACKS {
		ns := sq + d
		p := b.PieceAt(ns)
		if 6 < p && p < 12 { // black piece
			if ns <= IH8 { // check promotion
				for _, piece := range WHITE_PROMOTION_PIECES {
					moves[mi] = Move{sq, ns, true, false, false, true, piece, false}
					mi += 1
				}
			} else {
				moves[mi] = Move{sq, ns, true, false, false, false, WHITE_PAWN, false}
				mi += 1
			}
		} else if b.ep != nil && *b.ep == ns {
			moves[mi] = Move{sq, ns, true, false, false, false, WHITE_PAWN, false}
			mi += 1
		}
	} // one sq push
	if b.PieceAt(sq+WHITE_PAWN_DELTAS[0]) == EMPTY_SQUARE {
		ns := sq + WHITE_PAWN_DELTAS[0]
		if ns <= IH8 { // check promotion
			for _, piece := range WHITE_PROMOTION_PIECES {
				moves[mi] = Move{sq, ns, false, false, false, true, piece, false}
				mi += 1
			}
		} else {
			moves[mi] = Move{sq, ns, false, false, false, false, WHITE_PAWN, false}
			mi += 1
			if IA2 <= sq && sq <= IH2 && b.PieceAt(sq+WHITE_PAWN_DELTAS[1]) == EMPTY_SQUARE {
				moves[mi] = Move{sq, sq + WHITE_PAWN_DELTAS[1], false, false, false, false, WHITE_PAWN, true}
				mi += 1
			}
		}
	} // opening two sq push
	return moves[:mi]
}

func (b Board) BlackPawnMoves(sq int) []Move {
	moves := make([]Move, 14, 14)
	mi := 0
	for _, d := range BLACK_PAWN_ATTACKS {
		ns := sq + d
		p := b.PieceAt(ns)
		if 0 < p && p < 6 { // white piece
			if ns >= IA1 { // check promotion
				for _, piece := range BLACK_PROMOTION_PIECES {
					moves[mi] = Move{sq, ns, true, false, false, true, piece, false}
					mi += 1
				}
			} else {
				moves[mi] = Move{sq, ns, true, false, false, false, BLACK_PAWN, false}
				mi += 1
			}
		} else if b.ep != nil && *b.ep == ns {
			moves[mi] = Move{sq, ns, true, false, false, false, BLACK_PAWN, false}
			mi += 1
		}
	} // one sq push
	if b.PieceAt(sq+BLACK_PAWN_DELTAS[0]) == EMPTY_SQUARE {
		ns := sq + BLACK_PAWN_DELTAS[0]
		if ns >= IA1 { // check promotion
			for _, piece := range BLACK_PROMOTION_PIECES {
				moves[mi] = Move{sq, ns, false, false, false, true, piece, false}
				mi += 1
			}
		} else {
			moves[mi] = Move{sq, ns, false, false, false, false, BLACK_PAWN, false}
			mi += 1
			if IA7 <= sq && sq <= IH7 && b.PieceAt(sq+BLACK_PAWN_DELTAS[1]) == EMPTY_SQUARE {
				moves[mi] = Move{sq, sq + BLACK_PAWN_DELTAS[1], false, false, false, false, BLACK_PAWN, true}
				mi += 1
			}
		}
	} // opening two sq push
	return moves[:mi]
}

func (b Board) WhiteKnightMoves(sq int) []Move {
	moves := make([]Move, 8, 8)
	mi := 0
	for _, d := range KNIGHT_DELTAS {
		ns := d + sq
		p := b.PieceAt(ns)
		if 6 < p && p < 12 {
			moves[mi] = Move{sq, ns, true, false, false, false, WHITE_KNIGHT, false}
			mi += 1
		} else if p == EMPTY_SQUARE {
			moves[mi] = Move{sq, ns, false, false, false, false, WHITE_KNIGHT, false}
			mi += 1
		}
	}

	return moves[:mi]
}

func (b Board) BlackKnightMoves(sq int) []Move {
	moves := make([]Move, 8, 8)
	mi := 0
	for _, d := range KNIGHT_DELTAS {
		ns := d + sq
		p := b.PieceAt(ns)
		if 0 < p && p < 6 {
			moves[mi] = Move{sq, ns, true, false, false, false, BLACK_KNIGHT, false}
			mi += 1
		} else if p == EMPTY_SQUARE {
			moves[mi] = Move{sq, ns, false, false, false, false, BLACK_KNIGHT, false}
			mi += 1
		}
	}

	return moves[:mi]
}

func (b Board) WhiteBishopMoves(sq int) []Move {
	moves := make([]Move, 13, 13)
	mi := 0
	for _, d := range BISHOP_DELTAS {
		ns := d + sq
		for b.PieceAt(ns) == EMPTY_SQUARE {
			moves[mi] = Move{sq, ns, false, false, false, false, WHITE_BISHOP, false}
			mi += 1
			ns += d
		}
		p := b.PieceAt(ns)
		if 6 < p && p < 12 {
			moves[mi] = Move{sq, ns, true, false, false, false, WHITE_BISHOP, false}
			mi += 1
		}
	}
	return moves[:mi]
}

func (b Board) BlackBishopMoves(sq int) []Move {
	moves := make([]Move, 13, 13)
	mi := 0
	for _, d := range BISHOP_DELTAS {
		ns := d + sq
		for b.PieceAt(ns) == EMPTY_SQUARE {
			moves[mi] = Move{sq, ns, false, false, false, false, BLACK_BISHOP, false}
			mi += 1
			ns += d
		}
		p := b.PieceAt(ns)
		if 0 < p && p < 6 {
			moves[mi] = Move{sq, ns, true, false, false, false, BLACK_BISHOP, false}
			mi += 1
		}
	}
	return moves[:mi]
}

func (b Board) WhiteRookMoves(sq int) []Move {
	moves := make([]Move, 14, 14)
	mi := 0
	for _, d := range ROOK_DELTAS {
		ns := d + sq
		for b.PieceAt(ns) == EMPTY_SQUARE {
			moves[mi] = Move{sq, ns, false, false, false, false, WHITE_ROOK, false}
			mi += 1
			ns += d
		}
		p := b.PieceAt(ns)
		if 6 < p && p < 12 {
			moves[mi] = Move{sq, ns, true, false, false, false, WHITE_ROOK, false}
			mi += 1
		}
	}
	return moves[:mi]
}

func (b Board) BlackRookMoves(sq int) []Move {
	moves := make([]Move, 14, 14)
	mi := 0
	for _, d := range ROOK_DELTAS {
		ns := d + sq
		for b.PieceAt(ns) == EMPTY_SQUARE {
			moves[mi] = Move{sq, ns, false, false, false, false, BLACK_ROOK, false}
			mi += 1
			ns += d
		}
		p := b.PieceAt(ns)
		if 0 < p && p < 6 {
			moves[mi] = Move{sq, ns, true, false, false, false, BLACK_ROOK, false}
			mi += 1
		}
	}
	return moves[:mi]
}

func (b Board) WhiteQueenMoves(sq int) []Move {
	moves := append(b.WhiteBishopMoves(sq), b.WhiteRookMoves(sq)...)
	for i, _ := range moves {
		moves[i].promotionPiece = WHITE_QUEEN // might as well lean into this silly convention
	}
	return moves
}

func (b Board) BlackQueenMoves(sq int) []Move {
	moves := append(b.BlackBishopMoves(sq), b.BlackRookMoves(sq)...)
	for i, _ := range moves {
		moves[i].promotionPiece = BLACK_QUEEN
	}
	return moves
}

func (b Board) WhiteKingMoves(sq int) []Move {
	attackedSquares := toLookupMap(b.SquaresAttackedByBlackPieces())
	moves := make([]Move, 10, 10)
	mi := 0
	for _, d := range QUEEN_DELTAS {
		if attackedSquares[sq+d] {
			continue
		}

		p := b.PieceAt(sq + d)
		if p == EMPTY_SQUARE {
			moves[mi] = Move{sq, sq + d, false, false, false, false, WHITE_KING, false}
			mi += 1
		} else if 6 < p && p < 12 {
			moves[mi] = Move{sq, sq + d, true, false, false, false, WHITE_KING, false}
			mi += 1
		}
	}
	if b.checkWhiteCastleKingside(attackedSquares) {
		moves[mi] = Move{sq, sq + 2, false, true, false, false, WHITE_KING, false}
		mi += 1
	}
	if b.checkWhiteCastleQueenside(attackedSquares) {
		moves[mi] = Move{sq, sq - 2, false, false, true, false, WHITE_KING, false}
		mi += 1
	}
	return moves[:mi]
}

func (b Board) checkWhiteCastleKingside(attackedSquares map[int]bool) bool {
	if !b.castle[0] {
		return false
	}

	if b.PieceAt(IF1) != EMPTY_SQUARE || b.PieceAt(IG1) != EMPTY_SQUARE {
		return false
	}

	for i := IE1; i < IH1; i++ {
		if attackedSquares[i] == true {
			return false
		}
	}

	return true
}

func (b Board) checkWhiteCastleQueenside(attackedSquares map[int]bool) bool {
	if !b.castle[1] {
		return false
	}

	for i := IB1; i < IE1; i++ {
		if b.PieceAt(i) != EMPTY_SQUARE {
			return false
		}
	}

	for i := IC1; i < IF1; i++ {
		if attackedSquares[i] == true {
			return false
		}
	}

	return true
}

func (b Board) BlackKingMoves(sq int) []Move {
	attackedSquares := toLookupMap(b.SquaresAttackedByWhitePieces())
	moves := make([]Move, 10, 10)
	mi := 0
	for _, d := range QUEEN_DELTAS {
		if attackedSquares[sq+d] {
			continue
		}

		p := b.PieceAt(sq + d)
		if p == EMPTY_SQUARE {
			moves[mi] = Move{sq, sq + d, false, false, false, false, BLACK_KING, false}
			mi += 1
		} else if 0 < p && p < 6 {
			moves[mi] = Move{sq, sq + d, true, false, false, false, BLACK_KING, false}
			mi += 1
		}
	}
	if b.checkBlackCastleKingside(attackedSquares) {
		moves[mi] = Move{sq, sq + 2, false, true, false, false, BLACK_KING, false}
		mi += 1
	}
	if b.checkBlackCastleQueenside(attackedSquares) {
		moves[mi] = Move{sq, sq - 2, false, false, true, false, BLACK_KING, false}
		mi += 1
	}
	return moves[:mi]
}

func (b Board) checkBlackCastleKingside(attackedSquares map[int]bool) bool {
	if !b.castle[2] {
		return false
	}

	if b.PieceAt(IF8) != EMPTY_SQUARE || b.PieceAt(IG8) != EMPTY_SQUARE {
		return false
	}

	for i := IE8; i < IH8; i++ {
		if attackedSquares[i] == true {
			return false
		}
	}

	return true
}

func (b Board) checkBlackCastleQueenside(attackedSquares map[int]bool) bool {
	if !b.castle[3] {
		return false
	}

	for i := IB8; i < ID8; i++ {
		if b.PieceAt(i) != EMPTY_SQUARE {
			return false
		}
	}

	for i := IC8; i < IF8; i++ {
		if attackedSquares[i] == true {
			return false
		}
	}

	return true
}

func (b Board) SquaresAttackedByWhitePieces() []int {
	attackedSquares := make([]int, 0)
	for i := WHITE_PAWN; i <= WHITE_KING; i++ {
		sqs, ok := b.pieceSquares[i]
		if ok {
			for _, sq := range sqs {
				switch i {
				case WHITE_PAWN:
					attackedSquares = append(attackedSquares, b.PawnAttacks(sq, WHITE_PAWN_ATTACKS)...)
				case WHITE_KNIGHT:
					attackedSquares = append(attackedSquares, b.KnightAttacks(sq)...)
				case WHITE_BISHOP:
					attackedSquares = append(attackedSquares, b.BishopAttacks(sq)...)
				case WHITE_ROOK:
					attackedSquares = append(attackedSquares, b.RookAttacks(sq)...)
				case WHITE_QUEEN:
					attackedSquares = append(attackedSquares, b.QueenAttacks(sq)...)
				case WHITE_KING:
					attackedSquares = append(attackedSquares, b.KingAttacks(sq)...)
				}
			}
		}
	}

	return attackedSquares
}

func (b Board) SquaresAttackedByBlackPieces() []int {
	attackedSquares := make([]int, 0)
	for i := BLACK_PAWN; i <= BLACK_KING; i++ {
		sqs, ok := b.pieceSquares[i]
		if ok {
			for _, sq := range sqs {
				switch i {
				case BLACK_PAWN:
					attackedSquares = append(attackedSquares, b.PawnAttacks(sq, BLACK_PAWN_ATTACKS)...)
				case BLACK_KNIGHT:
					attackedSquares = append(attackedSquares, b.KnightAttacks(sq)...)
				case BLACK_BISHOP:
					attackedSquares = append(attackedSquares, b.BishopAttacks(sq)...)
				case BLACK_ROOK:
					attackedSquares = append(attackedSquares, b.RookAttacks(sq)...)
				case BLACK_QUEEN:
					attackedSquares = append(attackedSquares, b.QueenAttacks(sq)...)
				case BLACK_KING:
					attackedSquares = append(attackedSquares, b.KingAttacks(sq)...)
				}
			}
		}
	}

	return attackedSquares
}

func (b Board) PawnAttacks(sq int, deltas [2]int) []int {
	attacks := make([]int, 2, 2)
	mi := 0
	for _, d := range deltas {
		ns := d + sq
		if b.PieceAt(ns) != OFF_BOARD {
			attacks[mi] = ns
			mi += 1
		}
	}
	return attacks[:mi]
}

func (b Board) KnightAttacks(sq int) []int {
	attacks := make([]int, 8, 8)
	mi := 0
	for _, d := range KNIGHT_DELTAS {
		ns := d + sq
		if b.PieceAt(ns) != OFF_BOARD {
			attacks[mi] = ns
			mi += 1
		}
	}
	return attacks[:mi]
}

func (b Board) BishopAttacks(sq int) []int {
	attacks := make([]int, 13, 13)
	mi := 0
	for _, d := range BISHOP_DELTAS {
		ns := d + sq
		for b.PieceAt(ns) == EMPTY_SQUARE {
			attacks[mi] = ns
			mi += 1
			ns += d
		}
		if b.PieceAt(ns) != OFF_BOARD {
			attacks[mi] = ns
			mi += 1
		}
	}
	return attacks[:mi]
}

func (b Board) RookAttacks(sq int) []int {
	attacks := make([]int, 14, 14)
	mi := 0
	for _, d := range ROOK_DELTAS {
		ns := d + sq
		for b.PieceAt(ns) == EMPTY_SQUARE {
			attacks[mi] = ns
			mi += 1
			ns += d
		}
		if b.PieceAt(ns) != OFF_BOARD {
			attacks[mi] = ns
			mi += 1
		}
	}
	return attacks[:mi]
}

func (b Board) QueenAttacks(sq int) []int {
	return append(b.BishopAttacks(sq), b.RookAttacks(sq)...)
}

func (b Board) KingAttacks(sq int) []int {
	attacks := make([]int, 8, 8)
	mi := 0
	for _, d := range QUEEN_DELTAS {
		ns := sq + d
		if b.PieceAt(ns) != OFF_BOARD {
			attacks[mi] = ns
			mi += 1
		}
	}
	return attacks[:mi]
}

// TODO: generic
func toLookupMap(sqs []int) map[int]bool {
	m := make(map[int]bool)
	for _, s := range sqs {
		m[s] = true
	}
	return m
}
