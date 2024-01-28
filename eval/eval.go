package eval

import (
	"github.com/connorryanbaker/engine/board"
	"math"
)

func Eval(b board.Board) float64 {
	if b.Checkmate() {
		if b.Side == board.WHITE {
			return math.MaxFloat64 * -1
		}
		return math.MaxFloat64
	}

	if b.Drawn() {
		return 0
	}

	// TODO: tests to ensure balance of evals
	return evalMaterial(b) + evalPieceSquares(b)*5 + evalMobility(b)
}

func NegamaxEval(b board.Board) float64 {
	mult := float64(1)
	if b.Side == board.BLACK {
		mult = float64(-1)
	}

	return Eval(b) * mult
}
