package board

import "testing"

func TestPieceAtNewBoard(t *testing.T) {
	var tests = []struct {
		sq       int
		expected int
	}{
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
	}

	b := NewBoard()

	for _, tt := range tests {
		received := b.PieceAt(tt.sq)
		if received != tt.expected {
			t.Errorf("received %d, expected %d", received, tt.expected)
		}
	}
}
