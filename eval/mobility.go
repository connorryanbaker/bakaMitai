package eval

import "github.com/connorryanbaker/engine/board"

func evalMobility(b board.Board) float64 {
	return float64(len(b.LegalMoves())) * float64(10) * sign(b)
}

func sign(b board.Board) float64 {
	if b.Side == board.WHITE {
		return float64(1)
	}
	return float64(-1)
}
