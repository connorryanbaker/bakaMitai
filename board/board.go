package board

type Board struct {
	pieces       [120]int
	castle       [4]bool // wck,wcq,bck,bcq
	ep           *int
	side         int
	hply         int
	ply          int
	pieceSquares map[int][]int
	history      []History
}

// todo:
// recognize checkmate, stalemate, fifty move rule [ ]
// Moves() function to expose all available semi-legal moves in position [ ]
// add move hash to history for position comparison [ ]
// https://www.chessprogramming.org/Zobrist_Hashing [ ]
// printing for debugging / ui / accepting user input, play game [ ]
// then move on to the fun stuff

func NewBoard() Board {
	return Board{
		INIT_PIECES,
		INIT_CASTLE,
		nil,
		WHITE,
		0,
		0,
		INIT_PIECE_SQUARES, // we'll see if this works
		make([]History, 1),
	}
}

func (b Board) isEPCapture(m Move) bool {
	if b.ep == nil {
		return false
	}
	mp := b.PieceAt(m.from)

	if mp != WHITE_PAWN && mp != BLACK_PAWN {
		return false
	}

	return m.to == *b.ep
}

func (b *Board) UnmakeMove() {
	if len(b.history) == 0 {
		return
	}

	h := b.history[len(b.history)-1]
	m := h.move
	p := b.PieceAt(m.to)
	if m.promote {
		if b.side == BLACK {
			p = WHITE_PAWN
		} else {
			p = BLACK_PAWN
		}
	}
	b.pieces[m.from] = p
	b.pieces[m.to] = h.previousSquareOccupant
	if m.castleKingside {
		if b.side == BLACK {
			b.pieces[IF1] = EMPTY_SQUARE
			b.pieces[IH1] = WHITE_ROOK
		} else {
			b.pieces[IF8] = EMPTY_SQUARE
			b.pieces[IH8] = BLACK_ROOK
		}
	}
	if m.castleQueenside {
		if b.side == BLACK {
			b.pieces[ID1] = EMPTY_SQUARE
			b.pieces[IA1] = WHITE_ROOK
		} else {
			b.pieces[ID8] = EMPTY_SQUARE
			b.pieces[IA8] = BLACK_ROOK
		}
	}
	if m.capture && h.ep != nil && m.to == *h.ep && (p == WHITE_PAWN || p == BLACK_PAWN) {
		if b.side == BLACK {
			b.pieces[m.to+10] = BLACK_PAWN
		} else {
			b.pieces[m.to-10] = WHITE_PAWN
		}
	}
	b.ep = h.ep
	b.castle = h.castle
	b.hply = h.hply
	b.ply = h.ply
	b.side ^= 1
	b.popHistory()
	b.updatePieceSquares()
}

func (b *Board) MakeMove(m Move) bool {
	movingPiece := b.PieceAt(m.from)
	pieceAtDestSq := b.PieceAt(m.to)
	if b.PieceColor(movingPiece) != b.side {
		return false
	}

	if m.castleKingside {
		b.castleKingside(m)
		return true
	}

	if m.castleQueenside {
		b.castleQueenside(m)
		return true
	}

	if m.promote {
		return b.handlePromotion(m)
	}

	if b.isEPCapture(m) {
		return b.handleEPCapture(m)
	}

	if m.capture {
		return b.handleCapture(m)
	}

	// "quiet move"

	b.pushHistory(m)
	b.pieces[m.to] = movingPiece
	b.pieces[m.from] = EMPTY_SQUARE
	b.updatePieceSquares()
	if b.InCheck(b.side) {
		b.popHistory()
		b.pieces[m.from] = movingPiece
		b.pieces[m.to] = pieceAtDestSq
		b.updatePieceSquares()
		return false
	}

	if m.doublePawnPush {
		if b.side == WHITE {
			s := m.to + 10
			b.ep = &s
		} else {
			s := m.to - 10
			b.ep = &s
		}
	} else {
		b.ep = nil
	}

	b.updateCastlePermissions(m, movingPiece)
	if movingPiece == WHITE_PAWN || movingPiece == BLACK_PAWN {
		b.hply = 0
	} else {
		b.hply += 1
	}
	if b.side == BLACK {
		b.ply += 1
	}
	b.side ^= 1
	return true
}

func (b *Board) handleCapture(m Move) bool {
	movingPiece := b.PieceAt(m.from)
	capturedPiece := b.PieceAt(m.to)
	b.pushHistory(m)
	b.pieces[m.to] = movingPiece
	b.pieces[m.from] = EMPTY_SQUARE
	b.updatePieceSquares()
	if b.InCheck(b.side) {
		b.popHistory()
		b.pieces[m.to] = capturedPiece
		b.pieces[m.from] = movingPiece
		b.updatePieceSquares()
		return false
	}
	b.updateCastlePermissions(m, movingPiece)
	b.ep = nil
	b.hply = 0
	if b.side == BLACK {
		b.ply += 1
	}
	b.side ^= 1
	return true
}

func (b *Board) handlePromotion(m Move) bool {
	prevSq := b.PieceAt(m.to)
	movingPiece := b.PieceAt(m.from)
	b.pushHistory(m)
	b.pieces[m.to] = m.promotionPiece
	b.pieces[m.from] = EMPTY_SQUARE
	b.updatePieceSquares()
	if b.InCheck(b.side) {
		b.popHistory()
		b.pieces[m.to] = prevSq
		b.pieces[m.from] = movingPiece
		b.updatePieceSquares()
		return false
	}
	b.ep = nil
	b.hply = 0
	if b.side == BLACK {
		b.ply += 1
	}
	b.side ^= 1
	return true
}

