package board

import "testing"

func TestgetSmallestAttacker(tt *testing.T) {
	var tests = []struct {
		b             board
		expectedPiece int
		expectedSq    int
	}{
		{
			FromFENString("rnbqkb1r/ppp1pppp/5n2/3p4/4P3/2N5/PPPP1PPP/R1BQKBNR w KQkq - 0 1"),
			WHITE_PAWN,
			IE4,
		},
	}
}
