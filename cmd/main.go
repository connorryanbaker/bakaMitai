package main

import (
	"fmt"

	"github.com/connorryanbaker/engine/board"
	"github.com/connorryanbaker/engine/search"
)

func main() {
	b := board.NewBoard()
	// moves := search.Search(&b, 2)
	play(b)
}

func play(b board.Board) {
	for true {
		b.Print()
		moves := search.Search(&b, 2)
		b.MakeMove(moves[0])
		if b.Checkmate() {
			b.Print()
			fmt.Println("checkmate!")
			printHistory(b)
			return
		} else if b.Stalemate() {
			b.Print()
			fmt.Println("stalemate!")
			printHistory(b)
			return
		} else if b.FiftyMoveDraw() {
			b.Print()
			fmt.Println("50 move draw!")
			printHistory(b)
			return
		} else if b.ThreefoldRepetition() {
			b.Print()
			fmt.Println("3fold!")
			printHistory(b)
			return
		} else if b.InsufficientMaterial() {
			b.Print()
			fmt.Println("insufficient material!")
			printHistory(b)
			return
		}
	}
}

func printHistory(b board.Board) {
	h := b.History
	for i := 0; i < len(h); i++ {
		fmt.Printf("%d: FROM: %s TO: %s\n", i+1, board.SQ_NUM_TO_NAME[h[i].Move.From], board.SQ_NUM_TO_NAME[h[i].Move.To])
	}
}
