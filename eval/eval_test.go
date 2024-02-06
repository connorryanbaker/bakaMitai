package eval

import (
	"github.com/connorryanbaker/bakaMitai/board"
	"math"
	"testing"
)

func TestCheckmateEval(t *testing.T) {
	var tests = []struct {
		b board.Board
		e float64
	}{
		{
			board.FromFENString("8/8/8/8/8/5K2/6Q1/7k b - - 0 1"),
			math.MaxFloat64,
		},
		{
			board.FromFENString("8/8/8/8/8/5k2/6q1/7K w - - 0 1"),
			math.MaxFloat64 * -1,
		},
	}

	for _, tt := range tests {
		e := Eval(tt.b)
		if e != tt.e {
			t.Errorf("received unexpected eval; expected: %f, received: %f", e, tt.e)
		}
	}
}

func TestStalemateEval(t *testing.T) {
	var tests = []struct {
		b board.Board
		e float64
	}{
		{
			board.FromFENString("8/8/8/8/8/6Q1/5K2/7k b - - 0 1"),
			0,
		},
		{
			board.FromFENString("8/8/8/8/8/6q1/5k2/7K w - - 0 1"),
			0,
		},
	}

	for _, tt := range tests {
		e := Eval(tt.b)
		if e != tt.e {
			t.Errorf("received unexpected eval; expected: %f, received: %f", e, tt.e)
		}
	}
}

func TestInsufficientMaterialEval(t *testing.T) {
	var tests = []struct {
		b board.Board
		e float64
	}{
		{
			board.FromFENString("8/8/8/8/8/8/5K2/7k b - - 0 1"),
			0,
		},
	}

	for _, tt := range tests {
		e := Eval(tt.b)
		if e != tt.e {
			t.Errorf("received unexpected eval; expected: %f, received: %f", e, tt.e)
		}
	}
}

func TestMirrorEval(t *testing.T) {
	var tests = []struct {
		b board.Board
	}{
		{
			board.FromFENString("r1bqk1nr/pppp1ppp/2n5/2b1p3/2B1P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 0 1"),
		},
		{
			board.FromFENString("r1bqk2r/pppp1ppp/2n2n2/2b1p3/2B1P3/2N2N2/PPPP1PPP/R1BQK2R w KQkq - 0 1"),
		},
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPP1PPPP/RNBQKBNR w KQkq - 0 1"),
		},
		{
			board.FromFENString("rnbqkb1r/pppppppp/5n2/8/4P3/2N5/PPPP1PPP/R1BQKBNR b KQkq - 0 1"),
		},
		{
			board.FromFENString("rnbqkb1r/ppp1pppp/5n2/3p4/4P3/2N5/PPPP1PPP/R1BQKBNR w KQkq - 0 1"),
		},
		{
			board.FromFENString("rnbqkb1r/ppp1pppp/5n2/3P4/8/2N5/PPPP1PPP/R1BQKBNR b KQkq - 0 1"),
		},
	}

	for _, tt := range tests {
		e := Eval(tt.b)
		me := Eval(*board.Mirror(tt.b))
		if me != e*-1 {
			t.Errorf("Unexpected mirror evaluation; e: %f, me: %f", e, me)
		}
	}
}

func TestEval(t *testing.T) {
	var tests = []struct {
		b board.Board
		e float64
	}{
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPP1PPPP/RNBQKBNR w KQkq - 0 1"),
			float64(-0.9195),
		},
		{
			board.FromFENString("r1bqkbnr/pppppppp/8/8/8/8/PPP1PPPP/RNBQKBNR w KQkq - 0 1"),
			float64(1.9605),
		},
		{
			board.FromFENString("rnbqkb1r/pppppppp/5n2/1N6/8/8/PPPPPPPP/R1BQKBNR b KQkq - 0 1"),
			float64(-0.0285),
		},
		{
			board.FromFENString("rnbqkb1r/pppppppp/5n2/8/4P3/2N5/PPPP1PPP/R1BQKBNR b KQkq - 0 1"),
			float64(0.1675),
		},
		{
			board.FromFENString("rnbqkb1r/ppp1pppp/5n2/3p4/4P3/2N5/PPPP1PPP/R1BQKBNR w KQkq - 0 1"),
			float64(0.003),
		},
		{
			board.FromFENString("rnbqkb1r/ppp1pppp/5n2/3P4/8/2N5/PPPP1PPP/R1BQKBNR b KQkq - 0 1"),
			float64(1.095), // TODO: this should take SSE into account
		},
	}

	for _, tt := range tests {
		v := Eval(tt.b)
		if v != tt.e {
			t.Errorf("Unexpected evaluation; received: %f, expected: %f", v, tt.e)
		}
	}
}
