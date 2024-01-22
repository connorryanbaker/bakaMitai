package eval

import "github.com/connorryanbaker/engine/board"

var PIECE_VALUES = map[int]float64{
	board.WHITE_PAWN:   10,
	board.WHITE_KNIGHT: 30,
	board.WHITE_BISHOP: 35,
	board.WHITE_ROOK:   50,
	board.WHITE_QUEEN:  90,
	board.BLACK_PAWN:   -10,
	board.BLACK_KNIGHT: -30,
	board.BLACK_BISHOP: -35,
	board.BLACK_ROOK:   -50,
	board.BLACK_QUEEN:  -90,
}

func evalMaterial(b board.Board) float64 {
	var e float64
	for p, sqs := range b.Pieces() {
		e += PIECE_VALUES[p] * float64(len(sqs))
	}
	return e
}
