package board

import (
	"fmt"
	"sort"
)

var PIECE_COLORS = map[int]map[int]bool{
	WHITE: map[int]bool{
		WHITE_PAWN:   true,
		WHITE_KNIGHT: true,
		WHITE_BISHOP: true,
		WHITE_ROOK:   true,
		WHITE_QUEEN:  true,
		WHITE_KING:   true,
	},
	BLACK: map[int]bool{
		BLACK_PAWN:   true,
		BLACK_KNIGHT: true,
		BLACK_BISHOP: true,
		BLACK_ROOK:   true,
		BLACK_QUEEN:  true,
		BLACK_KING:   true,
	},
}

var PAWNS = map[int]bool{
	WHITE_PAWN: true,
	BLACK_PAWN: true,
}

var SLIDING_PIECES = map[int]bool{
	WHITE_KNIGHT: false,
	WHITE_BISHOP: true,
	WHITE_ROOK:   true,
	WHITE_QUEEN:  true,
	WHITE_KING:   false,
	BLACK_KNIGHT: false,
	BLACK_BISHOP: true,
	BLACK_ROOK:   true,
	BLACK_QUEEN:  true,
	BLACK_KING:   false,
}

var OFFSETS = map[int][]int{
	WHITE_KNIGHT: []int{-21, -19, -12, -8, 12, 21, 19, 8},
	BLACK_KNIGHT: []int{-21, -19, -12, -8, 12, 21, 19, 8},
	WHITE_BISHOP: []int{-11, -9, 9, 11},
	BLACK_BISHOP: []int{-11, -9, 9, 11},
	WHITE_ROOK:   []int{-10, -1, 1, 10},
	BLACK_ROOK:   []int{-10, -1, 1, 10},
	WHITE_QUEEN:  []int{-11, -10, -9, -1, 1, 9, 10, 11},
	BLACK_QUEEN:  []int{-11, -10, -9, -1, 1, 9, 10, 11},
	WHITE_KING:   []int{-11, -10, -9, -1, 1, 9, 10, 11},
	BLACK_KING:   []int{-11, -10, -9, -1, 1, 9, 10, 11},
}

var WHITE_PAWN_ATTACKS = [2]int{-11, -9}
var BLACK_PAWN_ATTACKS = [2]int{11, 9}
var WHITE_PAWN_DELTAS = [2]int{-10, -20}
var BLACK_PAWN_DELTAS = [2]int{10, 20}
var KNIGHT_DELTAS = [8]int{-21, -19, -12, -8, 12, 21, 19, 8}
var BISHOP_DELTAS = [4]int{-11, -9, 9, 11}
var ROOK_DELTAS = [4]int{-10, -1, 1, 10}
var QUEEN_DELTAS = [8]int{-11, -10, -9, -1, 1, 9, 10, 11}
var WHITE_PROMOTION_PIECES = [4]int{WHITE_QUEEN, WHITE_ROOK, WHITE_BISHOP, WHITE_KNIGHT}
var BLACK_PROMOTION_PIECES = [4]int{BLACK_QUEEN, BLACK_ROOK, BLACK_BISHOP, BLACK_KNIGHT}

type Move struct {
	From            int
	To              int
	Capture         bool
	CastleKingside  bool
	CastleQueenside bool
	Promote         bool
	PromotionPiece  int
	DoublePawnPush  bool
}

func (m Move) IsNull() bool {
	if m.From == 0 && m.To == 0 {
		return true
	}
	return false
}

func (m Move) Score(b Board) int {
	s := 0
	if m.Promote {
		s += 100
	}
	if m.Capture {
		s += 100
		mp := b.PieceAt(m.From)
		cp := b.PieceAt(m.To)
		s += CAPTURE_SCORE[mp][cp] * 100
	}
	if m.CastleKingside || m.CastleQueenside {
		s += 100
	}
	if m.DoublePawnPush && 2 < file(m.To) && file(m.To) < 6 {
		s += 100
	}
	return s
}

