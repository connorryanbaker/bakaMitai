package main

import (
	"fmt"

	"github.com/connorryanbaker/engine/board"
	"github.com/connorryanbaker/engine/search"
)

func main() {
	b := board.FromFENString("r2qk2r/pppb1ppp/8/1B2N3/1b2p3/8/PPPBPPPP/R2QK2R w KQkq - 1 1")
	// b := board.NewBoard()
	play(b)
}

func play(b board.Board) {
	for true {
		b.Print()
		_, m := search.Search(&b, 2)
		// fmt.Println(e)
		// fmt.Println(b.Hash())
		// fmt.Println(m)
		b.MakeMove(m)
		if b.Checkmate() {
			b.Print()
			fmt.Println("checkmate!")
			return
		} else if b.Stalemate() {
			b.Print()
			fmt.Println("stalemate!")
			return
		} else if b.FiftyMoveDraw() {
			b.Print()
			fmt.Println("50 move draw!")
			return
		} else if b.ThreefoldRepetition() {
			b.Print()
			fmt.Println("3fold!")
			return
		} else if b.InsufficientMaterial() {
			b.Print()
			fmt.Println("insufficient material!")
			return
		}
	}
}
