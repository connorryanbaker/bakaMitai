package engine

import (
	"github.com/connorryanbaker/bakaMitai/board"
	"github.com/connorryanbaker/bakaMitai/search"

	"time"
)

type Engine struct {
	Depth int
	PV    search.Line
}

func New() Engine {
	return Engine{
		PV: search.NewLine(1),
	}
}

func (e *Engine) GenMove(b *board.Board) board.Move {
	var m board.Move
	now := time.Now()
	exp := now.Add(time.Minute * 2)
	for depth := 1; ; depth++ {
		prev := e.PV.Moves
		if depth > len(e.PV.Moves) {
			e.PV = search.NewLine(depth)
			for j := 0; j < min(depth, len(prev)); j++ {
				e.PV.Moves[j] = prev[j]
			}
		}
		moves := search.Search(b, depth, &e.PV, exp)
		m = moves[0]
		e.siftPVMoves()
		if time.Now().After(exp) {
			break
		}
	}
	return m
}

func (e *Engine) siftPVMoves() {
	if e.Depth == 1 {
		return
	}

	for i := 1; i < e.Depth; i++ {
		e.PV.Moves[i-1] = e.PV.Moves[i]
	}
}
