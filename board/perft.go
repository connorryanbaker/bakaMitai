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

func bbperft(b *Board, depth int) uint64 {
	var nodes uint64
	if depth == 0 {
		return 1
	}

	for _, m := range b.GenerateBitboardMoves() {
		b.MakeMove(m)
		nodes += perft(b, depth-1)
		b.UnmakeMove()
	}
	return nodes
}
