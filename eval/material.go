package eval

import "github.com/connorryanbaker/bakaMitai/board"

var PIECE_VALUES = map[int]float64{
	board.WHITE_PAWN:   1000,
	board.WHITE_KNIGHT: 3000,
	board.WHITE_BISHOP: 3250,
	board.WHITE_ROOK:   5000,
	board.WHITE_QUEEN:  9000,
	board.BLACK_PAWN:   -1000,
	board.BLACK_KNIGHT: -3000,
	board.BLACK_BISHOP: -3250,
	board.BLACK_ROOK:   -5000,
	board.BLACK_QUEEN:  -9000,
}

func evalMaterial(b board.Board) float64 {
	var e float64
	for p, sqs := range b.PieceSquares {
		e += PIECE_VALUES[p] * float64(len(sqs))
	}
	return e
}
