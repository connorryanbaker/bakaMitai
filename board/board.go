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

// TODO: consolidate usage here or at least set a standard

func (b Board) PieceAt(idx int) int {
	return b.pieces[MAILBOX_64[idx]]
}

func (b Board) PieceAtFromConst(idx int) int {
  return b.pieces[idx]
}

func emptyPiecesArray() [120]int {
  p := [120]int{}
  for i := 0; i < 120; i++ {
    if i < 21 || i > 98 {
      p[i] = OFF_BOARD
    } else {
      p[i] = EMPTY_SQUARE
    }
  }
  return p
}
