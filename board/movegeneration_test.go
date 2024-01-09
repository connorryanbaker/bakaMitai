package board

import "testing"

func TestWhitePawnMovesOpeningMoveNoCapture(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

  // TODO: failure here indicates need to decide on how to transact w/ squares
  // passing in 0-63 results in deltas expecting 0 - 119 and using PieceAt will
  // convert
  // need means of converting from 120 to 64
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{A2, []Move{
			  {A2, A3, false, false, false, false, 1, false},
			  {A2, A4, false, false, false, false, 1, true},
		  },
		},
		// {B2, []Move{
		// 	{B2, B3, false, false, false, false, 1, false},
		// 	{B2, B4, false, false, false, false, 1, true},
		// },
		// },
		// {C2, []Move{
		// 	{C2, C3, false, false, false, false, 1, false},
		// 	{C2, C4, false, false, false, false, 1, true},
		// },
		// },
		// {D2, []Move{
		// 	{D2, D3, false, false, false, false, 1, false},
		// 	{D2, D4, false, false, false, false, 1, true},
		// },
		// },
		// {E2, []Move{
		// 	{E2, E3, false, false, false, false, 1, false},
		// 	{E2, E4, false, false, false, false, 1, true},
		// },
		// },
		// {F2, []Move{
		// 	{F2, F3, false, false, false, false, 1, false},
		// 	{F2, F4, false, false, false, false, 1, true},
		// },
		// },
		// {G2, []Move{
		// 	{G2, G3, false, false, false, false, 1, false},
		// 	{G2, G4, false, false, false, false, 1, true},
		// },
		// },
		// {H2, []Move{
		// 	{H2, H3, false, false, false, false, 1, false},
		// 	{H2, H4, false, false, false, false, 1, true},
		// },
		// },
	}

	for _, tt := range tests {
		moves := b.WhitePawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %d, received: %v, expected: %v", tt.sq, moves, tt.moves)
			}
		}
	}
}
