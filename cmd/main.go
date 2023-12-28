package main

import (
	"fmt"
	"github.com/connorryanbaker/engine/board"
)

func main() {
	b := board.NewBoard()
	fmt.Println(b.PieceAt(0))
}
