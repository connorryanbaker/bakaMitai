package board

// import (
//    "fmt"
// )

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

func BBperft(b *Board, depth int) uint64 {
	var nodes uint64
	if depth == 0 {
		return 1
	}

	bbm := b.GenerateBitboardMoves()
	// lm := b.LegalMoves()
	// if len(bbm) != len(lm) {
	//   fmt.Println("ERR!", len(lm), len(bbm))
	//   b.Print()
	//   fmt.Println("BBM")
	//   for _, m := range bbm {
	//     m.Print()
	//   }
	//   fmt.Println("LM")
	//   for _, m := range lm {
	//     m.Print()
	//   }
	//   return nodes
	// }

	for _, m := range bbm {
		b.MakeBBMove(m)
		nodes += BBperft(b, depth-1)
		b.UnmakeMove()
	}
	return nodes
}
