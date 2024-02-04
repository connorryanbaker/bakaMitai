package eval

import (
	"github.com/connorryanbaker/bakaMitai/board"

	"testing"
)

func TestDoubledPawnPenalties(t *testing.T) {
	var tests = []struct {
		b board.Board
		e int
	}{
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			0,
		},
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/5P2/PPPP1PPP/RNBQKBNR w KQkq - 0 1"),
			-1,
		},
		{
			board.FromFENString("rnbqkbnr/pppp1ppp/3p4/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			1,
		},
		{
			board.FromFENString("rnbqkbnr/pppp1ppp/3p4/8/8/6P1/PPPPPPP1/RNBQKBNR w KQkq - 0 1"),
			0,
		},
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/6P1/6P1/1PPPPPP1/RNBQKBNR w KQkq - 0 1"),
			-2,
		},
		{
			board.FromFENString("rnbqkbnr/pppp2pp/3p2p1/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			2,
		},
	}

	for _, tt := range tests {
		pcm := pawnCountMap(tt.b)
		v := doubledPawnPenalties(pcm)
		if tt.e != v {
			t.Errorf("Unexpected doubled pawn penalty eval; received: %d, expected: %d", v, tt.e)
			tt.b.Print()
		}
	}
}

func TestIsolatedPawnPenalties(t *testing.T) {
	var tests = []struct {
		b board.Board
		e int
	}{
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			0,
		},
		{
			board.FromFENString("rnbqkbnr/pppppppp/8/8/8/8/P1PPPPPP/RNBQKBNR w KQkq - 0 1"),
			-1,
		},
		{
			board.FromFENString("rnbqkbnr/pp1p1ppp/8/8/8/8/PP1PPPPP/RNBQKBNR w KQkq - 0 1"),
			1,
		},
		{
			board.FromFENString("rnbqkbnr/pp1p1p1p/8/8/8/8/PP1PPPPP/RNBQKBNR w KQkq - 0 1"),
			3,
		},
	}
	for _, tt := range tests {
		pcm := pawnCountMap(tt.b)
		v := isolatedPawnPenalties(pcm)
		if tt.e != v {
			t.Errorf("Unexpected isolated pawn penalty eval; received: %d, expected: %d", v, tt.e)
			tt.b.Print()
		}
	}
}
