package board

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
	for _, m := range bbm {
		b.MakeBBMove(m)
		nodes += BBperft(b, depth-1)
		b.UnmakeMove()
	}
	return nodes
}
