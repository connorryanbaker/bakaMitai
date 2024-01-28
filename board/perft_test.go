package board

import (
	"flag"
	"testing"
)

var testPerft = flag.Bool("perft", false, "run perft tests")

func TestPerft(t *testing.T) {
	if *testPerft != true {
		t.Skip("Skipping perft tests")
	}

	var tests = []struct {
		depth    int
		board    Board
		expected uint64
	}{
		{
			1,
			NewBoard(),
			20,
		},
		{
			2,
			NewBoard(),
			400,
		},
		{
			3,
			NewBoard(),
			8902,
		},
		{
			4,
			NewBoard(),
			197281,
		},
		{
			5,
			NewBoard(),
			4865609,
		},
		//{
		//	6,
		//	NewBoard(),
		//	119060324,
		//},
		// after profiling, looks like move generation
		// is a bottleneck - revisit after refactor
	}

	for _, tt := range tests {
		result := perft(&tt.board, tt.depth)
		if result != tt.expected {
			t.Errorf("Perft Err! Depth: %d, Expected: %d, Received: %d", tt.depth, tt.expected, result)
		} else {
			t.Logf("Passed depth %d\n", tt.depth)
		}
	}
}
