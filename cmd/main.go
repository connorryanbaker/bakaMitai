package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/connorryanbaker/engine/board"
)

func main() {
	b := board.FromFENString("7q/8/8/8/8/4k3/8/7K w - - 0 1")
	playRandomGame(b)
}

func playRandomGame(b board.Board) {
	for true {
		m := b.LegalMoves()
		b.MakeMove(m[rand.Intn(len(m))])
		b.Print()
		time.Sleep(1 * time.Millisecond)
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
