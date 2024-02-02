package search

import (
	"github.com/connorryanbaker/bakaMitai/board"
	"github.com/connorryanbaker/bakaMitai/eval"

	"math"
)

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

// s/o bruce moreland pv collection

type line struct {
	nummoves int
	moves    []board.Move
}

// tmp: hacky workaround for dev metrics
func SearchNodeCount(b *board.Board, depth int, nodes *int) []board.Move {
	pv := line{depth, make([]board.Move, depth)}
	negamaxNodeCount(b, depth, math.Inf(-1), math.Inf(1), pv, nodes)
	return pv.moves
}

func negamaxNodeCount(b *board.Board, depth int, alpha, beta float64, pv line, nodes *int) float64 {
	lpv := line{depth, make([]board.Move, depth)}
	moves := b.LegalMoves()
	if depth == 0 || len(moves) == 0 {
		*nodes += 1
		pv.nummoves = 0
		return eval.NegamaxEval(*b)
	}

	for _, m := range moves {
		b.MakeMove(m)
		v := -negamaxNodeCount(b, depth-1, -beta, -alpha, lpv, nodes)
		b.UnmakeMove()
		if v >= beta {
			return beta
		}
		if v > alpha {
			alpha = v
			pv.moves[0] = m
			for i := 0; i < lpv.nummoves-1; i++ {
				pv.moves[i+1] = lpv.moves[i]
			}
			pv.nummoves = lpv.nummoves + 1
		}
	}
	return alpha
}
