package board

import "testing"

func TestPieceAt(t *testing.T) {
	var tests = []struct {
		sq       int
		expected int
	}{
		{0, 10},
		{7, 10},
		{63, 4},
		{56, 4},
		{1, 8},
		{6, 8},
		{57, 2},
		{62, 2},
	}

	b := NewBoard()

	for _, tt := range tests {
		received := b.PieceAt(tt.sq)
		if received != tt.expected {
			t.Errorf("received %d, expected %d", received, tt.expected)
		}
	}
}
