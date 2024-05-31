package board

func BBperft(b *Board, depth int) uint64 {
	var nodes uint64
	if depth == 0 {
		return 1
	}

	bbm := b.GenerateBitboardMoves()
	for _, m := range bbm {
		b.MakeMove(m)
		nodes += BBperft(b, depth-1)
		b.UnmakeMove()
	}
	return nodes
}
