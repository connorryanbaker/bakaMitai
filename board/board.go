package board

type Board struct {
	pieces [120]int
	castle [4]bool
	ep     *int
	side   int
	hply   int
	ply    int
	// todo:
	// move hash for position comparison
	// https://www.chessprogramming.org/Zobrist_Hashing
	// fifty move rule
	// recognize check, checkmate, stalemate
	// move history: boardstate struct w/ hash, castling, ep, move, incheck?, repetitions
	// make move fn to maintain all this state
	// unmake move fn
	// fen parsing to load positions
	// move generation perhaps separate module
}

func NewBoard() Board {
	return Board{
		INIT_PIECES,
		INIT_CASTLE,
		nil,
		WHITE,
		0,
		0,
	}
}

func (b Board) PieceAt(idx int) int {
	return b.pieces[MAILBOX_64[idx]]
}

func (b Board) PieceFromConvertedIdx(idx int) int { // TODO: better name
	return b.pieces[idx]
}

func emptyPiecesArray() [120]int {
	p := [120]int{}
	for i := 0; i < 120; i++ {
		p[i] = OFF_BOARD
	}
	for i := 0; i < 64; i++ {
		p[MAILBOX_64[i]] = EMPTY_SQUARE
	}
	return p
}
