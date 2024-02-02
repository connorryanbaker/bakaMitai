package eval

import "github.com/connorryanbaker/bakaMitai/board"

func evalMobility(b board.Board) float64 {
	var wt float64
	var bt float64
	if b.Side == board.WHITE {
		wt = pieceMoveCount(b.LegalMoves(), b)
		b.Side ^= 1
		bt = pieceMoveCount(b.LegalMoves(), b)
	} else {
		bt = pieceMoveCount(b.LegalMoves(), b)
		b.Side ^= 1
		wt = pieceMoveCount(b.LegalMoves(), b)
	}
	return wt - bt
}

func pieceMoveCount(m []board.Move, b board.Board) float64 {
	var total float64
	for _, m := range m {
		p := b.PieceAt(m.From)
		switch p {
		case board.WHITE_KNIGHT, board.BLACK_KNIGHT, board.WHITE_QUEEN, board.BLACK_QUEEN:
			total += 1.5
		case board.WHITE_BISHOP, board.BLACK_BISHOP:
			total += 2.0
		case board.WHITE_ROOK, board.BLACK_ROOK:
			total += 3.0
		}
	}
	return total
}
