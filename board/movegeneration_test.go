package board

import "testing"

func TestWhitePawnMovesOpeningMoveNoCapture(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IA2, []Move{
			{IA2, IA3, false, false, false, false, 1, false},
			{IA2, IA4, false, false, false, false, 1, true},
		},
		},
		{IB2, []Move{
			{IB2, IB3, false, false, false, false, 1, false},
			{IB2, IB4, false, false, false, false, 1, true},
		},
		},
		{IC2, []Move{
			{IC2, IC3, false, false, false, false, 1, false},
			{IC2, IC4, false, false, false, false, 1, true},
		},
		},
		{ID2, []Move{
			{ID2, ID3, false, false, false, false, 1, false},
			{ID2, ID4, false, false, false, false, 1, true},
		},
		},
		{IE2, []Move{
			{IE2, IE3, false, false, false, false, 1, false},
			{IE2, IE4, false, false, false, false, 1, true},
		},
		},
		{IF2, []Move{
			{IF2, IF3, false, false, false, false, 1, false},
			{IF2, IF4, false, false, false, false, 1, true},
		},
		},
		{IG2, []Move{
			{IG2, IG3, false, false, false, false, 1, false},
			{IG2, IG4, false, false, false, false, 1, true},
		},
		},
		{IH2, []Move{
			{IH2, IH3, false, false, false, false, 1, false},
			{IH2, IH4, false, false, false, false, 1, true},
		},
		},
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

func TestBlackPawnMovesOpeningMoveNoCapture(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IA7, []Move{
			{IA7, IA6, false, false, false, false, 7, false},
			{IA7, IA5, false, false, false, false, 7, true},
		},
		},
		{IB7, []Move{
			{IB7, IB6, false, false, false, false, 7, false},
			{IB7, IB5, false, false, false, false, 7, true},
		},
		},
		{IC7, []Move{
			{IC7, IC6, false, false, false, false, 7, false},
			{IC7, IC5, false, false, false, false, 7, true},
		},
		},
		{ID7, []Move{
			{ID7, ID6, false, false, false, false, 7, false},
			{ID7, ID5, false, false, false, false, 7, true},
		},
		},
		{IE7, []Move{
			{IE7, IE6, false, false, false, false, 7, false},
			{IE7, IE5, false, false, false, false, 7, true},
		},
		},
		{IF7, []Move{
			{IF7, IF6, false, false, false, false, 7, false},
			{IF7, IF5, false, false, false, false, 7, true},
		},
		},
		{IG7, []Move{
			{IG7, IG6, false, false, false, false, 7, false},
			{IG7, IG5, false, false, false, false, 7, true},
		},
		},
		{IH7, []Move{
			{IH7, IH6, false, false, false, false, 7, false},
			{IH7, IH5, false, false, false, false, 7, true},
		},
		},
	}

	for _, tt := range tests {
		moves := b.BlackPawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %d, received: %v, expected: %v", tt.sq, moves, tt.moves)
			}
		}
	}
}
