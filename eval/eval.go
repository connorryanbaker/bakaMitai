package eval

import (
	"github.com/connorryanbaker/bakaMitai/board"

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

	return (evalMaterial(b) + evalPieceSquares(b) + evalPawnStructure(b)*100) / 1000
}

func NegamaxEval(b board.Board) float64 {
	mult := float64(1)
	if b.Side == board.BLACK {
		mult = float64(-1)
	}

	return Eval(b) * mult
}
