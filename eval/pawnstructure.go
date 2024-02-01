package eval

import "github.com/connorryanbaker/engine/board"

func evalPawnStructure(b board.Board) float64 {
	// return doubledPawnPenalties(b) + isolatedPawnPenalties(b)
	return float64(doubledPawnPenalties(b))
}

func doubledPawnPenalties(b board.Board) int {
	wdp := 0
	bdp := 0
	for i := 0; i < 8; i++ {
		r := board.MAILBOX_64[i]
		wpc := 0
		bpc := 0
		for ; r <= board.IH8; r += 10 {
			if b.PieceAt(r) == board.WHITE_PAWN {
				wpc += 1
			}
			if b.PieceAt(r) == board.BLACK_PAWN {
				bpc += 1
			}
		}
		if wpc > 1 {
			wdp += wpc
		}
		if bpc > 1 {
			bdp += bpc
		}
	}
	return bdp - wdp
}
