package search

import (
	"fmt"
	"github.com/connorryanbaker/engine/board"
	"github.com/connorryanbaker/engine/eval"
)

// TODO: alphabeta pruning and move sorting

func Search(b *board.Board, depth int) (float64, board.Move) {
	maximizing := b.Side() == board.WHITE
	bhash := b.Hash()
	e, m := minimax(b, maximizing, depth)
	if bhash != b.Hash() {
		panic("wtf")
	}
	return e, m
}

func minimax(b *board.Board, maximizing bool, depth int) (float64, board.Move) {
	if depth == 0 || len(b.LegalMoves()) == 0 {
		eval := eval.Eval(*b)
		history := b.History()
		return eval, history[len(history)-1].Move()
	}
	moves := b.LegalMoves()
	evals := make([]float64, len(moves))
	for i, m := range moves {
		bhash := b.Hash()
		b.MakeMove(m)
		eval, _ := minimax(b, !maximizing, depth-1)
		b.UnmakeMove()
		if bhash != b.Hash() {
			b.Print()
			fmt.Println(m)
			fmt.Println(b.History())
			panic(m)
		}
		evals[i] = eval
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
