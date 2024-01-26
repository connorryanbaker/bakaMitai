package main

import (
	"fmt"

	"github.com/connorryanbaker/engine/board"
	"github.com/connorryanbaker/engine/search"
)

func main() {
	// b := board.FromFENString("r2qkbnr/ppp1pppp/2n5/1B6/4p1b1/5N1P/PPPP1PP1/RNBQK2R b KQkq - 0 5")
	// b := board.FromFENString("r2qkbnr/ppp1pppp/2n5/1B6/4p1P1/5b1P/PPPP1P2/RNBQK2R b KQkq - 0 6")
	// b := board.FromFENString("r2qkbnr/ppp1pppp/2n5/1B6/4p3/5b1P/PPPP1PP1/RNBQK2R w KQkq - 0 6")
	b := board.NewBoard()
	play(b)
}

func play(b board.Board) {
	for true {
		b.Print()
		m := search.Search(&b, 4)
		b.MakeMove(m[0])
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
