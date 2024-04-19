package search

import (
	"github.com/connorryanbaker/bakaMitai/board"
	"github.com/connorryanbaker/bakaMitai/eval"

	"math"
	"time"
)

// s/o bruce moreland pv collection

var nodes int
var totalnodes int
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

func Search(b *board.Board, depth int, pv *Line, exp time.Time) []board.Move {
	negamax(b, depth, math.Inf(-1), math.Inf(1), pv, exp)
	return pv.Moves
}

func negamax(b *board.Board, depth int, alpha, beta float64, pv *Line, exp time.Time) float64 {
	if time.Now().After(exp) {
		return alpha
	}
	lpv := NewLine(depth)
	moves := b.GenerateBitboardMoves()
	if depth == 0 || len(moves) == 0 {
		pv.NumMoves = 0
		return quiesce(b, alpha, beta)
	}
	siftPV(pv.Moves[0], moves)
	depthBestEval := depth
	for _, m := range moves {
		b.MakeBBMove(m)
		draw := b.Drawn()
		var v float64
		if !draw {
			v = -negamax(b, depth-1, -beta, -alpha, &lpv, exp)
		}
		b.UnmakeMove()
		if v >= beta {
			return beta
		}
		if v > alpha || (v == math.MaxFloat64 && lpv.NumMoves < depthBestEval) {
			if lpv.NumMoves < depthBestEval {
				depthBestEval = lpv.NumMoves
			}

			alpha = v
			pv.Moves[0] = m
			for i := 0; i < min(len(pv.Moves)-1, len(lpv.Moves)); i++ {
				pv.Moves[i+1] = lpv.Moves[i]
			}
			pv.NumMoves = lpv.NumMoves + 1
		}
	}
	return alpha
}

func quiesce(b *board.Board, alpha, beta float64) float64 {
	standPat := eval.NegamaxEval(*b)
	if standPat >= beta {
		return beta
	}
	if alpha < standPat {
		alpha = standPat
	}
	captures := b.GenerateCaptures()
	for _, capture := range captures {
		if capture.See(b) >= 0 {
			b.MakeBBMove(capture)
			score := -1 * quiesce(b, -beta, -alpha)
			b.UnmakeMove()
			if score >= beta {
				return beta
			}
			if score > alpha {
				alpha = score
			}
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
