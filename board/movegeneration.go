package board

import (
	"errors"
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

func (b Board) MovesForPieceFromConvertedIdxSquare(sq int) ([]Move, error) {
	piece := b.PieceFromConvertedIdx(sq)
	switch piece {
	case WHITE_PAWN:
		return b.WhitePawnMoves(sq), nil
	// case WHITE_KNIGHT:
	//   return b.WhiteKnightMoves(sq), nil
	// case WHITE_BISHOP:
	//   return b.WhiteBishopMoves(sq), nil
	// case WHITE_ROOK:
	//   return b.WhiteRookMoves(sq), nil
	// case WHITE_QUEEN:
	//   return b.WhiteQueenMoves(sq), nil
	// case WHITE_KING:
	//   return b.WhiteKingMoves(sq), nil
	case BLACK_PAWN:
		return b.BlackPawnMoves(sq), nil
	// case BLACK_KNIGHT:
	//   return b.WhiteKnightMoves(sq), nil
	// case BLACK_BISHOP:
	//   return b.WhiteBishopMoves(sq), nil
	// case BLACK_ROOK:
	//   return b.WhiteRookMoves(sq), nil
	// case BLACK_QUEEN:
	//   return b.WhiteQueenMoves(sq), nil
	// case BLACK_KING:
	//   return b.WhiteKingMoves(sq), nil
	default:
		return nil, errors.New("invalid square")
	}
}

// TODO: EP
func (b Board) WhitePawnMoves(sq int) []Move { // pass in MAILBOX_64[A8 - H1]
	moves := make([]Move, 0) // TODO: len / capacity to avoid all appends
	for _, d := range WHITE_PAWN_ATTACKS {
		ns := sq + d
		if b.PieceFromConvertedIdx(ns) > 6 { // black piece
			if ns <= IH8 { // check promotion
				for _, piece := range WHITE_PROMOTION_PIECES {
					moves = append(moves, Move{sq, ns, true, false, false, true, piece, false})
				}
			} else {
				moves = append(moves, Move{sq, ns, true, false, false, false, WHITE_PAWN, false})
			}
		} else if b.ep != nil && *b.ep == ns {
			moves = append(moves, Move{sq, ns, true, false, false, false, WHITE_PAWN, false})
		}
	}
	// one sq push
	if b.PieceFromConvertedIdx(sq+WHITE_PAWN_DELTAS[0]) == EMPTY_SQUARE {
		ns := sq + WHITE_PAWN_DELTAS[0]
		if ns <= IH8 { // check promotion
			for _, piece := range WHITE_PROMOTION_PIECES {
				moves = append(moves, Move{sq, ns, false, false, false, true, piece, false})
			}
		} else {
			moves = append(moves, Move{sq, ns, false, false, false, false, WHITE_PAWN, false})
			if IA2 <= sq && sq <= IH2 && b.PieceFromConvertedIdx(sq+WHITE_PAWN_DELTAS[1]) == EMPTY_SQUARE {
				moves = append(moves, Move{sq, sq + WHITE_PAWN_DELTAS[1], false, false, false, false, WHITE_PAWN, true})
			}
		}
	}
	// opening two sq push
	return moves
}

func (b Board) BlackPawnMoves(sq int) []Move {
	moves := make([]Move, 0)
	for _, d := range BLACK_PAWN_ATTACKS {
		ns := sq + d
		if b.PieceFromConvertedIdx(ns) < 7 && 0 < b.PieceFromConvertedIdx(ns) { // white piece
			if ns >= IA1 { // check promotion
				for _, piece := range BLACK_PROMOTION_PIECES {
					moves = append(moves, Move{sq, ns, true, false, false, true, piece, false})
				}
			} else {
				moves = append(moves, Move{sq, ns, true, false, false, false, BLACK_PAWN, false})
			}
		} else if b.ep != nil && *b.ep == ns {
			moves = append(moves, Move{sq, ns, true, false, false, false, BLACK_PAWN, false})
		}
	}
	// one sq push
	if b.PieceFromConvertedIdx(sq+BLACK_PAWN_DELTAS[0]) == EMPTY_SQUARE {
		ns := sq + BLACK_PAWN_DELTAS[0]
		if ns >= IA1 { // check promotion
			for _, piece := range BLACK_PROMOTION_PIECES {
				moves = append(moves, Move{sq, ns, false, false, false, true, piece, false})
			}
		} else {
			moves = append(moves, Move{sq, ns, false, false, false, false, BLACK_PAWN, false})
			if IA7 <= sq && sq <= IH7 && b.PieceFromConvertedIdx(sq+BLACK_PAWN_DELTAS[1]) == EMPTY_SQUARE {
				moves = append(moves, Move{sq, sq + BLACK_PAWN_DELTAS[1], false, false, false, false, BLACK_PAWN, true})
			}
		}
	}
	// opening two sq push
	return moves
}