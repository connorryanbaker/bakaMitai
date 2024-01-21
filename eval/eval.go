package eval

import (
	"github.com/connorryanbaker/engine/board"
	"math"
)

// TODO: minimize number of time board generates moves / checks
// attacks etc.

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

	return evalMaterial(b)
}
