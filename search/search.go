package search

import (
	"github.com/connorryanbaker/engine/board"
	"github.com/connorryanbaker/engine/eval"

	"math"
	"sort"
)

// TODO: quiescence and transposition tables

func Search(b *board.Board, depth int) []board.Move {
	maximizing := b.Side == board.WHITE
	line := make([]board.Move, depth)
	alphaBeta(b, maximizing, math.Inf(-1), math.Inf(1), depth, line)
	return line
}

func alphaBeta(b *board.Board, maximizing bool, alpha, beta float64, depth int, line []board.Move) float64 {
	if depth == 0 || len(b.LegalMoves()) == 0 {
		return eval.Eval(*b)
	}

	moves := b.LegalMoves()
	sort.Slice(moves, func(a, b int) bool {
		return moves[a].Score() > moves[b].Score()
	})

	if maximizing {
		max := math.Inf(-1)
		for _, m := range moves {
			b.MakeMove(m)
			v := alphaBeta(b, !maximizing, alpha, beta, depth-1, line)
			b.UnmakeMove()
			if max < v {
				max = v
				line[len(line)-depth] = m
			}
			if alpha < max {
				alpha = max
			}
			if beta <= alpha {
				break
			}
		}
		return max
	}
	min := math.Inf(1)
	for _, m := range moves {
		b.MakeMove(m)
		v := alphaBeta(b, !maximizing, alpha, beta, depth-1, line)
		b.UnmakeMove()
		if v < min {
			min = v
			line[len(line)-depth] = m
		}
		if min < beta {
			beta = min
		}
		if beta <= alpha {
			break
		}
	}
	return min
}

func minimax(b *board.Board, maximizing bool, depth int) (float64, board.Move) {
	if depth == 0 || len(b.LegalMoves()) == 0 {
		eval := eval.Eval(*b)
		return eval, b.History[len(b.History)-1].Move
	}
	moves := b.LegalMoves()
	evals := make([]float64, len(moves))
	for i, m := range moves {
		b.MakeMove(m)
		eval, _ := minimax(b, !maximizing, depth-1)
		evals[i] = eval
		b.UnmakeMove()
	}
	var idx int
	if maximizing {
		idx = findIdx(evals, func(a, b float64) bool {
			return a > b
		})
	} else {
		idx = findIdx(evals, func(a, b float64) bool {
			return a < b
		})
	}

	return evals[idx], moves[idx]
}

func findIdx(e []float64, c func(a, b float64) bool) int {
	if len(e) == 0 {
		return -1
	}

	m := e[0]
	idx := 0
	for i := 1; i < len(e); i++ {
		if c(e[i], m) {
			m = e[i]
			idx = i
		}
	}

	return idx
}