func (b *Board) updateCastlePermissions(m Move, p int) {
	if b.side == WHITE {
		if !b.castle[0] && !b.castle[1] {
			return
		}
		if p != WHITE_ROOK && p != WHITE_KING {
			return
		}
		if p == WHITE_KING {
			b.castle[0] = false
			b.castle[1] = false
			return
		}
		if m.from == IA1 {
			b.castle[1] = false
			return
		}
		if m.from == IH1 {
			b.castle[0] = false
		}
	} else {
		if !b.castle[2] && !b.castle[3] {
			return
		}
		if p != BLACK_ROOK && p != BLACK_KING {
			return
		}
		if p == BLACK_KING {
			b.castle[2] = false
			b.castle[3] = false
			return
		}
		if m.from == IA8 {
			b.castle[3] = false
			return
		}
		if m.from == IH8 {
			b.castle[2] = false
		}
	}
}

func (b *Board) handleEPCapture(m Move) bool {
	b.pushHistory(m)
	if b.side == WHITE {
		b.pieces[m.to] = WHITE_PAWN
		b.pieces[m.from] = EMPTY_SQUARE
		b.pieces[m.to+10] = EMPTY_SQUARE
		b.updatePieceSquares()
		if b.InCheck(b.side) {
			b.popHistory()
			b.pieces[m.to+10] = BLACK_PAWN
			b.pieces[m.to] = EMPTY_SQUARE
			b.pieces[m.from] = WHITE_PAWN
			b.updatePieceSquares()
			return false
		}
	} else {
		b.pieces[m.to] = BLACK_PAWN
		b.pieces[m.from] = EMPTY_SQUARE
		b.pieces[m.to-10] = EMPTY_SQUARE
		b.updatePieceSquares()
		if b.InCheck(b.side) {
			b.popHistory()
			b.pieces[m.to-10] = WHITE_PAWN
			b.pieces[m.to] = EMPTY_SQUARE
			b.pieces[m.from] = BLACK_PAWN
			b.updatePieceSquares()
			return false
		}
	}
	b.hply = 0
	b.ep = nil
	if b.side == BLACK {
		b.ply += 1
	}
	b.side ^= 1
	return true
}

func (b *Board) castleKingside(m Move) {
	b.pushHistory(m)
	if b.side == WHITE {
		b.pieces[IG1] = WHITE_KING
		b.pieces[IF1] = WHITE_ROOK
		b.pieces[IE1] = EMPTY_SQUARE
		b.pieces[IH1] = EMPTY_SQUARE
		b.castle[0] = false
		b.castle[1] = false
	} else {
		b.pieces[IG8] = BLACK_KING
		b.pieces[IF8] = BLACK_ROOK
		b.pieces[IE8] = EMPTY_SQUARE
		b.pieces[IH8] = EMPTY_SQUARE
		b.castle[2] = false
		b.castle[3] = false
	}
	b.updatePieceSquares()
	b.ep = nil
	b.hply += 1
	if b.side == BLACK {
		b.ply += 1
	}
	b.side ^= 1
}

func (b *Board) castleQueenside(m Move) {
	b.pushHistory(m)
	if b.side == WHITE {
		b.pieces[IC1] = WHITE_KING
		b.pieces[ID1] = WHITE_ROOK
		b.pieces[IE1] = EMPTY_SQUARE
		b.pieces[IA1] = EMPTY_SQUARE
		b.castle[0] = false
		b.castle[1] = false
	} else {
		b.pieces[IC8] = BLACK_KING
		b.pieces[ID8] = BLACK_ROOK
		b.pieces[IE8] = EMPTY_SQUARE
		b.pieces[IA8] = EMPTY_SQUARE
		b.castle[2] = false
		b.castle[3] = false
	}
	b.updatePieceSquares()
	b.ep = nil
	b.hply += 1
	if b.side == BLACK {
		b.ply += 1
	}
	b.side ^= 1
}

func (b Board) PieceColor(p int) int {
	// todo: return error if invalid, error handling
	if EMPTY_SQUARE < p && p < BLACK_PAWN {
		return WHITE
	}
	return BLACK
}

func (b Board) PieceAt(idx int) int {
	return b.pieces[idx]
}

func (b Board) InCheck(side int) bool {
	if side == WHITE {
		kingPos := b.pieceSquares[WHITE_KING][0]
		attackedSquares := toLookupMap(b.SquaresAttackedByBlackPieces())
		return attackedSquares[kingPos] == true
	}

	kingPos := b.pieceSquares[BLACK_KING][0]
	attackedSquares := toLookupMap(b.SquaresAttackedByWhitePieces())
	return attackedSquares[kingPos] == true
}

func (b Board) Checkmate() bool {
	if !b.InCheck(b.side) {
		return false
	}
	return len(b.LegalMoves()) == 0
}

func (b Board) Stalemate() bool {
	if b.InCheck(b.side) {
		return false
	}
	return len(b.LegalMoves()) == 0
}

func (b *Board) updatePieceSquares() {
	nps := make(map[int][]int)
	for i := 0; i < 64; i++ {
		idx := MAILBOX_64[i]
		if b.pieces[idx] == EMPTY_SQUARE {
			continue
		}

		p := b.pieces[idx]
		if nps[p] != nil {
			nps[p] = append(nps[p], idx)
		} else {
			nps[p] = make([]int, 1)
			nps[p][0] = idx
		}
	}
	b.pieceSquares = nps
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
