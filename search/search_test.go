package search

import (
	"github.com/connorryanbaker/engine/board"
	"testing"
)

func TestFindMateInOne(t *testing.T) {
	var tests = []struct {
		b            board.Board
		expectedFrom int
		expectedTo   int
	}{
		{
			board.FromFENString("4K3/8/4k3/Q7/3P4/8/8/8 w - - 0 1"),
			board.IA5,
			board.IE5,
		},
		{
			board.FromFENString("1q6/8/4nn2/8/8/3K4/8/4k3 b - - 0 1"),
			board.IB8,
			board.IB3,
		},
		{
			board.FromFENString("8/8/4NN2/8/8/5k2/8/R4K2 w - - 0 1"),
			board.IA1,
			board.IA3,
		},
		{
			board.FromFENString("8/8/8/3bb3/8/4K3/8/r3k3 b - - 0 1"),
			board.IA1,
			board.IA3,
		},
	}

	for _, tt := range tests {
		_, m := Search(&tt.b, 2)
		if m.To != tt.expectedTo {
			t.Errorf("received unexpected to; expected: %d, received: %d", tt.expectedTo, m.To)
		}
		if m.From != tt.expectedFrom {
			t.Errorf("received unexpected from; expected: %d, received: %d", tt.expectedFrom, m.From)
		}
	}
}
