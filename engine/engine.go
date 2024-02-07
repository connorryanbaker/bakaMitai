package engine

import (
	"github.com/connorryanbaker/bakaMitai/board"
	"github.com/connorryanbaker/bakaMitai/search"
)

type Engine struct {
	Depth int
	PV    search.Line
}

func New(depth int) Engine {
	return Engine{
		Depth: depth,
		PV:    search.NewLine(depth),
	}
}

func (e *Engine) GenMove(b *board.Board) board.Move {
	moves := search.Search(b, e.Depth, &e.PV)
	m := moves[0]
	e.siftPVMoves()
	return m
}

func (e *Engine) siftPVMoves() {
	for i := 1; i < e.Depth; i++ {
		e.PV.Moves[i-1] = e.PV.Moves[i]
	}
}
