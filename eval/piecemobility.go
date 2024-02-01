package eval

import "github.com/connorryanbaker/engine/board"

func evalMobility(b board.Board) float64 {
	wt := 0
	bt := 0
	if b.Side == board.WHITE {
		wt = pieceMoveCount(b.LegalMoves(), b)
		b.Side ^= 1
		bt = pieceMoveCount(b.LegalMoves(), b)
	} else {
		bt = pieceMoveCount(b.LegalMoves(), b)
		b.Side ^= 1
		wt = pieceMoveCount(b.LegalMoves(), b)
	}
	return float64(wt - bt)
}

func pieceMoveCount(m []board.Move, b board.Board) int {
	total := 0
	for _, m := range m {
		p := b.PieceAt(m.From)
		switch p {
		case board.WHITE_KNIGHT, board.BLACK_KNIGHT:
			total += 1
		case board.WHITE_BISHOP, board.BLACK_BISHOP:
			total += 2
		case board.WHITE_ROOK, board.BLACK_ROOK, board.WHITE_QUEEN, board.BLACK_QUEEN:
			total += 3
		}
	}
	return total
}
