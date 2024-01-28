package search

import (
	"github.com/connorryanbaker/engine/board"
	"github.com/connorryanbaker/engine/eval"

	"math"
)

// TODO:
// test move ordering and node counts [ ]
// quiescence [ ]
// transposition tables [ ]

func Search(b *board.Board, depth int) []board.Move {
	pv := make([]board.Move, depth)
	negamax(b, depth, math.Inf(-1), math.Inf(1), pv)
	return pv
}

func negamax(b *board.Board, depth int, alpha, beta float64, line []board.Move) float64 {
	moves := b.LegalMoves()
	if depth == 0 || len(moves) == 0 {
		return eval.NegamaxEval(*b)
	}

	for _, m := range moves {
		b.MakeMove(m)
		v := -negamax(b, depth-1, -beta, -alpha, line)
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
