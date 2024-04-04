package board

func see(sq int, b *Board) int {
	_, move := getSmallestAttackerAndMove(sq, *b)
	value := 0

	if !move.IsNull() {
		capturedPiece := b.PieceAt(move.To)
		b.MakeBBMove(move)
		value = values[capturedPiece] - see(sq, b)
		b.UnmakeMove()
	}
	return value
}

var values = map[int]int{
	WHITE_PAWN:   100,
	BLACK_PAWN:   100,
	WHITE_KNIGHT: 300,
	BLACK_KNIGHT: 300,
	WHITE_BISHOP: 350,
	BLACK_BISHOP: 350,
	WHITE_ROOK:   500,
	BLACK_ROOK:   500,
	WHITE_QUEEN:  900,
	BLACK_QUEEN:  900,
}

const pawnCaptureScore = 1
const knightCaptureScore = 2
const bishopCaptureScore = 3
const rookCaptureScore = 4
const queenCaptureScore = 5
const kingCaptureScore = 6

func getSmallestAttackerAndMove(sq int, b Board) (int, Move) {
	var moves []Move
	sqBB := BB(1 << MAILBOX_TO_BB[sq])
	bb := b.newBitboard()
	captureMask, pushMask, movesForPinned, pinnedPieces := b.pinMasks(*bb)

	empty := bb.emptySquares()
	var fpieces BB
	var opieces BB
	var knights BB
	var bishops BB
	var rooks BB
	var queens BB
	var king BB
	if b.Side == WHITE {
		fpieces = bb.whitePieces()
		opieces = bb.blackPieces()
		knights = bb.whiteknights
		bishops = bb.whitebishops
		rooks = bb.whiterooks
		queens = bb.whitequeens
		king = bb.whiteking
	} else {
		fpieces = bb.blackPieces()
		opieces = bb.whitePieces()
		knights = bb.blackknights
		bishops = bb.blackbishops
		rooks = bb.blackrooks
		queens = bb.blackqueens
		king = bb.blackking
	}
	var score int

	if b.legalPawnCapturesBB(*bb, captureMask, pinnedPieces, movesForPinned)&sqBB > 0 {
		moves = b.generateBitboardPawnMoves(*bb, captureMask, pushMask, pinnedPieces, movesForPinned)
		score = pawnCaptureScore
	} else if b.legalKnightCapturesBB(knights, opieces, captureMask, pinnedPieces)&sqBB > 0 {
		moves = b.generateBitboardKnightMoves(*bb, captureMask, pushMask, pinnedPieces)
		score = knightCaptureScore
	} else if b.legalBishopCapturesBB(bishops, fpieces, opieces, empty, captureMask, pinnedPieces, movesForPinned)&sqBB > 0 {
		moves = b.generateBitboardBishopMoves(*bb, captureMask, pushMask, pinnedPieces, movesForPinned)
		score = bishopCaptureScore
	} else if b.legalRookCapturesBB(rooks, fpieces, opieces, empty, captureMask, pinnedPieces, movesForPinned)&sqBB > 0 {
		moves = b.generateBitboardRookMoves(*bb, captureMask, pushMask, pinnedPieces, movesForPinned)
		score = rookCaptureScore
	} else if b.legalQueenCapturesBB(queens, fpieces, opieces, empty, captureMask, pinnedPieces, movesForPinned)&sqBB > 0 {
		moves = b.generateBitboardQueenMoves(*bb, captureMask, pushMask, pinnedPieces, movesForPinned)
		score = queenCaptureScore
	} else if legalKingCapturesBB(king, opieces, bb.tabooSquares(b.Side))&sqBB > 0 {
		moves = b.generateBitboardKingMoves(*bb)
		score = kingCaptureScore
	}

	return score, extractMove(sq, moves)
}

func extractMove(sq int, moves []Move) Move {
	var move Move
	for _, m := range moves {
		if m.Capture && m.To == sq {
			return m
		}
	}
	return move
}
