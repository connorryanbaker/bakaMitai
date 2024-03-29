package board

import "testing"

// additional tests necessary

func TestHashIsConsistent(t *testing.T) {
	b := NewBoard()
	h1 := b.Hash()
	if h1 == 0 {
		t.Errorf("board hash should not equal 0: %d", h1)
	}
	h2 := b.Hash()
	if h1 != h2 {
		t.Errorf("board hash should not change without board changes: %d %d", h1, h2)
	}
}

func TestHashChangesWithMakeMoveAndUnmakemove(t *testing.T) {
	b := NewBoard()
	h1 := b.Hash()
	b.MakeMove(Move{
		IE2,
		IE4,
		false,
		false,
		false,
		false,
		WHITE_PAWN,
		true,
	})
	if b.Hash() == h1 {
		t.Errorf("board hash should update with moves: %d %d", h1, b.Hash())
	}
	if b.History[0].Hash != h1 {
		t.Errorf("board history hash should match: %d %d", b.History[0].Hash, h1)
	}
	b.UnmakeMove()
	if b.Hash() != h1 {
		t.Errorf("board hash should match original with unmake move: %d %d", h1, b.Hash())
	}
}
