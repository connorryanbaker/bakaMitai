package main

import (
	"fmt"

	"github.com/connorryanbaker/engine/board"
	"github.com/connorryanbaker/engine/search"
)

func main() {
	// b := board.FromFENString("1q6/8/8/8/8/5k2/8/7K b - - 0 1")
	b := board.NewBoard()
	playRandomGame(b)
}

func playRandomGame(b board.Board) {
	for true {
		b.Print()
		_, m := search.Search(&b, 2)
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
