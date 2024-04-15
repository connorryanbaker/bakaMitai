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

	return (evalMaterial(b) + evalPieceSquares(b) + evalPawnStructure(b)*100 + evalMobility(b)*5) / 1000
}

func NegamaxEval(b board.Board) float64 {
	mult := float64(1)
	if b.Side == board.BLACK {
		mult = float64(-1)
	}

	return Eval(b) * mult
}

func evalMobility(b board.Board) float64 {
	whiteMobility := b.RayMobility(board.WHITE, board.WHITE_ROOK) + b.RayMobility(board.WHITE, board.WHITE_BISHOP)
	blackMobility := b.RayMobility(board.BLACK, board.BLACK_ROOK) + b.RayMobility(board.BLACK, board.BLACK_BISHOP)
	return float64(whiteMobility - blackMobility)
}
