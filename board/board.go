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
	hashSeed     hash
	legalMoves   []Move
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
		make([]History, 0),
		newHashSeed(),
		nil,
	}
}

// todo: these should be public fields
func (b Board) History() []History {
	return b.history
}

func (b Board) Side() int {
	return b.side
}

func (b Board) Pieces() map[int][]int {
	return b.pieceSquares
}

func (b Board) isEPCapture(m Move) bool {
	if b.ep == nil {
		return false
	}
	mp := b.PieceAt(m.From)

	if mp != WHITE_PAWN && mp != BLACK_PAWN {
		return false
	}

	return m.To == *b.ep
}

func (b *Board) UnmakeMove() {
	if len(b.history) == 0 {
		return
	}

	h := b.history[len(b.history)-1]
	m := h.move
	p := b.PieceAt(m.To)
	if m.Promote {
		if b.side == BLACK {
			p = WHITE_PAWN
		} else {
			p = BLACK_PAWN
		}
	}
	b.pieces[m.From] = p
	b.pieces[m.To] = h.previousSquareOccupant
	if m.CastleKingside {
		if b.side == BLACK {
			b.pieces[IF1] = EMPTY_SQUARE
			b.pieces[IH1] = WHITE_ROOK
		} else {
			b.pieces[IF8] = EMPTY_SQUARE
			b.pieces[IH8] = BLACK_ROOK
		}
	}
	if m.CastleQueenside {
		if b.side == BLACK {
			b.pieces[ID1] = EMPTY_SQUARE
			b.pieces[IA1] = WHITE_ROOK
		} else {
			b.pieces[ID8] = EMPTY_SQUARE
			b.pieces[IA8] = BLACK_ROOK
		}
	}
	if m.Capture && h.ep != nil && m.To == *h.ep && (p == WHITE_PAWN || p == BLACK_PAWN) {
		if b.side == BLACK {
			b.pieces[m.To+10] = BLACK_PAWN
		} else {
			b.pieces[m.To-10] = WHITE_PAWN
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
	movingPiece := b.PieceAt(m.From)
	pieceAtDestSq := b.PieceAt(m.To)
	if b.PieceColor(movingPiece) != b.side {
		return false
	}

	if m.CastleKingside {
		b.castleKingside(m)
		b.legalMoves = nil
		return true
	}

	if m.CastleQueenside {
		b.castleQueenside(m)
		b.legalMoves = nil
		return true
	}

	if m.Promote {
		return b.handlePromotion(m)
	}

	if b.isEPCapture(m) {
		return b.handleEPCapture(m)
	}

	if m.Capture {
		return b.handleCapture(m)
	}

	// "quiet move"

	b.pushHistory(m)
	b.pieces[m.To] = movingPiece
	b.pieces[m.From] = EMPTY_SQUARE
	b.updatePieceSquares()
	if b.InCheck(b.side) {
		b.popHistory()
		b.pieces[m.From] = movingPiece
		b.pieces[m.To] = pieceAtDestSq
		b.updatePieceSquares()
		return false
	}

	if m.DoublePawnPush {
		if b.side == WHITE {
			s := m.To + 10
			b.ep = &s
		} else {
			s := m.To - 10
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
	b.legalMoves = nil
	return true
}

func (b *Board) handleCapture(m Move) bool {
	movingPiece := b.PieceAt(m.From)
	capturedPiece := b.PieceAt(m.To)
	b.pushHistory(m)
	b.pieces[m.To] = movingPiece
	b.pieces[m.From] = EMPTY_SQUARE
	b.updatePieceSquares()
	if b.InCheck(b.side) {
		b.popHistory()
		b.pieces[m.To] = capturedPiece
		b.pieces[m.From] = movingPiece
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
	b.legalMoves = nil
	return true
}

func (b *Board) handlePromotion(m Move) bool {
	prevSq := b.PieceAt(m.To)
	movingPiece := b.PieceAt(m.From)
	b.pushHistory(m)
	b.pieces[m.To] = m.PromotionPiece
	b.pieces[m.From] = EMPTY_SQUARE
	b.updatePieceSquares()
	if b.InCheck(b.side) {
		b.popHistory()
		b.pieces[m.To] = prevSq
		b.pieces[m.From] = movingPiece
		b.updatePieceSquares()
		return false
	}
	b.ep = nil
	b.hply = 0
	if b.side == BLACK {
		b.ply += 1
	}
	b.side ^= 1
	b.legalMoves = nil
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
		if m.From == IA1 {
			b.castle[1] = false
			return
		}
		if m.From == IH1 {
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
		if m.From == IA8 {
			b.castle[3] = false
			return
		}
		if m.From == IH8 {
			b.castle[2] = false
		}
	}
}

func (b *Board) handleEPCapture(m Move) bool {
	b.pushHistory(m)
	if b.side == WHITE {
		b.pieces[m.To] = WHITE_PAWN
		b.pieces[m.From] = EMPTY_SQUARE
		b.pieces[m.To+10] = EMPTY_SQUARE
		b.updatePieceSquares()
		if b.InCheck(b.side) {
			b.popHistory()
			b.pieces[m.To+10] = BLACK_PAWN
			b.pieces[m.To] = EMPTY_SQUARE
			b.pieces[m.From] = WHITE_PAWN
			b.updatePieceSquares()
			return false
		}
	} else {
		b.pieces[m.To] = BLACK_PAWN
		b.pieces[m.From] = EMPTY_SQUARE
		b.pieces[m.To-10] = EMPTY_SQUARE
		b.updatePieceSquares()
		if b.InCheck(b.side) {
			b.popHistory()
			b.pieces[m.To-10] = WHITE_PAWN
			b.pieces[m.To] = EMPTY_SQUARE
			b.pieces[m.From] = BLACK_PAWN
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
	b.legalMoves = nil
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

func (b Board) FiftyMoveDraw() bool {
	return b.hply == 100
}

func (b Board) ThreefoldRepetition() bool {
	currentHash := b.Hash()
	count := 1
	for _, v := range b.history {
		if v.hash == currentHash {
			count += 1
		}
	}
	return count >= 3
}

func (b Board) InsufficientMaterial() bool {
	return len(b.pieceSquares) == 2
}

func (b Board) Drawn() bool {
	return b.InsufficientMaterial() || b.ThreefoldRepetition() || b.FiftyMoveDraw() || b.Stalemate()
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