func (m Move) Print() {
	fmt.Printf("%s", m.ToString())
}

func (m Move) ToString() string {
	return fmt.Sprintf("%s - %s\n", SQ_NUM_TO_NAME[m.From], SQ_NUM_TO_NAME[m.To])
}

func EqualMoves(m1, m2 Move) bool {
	if m1.From != m2.From {
		return false
	}
	if m1.To != m2.To {
		return false
	}
	if m1.Capture != m2.Capture {
		return false
	}
	if m1.CastleKingside != m2.CastleKingside {
		return false
	}
	if m1.CastleQueenside != m2.CastleQueenside {
		return false
	}
	if m1.Promote != m2.Promote {
		return false
	}
	if m1.PromotionPiece != m2.PromotionPiece {
		return false
	}
	if m1.DoublePawnPush != m2.DoublePawnPush {
		return false
	}
	return true
}

func (b Board) LegalMoves() []Move {
	if b.legalMoves != nil {
		return b.legalMoves
	}
	moves := make([]Move, 0)
	// for _, m := range b.Moves() {
	for _, m := range b.genMoves() {
		r := b.MakeMove(m)
		if r == true {
			b.UnmakeMove()
			moves = append(moves, m)
		}
	}

	sort.Slice(moves, func(i, j int) bool {
		return moves[i].Score(b) > moves[j].Score(b)
	})

	b.legalMoves = moves
	return b.legalMoves
}

func (b Board) genMoves() []Move {
	m := make([]Move, 0)
	for i := 0; i < 64; i++ {
		sq := MAILBOX_64[i]
		p := b.PieceAt(sq)
		if _, ok := PIECE_COLORS[b.Side][p]; !ok {
			continue
		}
		if _, ok := PAWNS[p]; ok {
			m = append(m, b.genPawnMoves(sq)...)
		} else {
			m = append(m, b.genOffsetMoves(p, sq)...)
		}
	}
	return m
}

func (b Board) genPawnMoves(sq int) []Move {
	moves := make([]Move, 0)
	if b.Side == WHITE {
		for _, d := range WHITE_PAWN_ATTACKS {
			ns := sq + d
			p := b.PieceAt(ns)
			if _, ok := PIECE_COLORS[b.Side^1][p]; ok {
				if ns <= IH8 { // check promotion
					for _, piece := range WHITE_PROMOTION_PIECES {
						moves = append(moves, Move{sq, ns, true, false, false, true, piece, false})
					}
				} else {
					moves = append(moves, Move{sq, ns, true, false, false, false, WHITE_PAWN, false})
				}
			} else if b.Ep != nil && *b.Ep == ns {
				moves = append(moves, Move{sq, ns, true, false, false, false, WHITE_PAWN, false})
			}
		} // one sq push
		if b.PieceAt(sq+WHITE_PAWN_DELTAS[0]) == EMPTY_SQUARE {
			ns := sq + WHITE_PAWN_DELTAS[0]
			if ns <= IH8 { // check promotion
				for _, piece := range WHITE_PROMOTION_PIECES {
					moves = append(moves, Move{sq, ns, false, false, false, true, piece, false})
				}
			} else {
				moves = append(moves, Move{sq, ns, false, false, false, false, WHITE_PAWN, false})
				if IA2 <= sq && sq <= IH2 && b.PieceAt(sq+WHITE_PAWN_DELTAS[1]) == EMPTY_SQUARE {
					moves = append(moves, Move{sq, sq + WHITE_PAWN_DELTAS[1], false, false, false, false, WHITE_PAWN, true})
				} // opening two sq push
			}
		}
	} else {
		for _, d := range BLACK_PAWN_ATTACKS {
			ns := sq + d
			p := b.PieceAt(ns)
			if _, ok := PIECE_COLORS[b.Side^1][p]; ok {
				if ns >= IA1 { // check promotion
					for _, piece := range BLACK_PROMOTION_PIECES {
						moves = append(moves, Move{sq, ns, true, false, false, true, piece, false})
					}
				} else {
					moves = append(moves, Move{sq, ns, true, false, false, false, BLACK_PAWN, false})
				}
			} else if b.Ep != nil && *b.Ep == ns {
				moves = append(moves, Move{sq, ns, true, false, false, false, BLACK_PAWN, false})
			}
		} // one sq push
		if b.PieceAt(sq+BLACK_PAWN_DELTAS[0]) == EMPTY_SQUARE {
			ns := sq + BLACK_PAWN_DELTAS[0]
			if ns >= IA1 { // check promotion
				for _, piece := range BLACK_PROMOTION_PIECES {
					moves = append(moves, Move{sq, ns, false, false, false, true, piece, false})
				}
			} else {
				moves = append(moves, Move{sq, ns, false, false, false, false, BLACK_PAWN, false})
				if IA7 <= sq && sq <= IH7 && b.PieceAt(sq+BLACK_PAWN_DELTAS[1]) == EMPTY_SQUARE {
					moves = append(moves, Move{sq, sq + BLACK_PAWN_DELTAS[1], false, false, false, false, BLACK_PAWN, true})
				}
			} // opening two sq push
		}
	}
	return moves
}

