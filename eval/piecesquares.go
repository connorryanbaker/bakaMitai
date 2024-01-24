package eval

import (
	"github.com/connorryanbaker/engine/board"
)

// definitions from chess programming simplified evaluation function
var WHITE_PAWN_SQ_VALUES = [64]float64{
	0, 0, 0, 0, 0, 0, 0, 0,
	50, 50, 50, 50, 50, 50, 50, 50,
	10, 10, 20, 30, 30, 20, 10, 10,
	5, 5, 10, 25, 25, 10, 5, 5,
	0, 0, 0, 20, 20, 0, 0, 0,
	5, -5, -10, 0, 0, -10, -5, 5,
	5, 10, 10, -20, -20, 10, 10, 5,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var BLACK_PAWN_SQ_VALUES = [64]float64{
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 10, 10, -20, -20, 10, 10, 5,
	5, -5, -10, 0, 0, -10, -5, 5,
	0, 0, 0, 20, 20, 0, 0, 0,
	5, 5, 10, 25, 25, 10, 5, 5,
	10, 10, 20, 30, 30, 20, 10, 10,
	50, 50, 50, 50, 50, 50, 50, 50,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var KNIGHT_SQ_VALUES = [64]float64{
	-50, -40, -30, -30, -30, -30, -40, -50,
	-40, -20,   0,   0,   0,   0, -20, -40,
	-30,   0,  15,  15,  15,  15,   0, -30,
	-30,   5,  20,  20,  20,  20,   5, -30,
	-30,   0,  20,  20,  20,  20,   0, -30,
	-30,   5,  15,  15,  15,  15,   5, -30,
	-40, -20,   0,   5,   5,   0, -20, -40,
	-50, -40, -30, -30, -30, -30, -40, -50,
}

var WHITE_BISHOP_SQ_VALUES = [64]float64{
	-20, -10, -10, -10, -10, -10, -10, -20,
	-10,   0,   0,   0,   0,   0,   0, -10,
	-10,   0,   5,  10,  10,   5,   0,   5,
	-10,   5,   5,  10,  10,   5,   5, -10,
	-10,   0,  10,  10,  10,  10,   0, -10,
	-10,  10,  10,  10,  10,  10,  10, -10,
	-10,   5,   0,   0,   0,   0,   5, -10,
	-20, -10, -10, -10, -10, -10, -10, -20,
}

var BLACK_BISHOP_SQ_VALUES = [64]float64{
	-20, -10, -10, -10, -10, -10, -10, -20,
	-10,   5,   0,   0,   0,   0,   5, -10,
	-10,  10,  10,  10,  10,  10,  10, -10,
	-10,   0,  10,  10,  10,  10,   0, -10,
	-10,   5,   5,  10,  10,   5,   5, -10,
	-10,   0,   5,  10,  10,   5,   0,   5,
	-10,   0,   0,   0,   0,   0,   0, -10,
	-20, -10, -10, -10, -10, -10, -10, -20,
}

var WHITE_ROOK_SQ_VALUES = [64]float64{
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 10, 10, 10, 10, 10, 10, 5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	7, 7, 7, 7, 7, 7, 7, 7, // test to bait rook to move
	-5, 0, 0, 0, 0, 0, 0, -5,
	0, 0, 0, 5, 5, 3, 0, 0,
}

var BLACK_ROOK_SQ_VALUES = [64]float64{
	0, 0, 0, 5, 5, 3, 0, 0,
	-5, 0, 0, 0, 0, 0, 0, -5,
	7, 7, 7, 7, 7, 7, 7, 7,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	5, 10, 10, 10, 10, 10, 10, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
}

var WHITE_KING_MIDDLE_GAME_SQ_VALUES = [64]float64{
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-20, -30, -30, -40, -40, -30, -30, -20,
	-10, -20, -20, -20, -20, -20, -20, -10,
	20, 20, 0, 0, 0, 0, 20, 20,
	20, 30, 10, 0, 0, 10, 30, 20,
}
var BLACK_KING_MIDDLE_GAME_SQ_VALUES = [64]float64{
	20, 30, 10, 0, 0, 10, 30, 20,
	20, 20, 0, 0, 0, 0, 20, 20,
	-10, -20, -20, -20, -20, -20, -20, -10,
	-20, -30, -30, -40, -40, -30, -30, -20,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
}

var KING_END_GAME_SQ_VALUES = [64]float64{
	-50, -40, -30, -20, -20, -30, -40, -50,
	-30, -20, -10, 0, 0, -10, -20, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 30, 40, 40, 30, -10, -30,
	-30, -10, 20, 30, 30, 20, -10, -30,
	-30, -30, 0, 0, 0, 0, -30, -30,
	-50, -30, -30, -30, -30, -30, -30, -50,
}

var PIECE_TO_LOOKUP = map[int][64]float64{
	board.WHITE_PAWN:   WHITE_PAWN_SQ_VALUES,
	board.BLACK_PAWN:   BLACK_PAWN_SQ_VALUES,
	board.WHITE_KNIGHT: KNIGHT_SQ_VALUES,
	board.BLACK_KNIGHT: KNIGHT_SQ_VALUES,
	board.WHITE_BISHOP: WHITE_BISHOP_SQ_VALUES,
	board.BLACK_BISHOP: BLACK_BISHOP_SQ_VALUES,
	board.WHITE_ROOK:   WHITE_ROOK_SQ_VALUES,
	board.BLACK_ROOK:   BLACK_ROOK_SQ_VALUES,
	board.WHITE_KING: WHITE_KING_MIDDLE_GAME_SQ_VALUES,
	board.BLACK_KING: BLACK_KING_MIDDLE_GAME_SQ_VALUES,
}

var PIECE_WEIGHTS = map[int]float64{
	board.WHITE_PAWN:   3,
	board.BLACK_PAWN:   -3,
	board.WHITE_KNIGHT: 3,
	board.BLACK_KNIGHT: -3,
	board.WHITE_BISHOP: 3.5,
	board.BLACK_BISHOP: -3.5,
	board.WHITE_ROOK:   3,
	board.BLACK_ROOK:   -3,
	board.WHITE_KING: 10,
	board.BLACK_KING: -10,
}

func evalPieceSquares(b board.Board) float64 {
	var e float64
	for p, sqs := range b.PieceSquares {
		lookup, ok := PIECE_TO_LOOKUP[p]
    if !ok {
      continue
    }

		if p == board.WHITE_KING && b.PieceSquares[board.BLACK_QUEEN] == nil {
			lookup = KING_END_GAME_SQ_VALUES
		}
		if p == board.BLACK_KING && b.PieceSquares[board.WHITE_QUEEN] == nil {
			lookup = KING_END_GAME_SQ_VALUES
		}

		for _, sq := range sqs {
			v := lookup[board.SQ_NAME_TO_SQ_64[sq]]
			e += v * PIECE_WEIGHTS[p]
		}
	}
	return e
}
