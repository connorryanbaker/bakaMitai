package board

type Board struct {
	pieces [120]int
	castle [2]bool
	ep     *int
	side   int
	// todo:
	// repetition hash/counter
	// side to move
	// move hash for position comparison
	// fifty move rule
	// recognize check, checkmate, stalemate
	// move history: boardstate struct w/ hash, castling, ep, move, incheck?, repetitions
	// make move fn to maintain all this state
	// fen parsing to load positions
	// move generation perhaps separate module
}

func NewBoard() Board {
	return Board{
		INIT_PIECES,
		INIT_CASTLE,
		nil,
		WHITE,
	}
}

func (b Board) PieceAt(idx int) int {
	return b.pieces[MAILBOX_64[idx]]
}