func (b Board) genOffsetMoves(p, sq int) []Move {
	if p == WHITE_KING || p == BLACK_KING {
		return b.genKingMoves(p, sq)
	}

	moves := make([]Move, 0)

	for _, d := range OFFSETS[p] {
		nsq := sq + d
		if v, _ := SLIDING_PIECES[p]; v {
			for b.PieceAt(nsq) == EMPTY_SQUARE {
				moves = append(moves, Move{sq, nsq, false, false, false, false, p, false})
				nsq += d
			}
		}
		mp := b.PieceAt(nsq)
		if mp == EMPTY_SQUARE {
			moves = append(moves, Move{sq, nsq, false, false, false, false, p, false})
		}
		if _, ok := PIECE_COLORS[b.Side^1][mp]; ok {
			moves = append(moves, Move{sq, nsq, true, false, false, false, p, false})
		}
	}

	return moves
}

func (b Board) genKingMoves(p, sq int) []Move {
	moves := make([]Move, 0)

	for _, d := range OFFSETS[p] {
		if b.isAttacked(sq+d, b.Side) {
			continue
		}
		mp := b.PieceAt(sq + d)
		if mp == EMPTY_SQUARE {
			moves = append(moves, Move{sq, sq + d, false, false, false, false, p, false})
		}
		if _, ok := PIECE_COLORS[b.Side^1][mp]; ok {
			moves = append(moves, Move{sq, sq + d, true, false, false, false, p, false})
		}
	}
	kc := b.genKingsideCastle(p, sq)
	qc := b.genQueensideCastle(p, sq)
	if kc != nil {
		moves = append(moves, *kc)
	}
	if qc != nil {
		moves = append(moves, *qc)
	}
	return moves
}

func (b Board) genKingsideCastle(p, sq int) *Move {
	if p == WHITE_KING {
		if !b.Castle[0] {
			return nil
		}
		if b.PieceAt(IF1) != EMPTY_SQUARE || b.PieceAt(IG1) != EMPTY_SQUARE {
			return nil
		}

		for i := IE1; i < IH1; i++ {
			if b.isAttacked(i, WHITE) {
				return nil
			}
		}

		return &Move{sq, IG1, false, true, false, false, p, false}
	} else {
		if !b.Castle[2] {
			return nil
		}
		if b.PieceAt(IF8) != EMPTY_SQUARE || b.PieceAt(IG8) != EMPTY_SQUARE {
			return nil
		}

		for i := IE8; i < IH8; i++ {
			if b.isAttacked(i, BLACK) {
				return nil
			}
		}

		return &Move{sq, IG8, false, true, false, false, p, false}
	}
}

