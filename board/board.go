package board

type Board struct {
	pieces       [120]int
	Castle       [4]bool // wck,wcq,bck,bcq
	Ep           *int
	Side         int
	Hply         int
	Ply          int
	PieceSquares map[int][]int
	History      []History
	hashSeed     hash
}

func NewBoard() Board {
	b := Board{
		INIT_PIECES,
		INIT_CASTLE,
		nil,
		WHITE,
		0,
		0,
		INIT_PIECE_SQUARES, // we'll see if this works
		make([]History, 0),
		newHashSeed(),
	}
	return b
}

func (b Board) isEPCapture(m Move) bool {
	if b.Ep == nil {
		return false
	}
	mp := b.PieceAt(m.From)

	if mp != WHITE_PAWN && mp != BLACK_PAWN {
		return false
	}

	return m.To == *b.Ep
}

func (b *Board) UnmakeMove() {
	h := b.History[len(b.History)-1]
	m := h.Move
	p := b.PieceAt(m.To)
	if m.Promote {
		if b.Side == BLACK {
			p = WHITE_PAWN
		} else {
			p = BLACK_PAWN
		}
	}
	b.pieces[m.From] = p
	b.pieces[m.To] = h.previousSquareOccupant
	if m.CastleKingside {
		if b.Side == BLACK {
			b.pieces[IF1] = EMPTY_SQUARE
			b.pieces[IH1] = WHITE_ROOK
		} else {
			b.pieces[IF8] = EMPTY_SQUARE
			b.pieces[IH8] = BLACK_ROOK
		}
	}
	if m.CastleQueenside {
		if b.Side == BLACK {
			b.pieces[ID1] = EMPTY_SQUARE
			b.pieces[IA1] = WHITE_ROOK
		} else {
			b.pieces[ID8] = EMPTY_SQUARE
			b.pieces[IA8] = BLACK_ROOK
		}
	}
	if m.Capture && h.ep != nil && m.To == *h.ep && (p == WHITE_PAWN || p == BLACK_PAWN) {
		if b.Side == BLACK {
			b.pieces[m.To+10] = BLACK_PAWN
		} else {
			b.pieces[m.To-10] = WHITE_PAWN
		}
	}
	b.Ep = h.ep
	b.Castle = h.castle
	b.Hply = h.hply
	b.Ply = h.ply
	b.Side ^= 1
	b.popHistory()
	b.updatePieceSquares()
}

func (b *Board) MakeMove(m Move) bool {
	movingPiece := b.PieceAt(m.From)
	if b.PieceColor(movingPiece) != b.Side {
		panic(1)
	}

	if m.CastleKingside {
		b.castleKingside(m)
		return true
	}

	if m.CastleQueenside {
		b.castleQueenside(m)
		return true
	}

	if m.Promote {
		return b.handleBBPromotion(m)
	}

	if b.isEPCapture(m) {
		return b.handleBBEPCapture(m)
	}

	if m.Capture {
		return b.handleBBCapture(m)
	}

	// "quiet move"

	b.pushHistory(m)
	b.pieces[m.To] = movingPiece
	b.pieces[m.From] = EMPTY_SQUARE
	b.updatePieceSquares()

	if m.DoublePawnPush {
		if b.Side == WHITE {
			s := m.To + 10
			b.Ep = &s
		} else {
			s := m.To - 10
			b.Ep = &s
		}
	} else {
		b.Ep = nil
	}

	b.updateCastlePermissions(m, movingPiece, EMPTY_SQUARE)
	if movingPiece == WHITE_PAWN || movingPiece == BLACK_PAWN {
		b.Hply = 0
	} else {
		b.Hply += 1
	}
	if b.Side == BLACK {
		b.Ply += 1
	}
	b.Side ^= 1
	return true
}

func (b *Board) handleBBCapture(m Move) bool {
	movingPiece := b.PieceAt(m.From)
	capturedPiece := b.PieceAt(m.To)
	b.pushHistory(m)
	b.pieces[m.To] = movingPiece
	b.pieces[m.From] = EMPTY_SQUARE
	b.updatePieceSquares()
	b.updateCastlePermissions(m, movingPiece, capturedPiece)
	b.Ep = nil
	b.Hply = 0
	if b.Side == BLACK {
		b.Ply += 1
	}
	b.Side ^= 1
	return true
}

func (b *Board) handleBBPromotion(m Move) bool {
	b.pushHistory(m)
	b.pieces[m.To] = m.PromotionPiece
	b.pieces[m.From] = EMPTY_SQUARE
	b.updatePieceSquares()
	b.Ep = nil
	b.Hply = 0
	if b.Side == BLACK {
		b.Ply += 1
	}
	b.Side ^= 1
	return true
}

