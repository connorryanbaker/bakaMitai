package eval

import (
	"github.com/connorryanbaker/engine/board"
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
	}

	for _, tt := range tests {
		e := Eval(tt.b)
		me := Eval(*board.Mirror(tt.b))
		if me != e*-1 {
			t.Errorf("Unexpected mirror evaluation; e: %f, me: %f", e, me)
		}
	}
}
