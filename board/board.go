package board

type Board struct {
	pieces       [120]int
	castle       [4]bool
	ep           *int
	side         int
	hply         int
	ply          int
	pieceSquares map[int][]int
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
		INIT_PIECE_SQUARES, // we'll see if this works
	}
}

func (b Board) PieceAt(idx int) int {
	return b.pieces[idx]
}

// this could / should(?) use side struct field
// taking param now for easier testing
func (b Board) InCheck(side int) bool {
	if side == WHITE {
		kingPos := b.pieceSquares[WHITE_KING][0]
		for _, sq := range b.SquaresAttackedByBlackPieces() {
			if sq == kingPos {
				return true
			}
		}
		return false
	}

	kingPos := b.pieceSquares[BLACK_KING][0]
	for _, sq := range b.SquaresAttackedByWhitePieces() {
		if sq == kingPos {
			return true
		}
	}
	return false
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
