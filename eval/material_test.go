package eval

import (
	"github.com/connorryanbaker/engine/board"
	"testing"
)

func TestEvalMaterial(t *testing.T) {
	var tests = []struct {
		b board.Board
		e float64
		d string
	}{
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			float64(0),
			"start position",
		},
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPP1/RNBQKBNR w KQkq - 0 1"),
			float64(-10),
			"start position, h2 pawn missing",
		},
		{
			board.FromFENString("rnbqkbnr/ppppppp1/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			float64(10),
			"start position, h7 pawn missing",
		},
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/R1BQKBNR w KQkq - 0 1"),
			float64(-30),
			"start position, NB1 missing",
		},
		{
			board.FromFENString("r1bqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			float64(30),
			"start position, NB8 missing",
		},
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RN1QKBNR w KQkq - 0 1"),
			float64(-35),
			"start position, BC1 missing",
		},
		{
			board.FromFENString("rn1qkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			float64(35),
			"start position, BC8 missing",
		},
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/1NBQKBNR w KQkq - 0 1"),
			float64(-50),
			"start position, RA1 missing",
		},
		{
			board.FromFENString("1nbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			float64(50),
			"start position, RA8 missing",
		},
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNB1KBNR w KQkq - 0 1"),
			float64(-90),
			"start position, QD1 missing",
		},
		{
			board.FromFENString("rnb1kbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			float64(90),
			"start position, QD8 missing",
		},
	}

	for _, tt := range tests {
		e := evalMaterial(tt.b)
		if e != tt.e {
			t.Errorf("%s: received unexpected evalMaterial; expected: %f, received: %f", tt.d, tt.e, e)
		}
	}
}