func (b Board) genQueensideCastle(p, sq int) *Move {
	if p == WHITE_KING {
		if !b.Castle[1] {
			return nil
		}

		for i := IB1; i < IE1; i++ {
			if b.PieceAt(i) != EMPTY_SQUARE {
				return nil
			}
		}

		for i := IC1; i < IF1; i++ {
			if b.isAttacked(i, WHITE) {
				return nil
			}
		}

		return &Move{sq, IC1, false, false, true, false, p, false}
	} else {
		if !b.Castle[3] {
			return nil
		}

		for i := IB8; i < IE8; i++ {
			if b.PieceAt(i) != EMPTY_SQUARE {
				return nil
			}
		}

		for i := IC8; i < IF8; i++ {
			if b.isAttacked(i, BLACK) {
				return nil
			}
		}

		return &Move{sq, IC8, false, false, true, false, p, false}
	}
}

// side indicates who we're checking is / isn't attacked
func (b Board) isAttacked(sq, side int) bool {
	return b.isPawnAttacked(sq, side) || b.isJumpAttacked(sq, side) || b.isCardinalAttacked(sq, side) || b.isOrdinalAttacked(sq, side) || b.isKingAttacked(sq, side)
}

func (b Board) isPawnAttacked(sq, side int) bool {
	if b.PieceAt(sq) == OFF_BOARD {
		return false
	}
	if side == WHITE {
		if b.PieceAt(sq-9) == BLACK_PAWN || b.PieceAt(sq-11) == BLACK_PAWN {
			return true
		}
	} else {
		if b.PieceAt(sq+9) == WHITE_PAWN || b.PieceAt(sq+11) == WHITE_PAWN {
			return true
		}
	}
	return false
}

func (b Board) isJumpAttacked(sq, side int) bool {
	if b.PieceAt(sq) == OFF_BOARD {
		return false
	}
	enemyPiece := WHITE_KNIGHT
	if side == WHITE {
		enemyPiece = BLACK_KNIGHT
	}
	for _, d := range KNIGHT_DELTAS {
		if b.PieceAt(sq+d) == enemyPiece {
			return true
		}
	}
	return false
}

var CARDINAL_ATTACKERS = map[int]map[int]bool{
	WHITE: map[int]bool{
		BLACK_ROOK:  true,
		BLACK_QUEEN: true,
	},
	BLACK: map[int]bool{
		WHITE_ROOK:  true,
		WHITE_QUEEN: true,
	},
}

func (b Board) isCardinalAttacked(sq, side int) bool {
	if b.PieceAt(sq) == OFF_BOARD {
		return false
	}
	for _, d := range ROOK_DELTAS {
		nsq := sq + d
		for b.PieceAt(nsq) == EMPTY_SQUARE {
			nsq += d
		}
		p := b.PieceAt(nsq)
		if _, ok := CARDINAL_ATTACKERS[side][p]; ok {
			return true
		}
	}

	return false
}

var ORDINAL_ATTACKERS = map[int]map[int]bool{
	WHITE: map[int]bool{
		BLACK_BISHOP: true,
		BLACK_QUEEN:  true,
	},
	BLACK: map[int]bool{
		WHITE_BISHOP: true,
		WHITE_QUEEN:  true,
	},
}

func (b Board) isOrdinalAttacked(sq, side int) bool {
	if b.PieceAt(sq) == OFF_BOARD {
		return false
	}
	for _, d := range BISHOP_DELTAS {
		nsq := sq + d
		for b.PieceAt(nsq) == EMPTY_SQUARE {
			nsq += d
		}
		p := b.PieceAt(nsq)
		if _, ok := ORDINAL_ATTACKERS[side][p]; ok {
			return true
		}
	}

	return false
}

var KING_ATTACKER = map[int]int{
	WHITE: BLACK_KING,
	BLACK: WHITE_KING,
}

func (b Board) isKingAttacked(sq, side int) bool {
	if b.PieceAt(sq) == OFF_BOARD {
		return false
	}
	for _, d := range QUEEN_DELTAS {
		if b.PieceAt(sq+d) == KING_ATTACKER[side] {
			return true
		}
	}

	return false
}
