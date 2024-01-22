package eval

import (
  "github.com/connorryanbaker/engine/board"
  "testing"
)

func TestEvalPieceSquaresInitialPosition(t *testing.T) {
  var tests = []struct{
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
