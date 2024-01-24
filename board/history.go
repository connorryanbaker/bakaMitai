package board

type History struct {
	previousSquareOccupant int
	Move                   Move
	castle                 [4]bool
	ep                     *int
	hply                   int
	ply                    int
	Hash                   uint64
}

func (b *Board) pushHistory(m Move) {
	h := History{
		previousSquareOccupant: b.PieceAt(m.To),
		Move:                   m,
		castle:                 b.Castle,
		ep:                     b.Ep,
		hply:                   b.Hply,
		ply:                    b.Ply,
		Hash:                   b.Hash(),
	}
	b.History = append(b.History, h)
}

func (b *Board) popHistory() {
	b.History[len(b.History)-1] = History{}
	b.History = b.History[:len(b.History)-1]
}
