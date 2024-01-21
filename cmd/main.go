package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/connorryanbaker/engine/board"
	"github.com/connorryanbaker/engine/eval"
)

func main() {
	// b := board.FromFENString("6q1/8/8/8/8/5k2/8/7K b - - 0 1")
	b := board.NewBoard()
	playRandomGame(b)
}

func playRandomGame(b board.Board) {
	for true {
		m := b.LegalMoves()

		sort.Slice(m, func(i, j int) bool {
			b.MakeMove(m[i])
			e1 := eval.Eval(b)
			b.UnmakeMove()
			b.MakeMove(m[j])
			e2 := eval.Eval(b)
			b.UnmakeMove()
			if b.Side() == board.WHITE {
				return e1 > e2
			}
			return e2 > e1
		})
		b.MakeMove(m[0])
		b.Print()
		time.Sleep(1 * time.Second)
		if b.Checkmate() {
			fmt.Println("checkmate!")
			return
		} else if b.Stalemate() {
			fmt.Println("stalemate!")
			return
		} else if b.FiftyMoveDraw() {
			fmt.Println("50 move draw!")
			return
		} else if b.ThreefoldRepetition() {
			fmt.Println("3fold!")
			return
		} else if b.InsufficientMaterial() {
			fmt.Println("insufficient material!")
			return
		}
	}
}