func (b *Board) updateCastleMovingPiece(m Move, p int) {
	if b.Side == WHITE {
		if !b.Castle[0] && !b.Castle[1] {
			return
		}
		if p != WHITE_ROOK && p != WHITE_KING {
			return
		}
		if p == WHITE_KING {
			b.Castle[0] = false
			b.Castle[1] = false
			return
		}
		if m.From == IA1 {
			b.Castle[1] = false
			return
		}
		if m.From == IH1 {
			b.Castle[0] = false
		}
	} else {
		if !b.Castle[2] && !b.Castle[3] {
			return
		}
		if p != BLACK_ROOK && p != BLACK_KING {
			return
		}
		if p == BLACK_KING {
			b.Castle[2] = false
			b.Castle[3] = false
			return
		}
		if m.From == IA8 {
			b.Castle[3] = false
			return
		}
		if m.From == IH8 {
			b.Castle[2] = false
		}
	}
}

func (b *Board) updateCastleCapturedPiece(m Move, p int) {
	if !m.Capture {
		return
	}

	if b.Side == WHITE {
		if p != BLACK_ROOK {
			return
		}
		if m.To == IA8 {
			b.Castle[3] = false
		}
		if m.To == IH8 {
			b.Castle[2] = false
		}
	} else {
		if p != WHITE_ROOK {
			return
		}
		if m.To == IA1 {
			b.Castle[1] = false
		}
		if m.To == IH1 {
			b.Castle[0] = false
		}
	}
}

func (b *Board) updateCastlePermissions(m Move, mp, cp int) {
	b.updateCastleMovingPiece(m, mp)
	b.updateCastleCapturedPiece(m, cp)
}

func (b *Board) handleBBEPCapture(m Move) bool {
	b.pushHistory(m)
	if b.Side == WHITE {
		b.pieces[m.To] = WHITE_PAWN
		b.pieces[m.From] = EMPTY_SQUARE
		b.pieces[m.To+10] = EMPTY_SQUARE
		b.updatePieceSquares()
	} else {
		b.pieces[m.To] = BLACK_PAWN
		b.pieces[m.From] = EMPTY_SQUARE
		b.pieces[m.To-10] = EMPTY_SQUARE
		b.updatePieceSquares()
	}
	b.Hply = 0
	b.Ep = nil
	if b.Side == BLACK {
		b.Ply += 1
	}
	b.Side ^= 1
	return true
}

func (b *Board) castleKingside(m Move) {
	b.pushHistory(m)
	if b.Side == WHITE {
		b.pieces[IG1] = WHITE_KING
		b.pieces[IF1] = WHITE_ROOK
		b.pieces[IE1] = EMPTY_SQUARE
		b.pieces[IH1] = EMPTY_SQUARE
		b.Castle[0] = false
		b.Castle[1] = false
	} else {
		b.pieces[IG8] = BLACK_KING
		b.pieces[IF8] = BLACK_ROOK
		b.pieces[IE8] = EMPTY_SQUARE
		b.pieces[IH8] = EMPTY_SQUARE
		b.Castle[2] = false
		b.Castle[3] = false
	}
	b.updatePieceSquares()
	b.Ep = nil
	b.Hply += 1
	if b.Side == BLACK {
		b.Ply += 1
	}
	b.Side ^= 1
}

func (b *Board) castleQueenside(m Move) {
	b.pushHistory(m)
	if b.Side == WHITE {
		b.pieces[IC1] = WHITE_KING
		b.pieces[ID1] = WHITE_ROOK
		b.pieces[IE1] = EMPTY_SQUARE
		b.pieces[IA1] = EMPTY_SQUARE
		b.Castle[0] = false
		b.Castle[1] = false
	} else {
		b.pieces[IC8] = BLACK_KING
		b.pieces[ID8] = BLACK_ROOK
		b.pieces[IE8] = EMPTY_SQUARE
		b.pieces[IA8] = EMPTY_SQUARE
		b.Castle[2] = false
		b.Castle[3] = false
	}
	b.updatePieceSquares()
	b.Ep = nil
	b.Hply += 1
	if b.Side == BLACK {
		b.Ply += 1
	}
	b.Side ^= 1
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
	// TODO: cleanup
	bb := b.newBitboard()
	return b.inCheck(*bb)
}

func (b Board) Checkmate() bool {
	if !b.InCheck(b.Side) {
		return false
	}
	return len(b.GenerateBitboardMoves()) == 0
}

func (b Board) Stalemate() bool {
	if b.InCheck(b.Side) {
		return false
	}
	return len(b.GenerateBitboardMoves()) == 0
}

func (b Board) FiftyMoveDraw() bool {
	return b.Hply == 100
}

func (b Board) ThreefoldRepetition() bool {
	currentHash := b.Hash()
	count := 1
	for _, v := range b.History {
		if v.Hash == currentHash {
			count += 1
		}
	}
	return count >= 3
}

func (b Board) InsufficientMaterial() bool {
	if len(b.PieceSquares) == 2 {
		return true
	}
	if len(b.PieceSquares) == 3 {
		if _, ok := b.PieceSquares[WHITE_KNIGHT]; ok {
			return true
		}
		if _, ok := b.PieceSquares[BLACK_KNIGHT]; ok {
			return true
		}
		if p, ok := b.PieceSquares[WHITE_BISHOP]; ok {
			return len(p) == 1
		}
		if p, ok := b.PieceSquares[BLACK_BISHOP]; ok {
			return len(p) == 1
		}
	}
	return false
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
	b.PieceSquares = nps
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
