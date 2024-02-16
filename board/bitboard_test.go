package board

import "testing"

func TestInitAttacks(t *testing.T) {
  b := NewBoard()
  b.Print()
  moves := b.GenerateBitboardMoves()
  for _, m := range moves {
    m.Print()
  }
	if 1 != 1 {
		t.Errorf("uh oh")
	}
}
