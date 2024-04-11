package search

import (
	"github.com/connorryanbaker/bakaMitai/board"

	"testing"
	"time"
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
		pv := NewLine(4)
		l := Search(&tt.b, 4, &pv, time.Now().Add(time.Second*30))
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
				{board.IA7, board.IA8, false, false, false, true, board.WHITE_ROOK, false},
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
		// TODO: debug
		// {
		// 	board.FromFENString("2k1r2r/ppp3p1/3b4/3pq2b/7p/2NP1P2/PPP2Q1P/R5RK b - - 0 1"),
		// 	[]board.Move{
		// 		{board.IH5, board.IF3, true, false, false, false, board.BLACK_BISHOP, false},
		// 		{board.IF2, board.IF3, true, false, false, false, board.WHITE_QUEEN, false},
		// 		{board.IE5, board.IH2, true, false, false, false, board.BLACK_QUEEN, false},
		// 	},
		// },
		// {
		// 	board.FromFENString("3k4/1p3Bp1/p5r1/2b5/P3P1N1/5Pp1/1P1r4/2R4K b - - 0 1"),
		// 	[]board.Move{
		// 		{board.ID2, board.IH2, false, false, false, false, board.BLACK_ROOK, false},
		// 		{board.IG4, board.IH2, true, false, false, false, board.WHITE_KNIGHT, false},
		// 		{board.IG3, board.IG2, false, false, false, false, board.BLACK_PAWN, false},
		// 	},
		// },
	}

	for _, tt := range tests {
		pv := NewLine(4)
		tt.b.Print()
		l := Search(&tt.b, 4, &pv, time.Now().Add(time.Second*30))
		for i := 0; i < 3; i++ {
			if !board.EqualMoves(tt.l[i], l[i]) {
				t.Errorf("line[%d] to; received: %s, expected: %s", i, board.SQ_NUM_TO_NAME[l[i].To], board.SQ_NUM_TO_NAME[tt.l[i].To])
				t.Errorf("line[%d] from; received: %s, expected: %s", i, board.SQ_NUM_TO_NAME[l[i].From], board.SQ_NUM_TO_NAME[tt.l[i].From])
				t.Errorf("received line: %v", l)
				t.Errorf("expected line: %v", tt.l)
				tt.b.Print()
			}
		}
	}
}

func TestFindMateInThree(t *testing.T) {
	var tests = []struct {
		b board.Board
		l []board.Move
	}{
		{
			board.FromFENString("r1b1kb1r/pppp1ppp/5q2/4n3/3KP3/2N3PN/PPP4P/R1BQ1B1R b kq - 0 1"),
			[]board.Move{
				{board.IF8, board.IC5, false, false, false, false, board.BLACK_BISHOP, false},
				{board.ID4, board.IC5, true, false, false, false, board.WHITE_KING, false},
				{board.IF6, board.IB6, false, false, false, false, board.BLACK_QUEEN, false},
				{board.IC5, board.ID5, false, false, false, false, board.WHITE_KING, false},
				{board.IB6, board.ID6, false, false, false, false, board.BLACK_QUEEN, false},
			},
		},
	}
	for _, tt := range tests {
		pv := NewLine(6)
		l := Search(&tt.b, 6, &pv, time.Now().Add(time.Second*60))
		for i := 0; i < 5; i++ {
			if !board.EqualMoves(tt.l[i], l[i]) {
				t.Errorf("line[%d] to; received: %s, expected: %s", i, board.SQ_NUM_TO_NAME[l[i].To], board.SQ_NUM_TO_NAME[tt.l[i].To])
				t.Errorf("line[%d] from; received: %s, expected: %s", i, board.SQ_NUM_TO_NAME[l[i].From], board.SQ_NUM_TO_NAME[tt.l[i].From])
				t.Errorf("received line: %v", l)
				t.Errorf("expected line: %v", tt.l)
				tt.b.Print()
			}
		}
	}
}
