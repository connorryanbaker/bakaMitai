package board

import "fmt"

func perft(b *Board, depth int) uint64 {
	var nodes uint64
	if depth == 0 {
		return 1
	}

	for _, m := range b.LegalMoves() {
		b.MakeMove(m)
		nodes += perft(b, depth-1)
		b.UnmakeMove()
	}
	return nodes
}

func bbperft(b *Board, depth int) uint64 {
	var nodes uint64
	if depth == 0 {
		return 1
	}

	bbm := b.GenerateBitboardMoves()
	lm := b.LegalMoves()
	if len(bbm) != len(lm) {
		fmt.Println("ERR!")
		fmt.Println(b.History)
		b.Print()
		fmt.Println("LEGALMOVES", len(lm))
		for _, m := range lm {
			m.Print()
		}
		fmt.Println("BITBOARDMOVES", len(bbm))
		for _, m := range bbm {
			m.Print()
		}
	}

	for _, m := range bbm {
		v := b.MakeMove(m)
		nodes += bbperft(b, depth-1)
		if v {
			b.UnmakeMove()
		}
	}
	return nodes
}
