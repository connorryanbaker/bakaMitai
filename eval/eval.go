package eval

import (
	"github.com/connorryanbaker/engine/board"
	"math"
)

func Eval(b board.Board) float64 {
	if b.Checkmate() {
		if b.Side() == board.WHITE {
			return math.Inf(-1)
		}
		return math.Inf(1)
	}

	if b.Drawn() {
		return 0
	}

	// TODO: tests to ensure balance of evals
	return evalMaterial(b) + evalPieceSquares(b)*5
}
