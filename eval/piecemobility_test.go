package eval

import (
	"github.com/connorryanbaker/bakaMitai/board"

	"testing"
)

func TestEvalMobility(t *testing.T) {
	var tests = []struct {
		b board.Board
		e float64
	}{
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			float64(0),
		},
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/5N2/PPPPPPPP/RNBQKB1R w KQkq - 0 1"),
			float64(7.5),
		},
		{
			board.FromFENString("r1bqkbnr/pppppppp/2n5/8/8/5N2/PPPPPPPP/RNBQKB1R w KQkq - 0 1"),
			float64(0),
		},
		{
			board.FromFENString("r1bqkbnr/pppppppp/2n5/8/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 0 1"),
			float64(11.5),
		},
		{
			board.FromFENString("r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 0 1"),
			float64(-6),
		},
		{
			board.FromFENString("r1bqkbnr/pppp1ppp/2n5/4p3/4P3/2N5/PPPP1PPP/R1BQKBNR w KQkq - 0 1"),
			float64(0),
		},
		{
			board.FromFENString("r1bqkbnr/pppp1ppp/2n5/4p3/4P3/2N5/PPPP1PP1/R1BQKBNR w KQkq - 0 1"),
			float64(18),
		},
		{
			board.FromFENString("r1bqkbnr/1ppp1ppp/2n5/4p3/4P3/2N5/PPPP1PP1/R1BQKBNR w KQkq - 0 1"),
			float64(-1.5),
		},
		{
			board.FromFENString("r1bqkbnr/1ppp1pp1/2n5/4p3/4P3/2N5/1PPP1PP1/R1BQKBNR w KQkq - 0 1"),
			float64(0),
		},
	}

	for _, tt := range tests {
		v := evalMobility(tt.b)
		if v != tt.e {
			tt.b.Print()
			t.Errorf("Unexpected mobility eval; received: %f, expected: %f", v, tt.e)
		}
	}
}
