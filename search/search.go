package search

import (
	"github.com/connorryanbaker/engine/board"
	"github.com/connorryanbaker/engine/eval"

	"math"
)

func Search(b *board.Board, depth int, nodes *int) []board.Move {
	pv := make([]board.Move, depth)
	negamax(b, depth, math.Inf(-1), math.Inf(1), pv, nodes)
	return pv
}

func negamax(b *board.Board, depth int, alpha, beta float64, line []board.Move, nodes *int) float64 {
	moves := b.LegalMoves()
	if depth == 0 || len(moves) == 0 {
		*nodes += 1
		return eval.NegamaxEval(*b)
	}

	for _, m := range moves {
		b.MakeMove(m)
		v := -negamax(b, depth-1, -beta, -alpha, line, nodes)
		b.UnmakeMove()
		if v >= beta {
			return beta
		}
		if v > alpha {
			alpha = v
			line[len(line)-depth] = m
		}
	}
	return alpha
}
