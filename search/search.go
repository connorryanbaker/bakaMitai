package search

import (
	"github.com/connorryanbaker/bakaMitai/board"
	"github.com/connorryanbaker/bakaMitai/eval"

	"math"
)

// s/o bruce moreland pv collection

var nodes int
var nc = make([]int, 0)

type Line struct {
	NumMoves int
	Moves    []board.Move
}

func NewLine(depth int) Line {
	return Line{
		NumMoves: depth,
		Moves:    make([]board.Move, depth),
	}
}

func Search(b *board.Board, depth int, pv *Line) []board.Move {
	nodes = 0
	negamax(b, depth, math.Inf(-1), math.Inf(1), pv)
	nc = append(nc, nodes)
	s := 0
	for _, v := range nc {
		s += v
	}
	return pv.Moves
}

func negamax(b *board.Board, depth int, alpha, beta float64, pv *Line) float64 {
	lpv := NewLine(depth)
	moves := siftPV(pv.Moves[0], b.LegalMoves())
	if depth == 0 || len(moves) == 0 {
		nodes += 1
		pv.NumMoves = 0
		return eval.NegamaxEval(*b)
	}

	for _, m := range moves {
		b.MakeMove(m)
		// todo: continue check on shortcircuiting if checkmate / draw here
		var v float64
		if !b.Drawn() {
			v = -negamax(b, depth-1, -beta, -alpha, &lpv)
		}
		b.UnmakeMove()
		if v >= beta {
			return beta
		}
		if v > alpha {
			alpha = v
			pv.Moves[0] = m
			for i := 0; i < lpv.NumMoves; i++ {
				pv.Moves[i+1] = lpv.Moves[i]
			}
			pv.NumMoves = lpv.NumMoves + 1
		}
	}
	return alpha
}

func siftPV(pvMove board.Move, legalMoves []board.Move) []board.Move {
	if pvMove.IsNull() {
		return legalMoves
	}
	for i, m := range legalMoves {
		if board.EqualMoves(pvMove, m) && i > 0 {
			for j := i; j > 0; j-- {
				tmp := legalMoves[j-1]
				legalMoves[j-1] = legalMoves[j]
				legalMoves[j] = tmp
			}
		}
	}
	return legalMoves
}
