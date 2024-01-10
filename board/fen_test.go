package board

import "testing"

type expectation struct {
	sq    int
	piece int
}

func TestFromFENString(t *testing.T) {
	var tests = []struct {
		fen          string
		expectations []expectation
		side         int
		castle       [4]bool
		ep           *int
		hply         int
		ply          int
	}{
		{
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			[]expectation{
				{IA8, BLACK_ROOK},
				{IB8, BLACK_KNIGHT},
				{IC8, BLACK_BISHOP},
				{ID8, BLACK_QUEEN},
				{IE8, BLACK_KING},
				{IF8, BLACK_BISHOP},
				{IG8, BLACK_KNIGHT},
				{IH8, BLACK_ROOK},
				{IA7, BLACK_PAWN},
				{IB7, BLACK_PAWN},
				{IC7, BLACK_PAWN},
				{ID7, BLACK_PAWN},
				{IE7, BLACK_PAWN},
				{IF7, BLACK_PAWN},
				{IG7, BLACK_PAWN},
				{IH7, BLACK_PAWN},
				{IA6, EMPTY_SQUARE},
				{IB6, EMPTY_SQUARE},
				{IC6, EMPTY_SQUARE},
				{ID6, EMPTY_SQUARE},
				{IE6, EMPTY_SQUARE},
				{IF6, EMPTY_SQUARE},
				{IG6, EMPTY_SQUARE},
				{IH6, EMPTY_SQUARE},
				{IA5, EMPTY_SQUARE},
				{IB5, EMPTY_SQUARE},
				{IC5, EMPTY_SQUARE},
				{ID5, EMPTY_SQUARE},
				{IE5, EMPTY_SQUARE},
				{IF5, EMPTY_SQUARE},
				{IG5, EMPTY_SQUARE},
				{IH5, EMPTY_SQUARE},
				{IA4, EMPTY_SQUARE},
				{IB4, EMPTY_SQUARE},
				{IC4, EMPTY_SQUARE},
				{ID4, EMPTY_SQUARE},
				{IE4, EMPTY_SQUARE},
				{IF4, EMPTY_SQUARE},
				{IG4, EMPTY_SQUARE},
				{IH4, EMPTY_SQUARE},
				{IA3, EMPTY_SQUARE},
				{IB3, EMPTY_SQUARE},
				{IC3, EMPTY_SQUARE},
				{ID3, EMPTY_SQUARE},
				{IE3, EMPTY_SQUARE},
				{IF3, EMPTY_SQUARE},
				{IG3, EMPTY_SQUARE},
				{IH3, EMPTY_SQUARE},
				{IA2, WHITE_PAWN},
				{IB2, WHITE_PAWN},
				{IC2, WHITE_PAWN},
				{ID2, WHITE_PAWN},
				{IE2, WHITE_PAWN},
				{IF2, WHITE_PAWN},
				{IG2, WHITE_PAWN},
				{IH2, WHITE_PAWN},
				{IA1, WHITE_ROOK},
				{IB1, WHITE_KNIGHT},
				{IC1, WHITE_BISHOP},
				{ID1, WHITE_QUEEN},
				{IE1, WHITE_KING},
				{IF1, WHITE_BISHOP},
				{IG1, WHITE_KNIGHT},
				{IH1, WHITE_ROOK},
			},
			0,
			[4]bool{true, true, true, true},
			nil,
			0,
			0,
		},
	}

	// TODO: more tests obviously

	for _, tt := range tests {
		b := FromFENString(tt.fen)
		for _, expectation := range tt.expectations {
			received := b.PieceAt(expectation.sq)
			if received != expectation.piece {
				t.Errorf("sq: %d, received %d, expected %d", expectation.sq, received, expectation.piece)
			}
		}
		if b.side != tt.side {
			t.Errorf("side - received %d, expected %d", b.side, tt.side)
		}
		if b.castle != tt.castle {
			t.Errorf("castle - received %v, expected %v", b.castle, tt.castle)
		}
		if b.ep != tt.ep {
			t.Errorf("ep - received %p, expected %p", b.ep, tt.ep)
		}
		if b.hply != tt.hply {
			t.Errorf("hply - received %d, expected %d", b.hply, tt.hply)
		}
		if b.ply != tt.ply {
			t.Errorf("ply - received %d, expected %d", b.ply, tt.ply)
		}
	}
}
