package search

import (
	"github.com/connorryanbaker/bakaMitai/board"
	"testing"
)

// TODO: test picking move in lost position
// may benefit from switching to negamax eval calls in root
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
		l := Search(&tt.b, 2)
		m := l[0]
		if m.To != tt.expectedTo {
			tt.b.Print()
			t.Errorf("received unexpected to; expected: %s, received: %s, %v", board.SQ_NUM_TO_NAME[tt.expectedTo], board.SQ_NUM_TO_NAME[m.To], l)
		}
		if m.From != tt.expectedFrom {
			t.Errorf("received unexpected from; expected: %s, received: %s", board.SQ_NUM_TO_NAME[tt.expectedFrom], board.SQ_NUM_TO_NAME[m.From])
		}
	}
}

func TestFindMateInTwo(t *testing.T) {
	var tests = []struct {
		b board.Board
		l []board.Move
	}{
		{
			board.FromFENString("7k/P7/5K2/8/8/8/8/8 w - - 0 1"),
			[]board.Move{
				{board.IF6, board.IG6, false, false, false, false, board.WHITE_KING, false},
				{board.IH8, board.IG8, false, false, false, false, board.BLACK_KING, false},
				{board.IA7, board.IA8, false, false, false, true, board.WHITE_QUEEN, false},
			},
		},
		{
			board.FromFENString("r2qkb1r/pp2nppp/3p4/2pNN1B1/2BnP3/3P4/PPP2PPP/R2bK2R w KQkq - 1 0"),
			[]board.Move{
				{board.ID5, board.IF6, false, false, false, false, board.WHITE_KNIGHT, false},
				{board.IG7, board.IF6, true, false, false, false, board.BLACK_PAWN, false},
				{board.IC4, board.IF7, true, false, false, false, board.WHITE_BISHOP, false},
			},
		},
	}

	for _, tt := range tests {
		l := Search(&tt.b, 3)
		for i, _ := range l {
			if tt.l[i].To != l[i].To {
				t.Errorf("line[%d] to; received: %s, expected: %s", i, board.SQ_NUM_TO_NAME[l[i].To], board.SQ_NUM_TO_NAME[tt.l[i].To])
				t.Errorf("%v", l)
			}
			if tt.l[i].From != l[i].From {
				t.Errorf("line[%d] from; received: %s, expected: %s", i, board.SQ_NUM_TO_NAME[l[i].From], board.SQ_NUM_TO_NAME[tt.l[i].From])
				t.Errorf("%v", l)
			}
		}
	}
}
