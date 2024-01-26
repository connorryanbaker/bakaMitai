package eval

import (
	"github.com/connorryanbaker/engine/board"
	"testing"
)

func TestEvalPieceSquaresInitialPositionEqual(t *testing.T) {
	var tests = []struct {
		b board.Board
		e float64
		d string
	}{
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 1 1"),
			float64(0),
			"opening position",
		},
	}

	for _, tt := range tests {
		e := evalPieceSquares(tt.b)
		if e != tt.e {
			t.Errorf("%s; expected: %f, received: %f", tt.d, tt.e, e)
		}
	}
}

func TestKnightDevelopmentBalance(t *testing.T) {
	b1ND5 := board.FromFENString("rnbqkbnr/pppppppp/8/3N4/8/8/PPPPPPPP/R1BQKBNR w KQkq - 0 1")
	b2NC3NF3 := board.FromFENString("rnbqkbnr/pppppppp/8/8/8/2N2N2/PPPPPPPP/R1BQKB1R w KQkq - 0 1")

	e1 := evalPieceSquares(b1ND5)
	e2 := evalPieceSquares(b2NC3NF3)

	if e1 > e2 {
		t.Errorf("One knight on D5: %f, Two knights developed C3 F3: %f", e1, e2)
	}
}
