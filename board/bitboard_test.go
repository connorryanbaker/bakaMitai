package board

import "testing"

func TestOrdering(t *testing.T) {
	var tests = []struct {
		b  Board
		xm Move
	}{
		{
			FromFENString("rnbqkbnr/pp2pppp/3p4/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 0 1"),
			Move{
				IF1,
				IB5,
				false,
				false,
				false,
				false,
				WHITE_BISHOP,
				false,
			},
		},
		{
			FromFENString("rnbqkbnr/1p2pppp/p2p4/1Bp5/4P3/2N2N2/PPPP1PPP/R1BQK2R b KQkq - 0 1"),
			Move{
				IA6,
				IB5,
				true,
				false,
				false,
				false,
				BLACK_PAWN,
				false,
			},
		},
	}
	for _, tt := range tests {
		moves := tt.b.GenerateBitboardMoves()
		if !EqualMoves(tt.xm, moves[0]) {
			t.Errorf("expected: %s, received: %s", tt.xm.ToString(), moves[0].ToString())
		}
	}
}
