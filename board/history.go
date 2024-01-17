package board

type History struct {
	previousSquareOccupant int
	move                   Move
	castle                 [4]bool
	ep                     *int
	hply                   int
	ply                    int
}

func (b *Board) pushHistory(m Move) {
	h := History{
		previousSquareOccupant: b.PieceAt(m.to),
		move:                   m,
		castle:                 b.castle,
		ep:                     b.ep,
		hply:                   b.hply,
		ply:                    b.ply,
	}
	b.history = append(b.history, h)
}

func (b *Board) popHistory() {
	b.history[len(b.history)-1] = History{}
	b.history = b.history[:len(b.history)-1]
}