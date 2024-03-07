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
		// {
		// 	5,
		// 	NewBoard(),
		// 	4865609,
		// },
		//{
		//	6,
		//	NewBoard(),
		//	119060324,
		//},
		// after isAttacked change 4282.39s is time to beat :D
		// need to use bitboards to circumvent checking legality by making
		// then unmaking moves
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

// Some great examples from here: https://gist.github.com/peterellisjones/8c46c28141c162d1d8a0f0badbc9cff9
func TestCustomFENPerft(t *testing.T) {
	var tests = []struct {
		depth    int
		board    Board
		expected uint64
	}{
		{
			1,
			FromFENString("r6r/1b2k1bq/8/8/7B/8/8/R3K2R b KQ - 3 2"),
			8,
		},
		{
			1,
			FromFENString("8/8/8/2k5/2pP4/8/B7/4K3 b - d3 0 3"),
			8,
		},
		{
			1,
			FromFENString("r1bqkbnr/pppppppp/n7/8/8/P7/1PPPPPPP/RNBQKBNR w KQkq - 2 2"),
			19,
		},
		{
			1,
			FromFENString("r3k2r/p1pp1pb1/bn2Qnp1/2qPN3/1p2P3/2N5/PPPBBPPP/R3K2R b KQkq - 3 2"),
			5,
		},
		{
			1,
			FromFENString("2kr3r/p1ppqpb1/bn2Qnp1/3PN3/1p2P3/2N5/PPPBBPPP/R3K2R b KQ - 3 2"),
			44,
		},
		{
			1,
			FromFENString("rnb2k1r/pp1Pbppp/2p5/q7/2B5/8/PPPQNnPP/RNB1K2R w KQ - 3 9"),
			39,
		},
		{
			1,
			FromFENString("2r5/3pk3/8/2P5/8/2K5/8/8 w - - 5 4"),
			9,
		},
		{
			3,
			FromFENString("rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8"),
			62379,
		},
		{
			3,
			FromFENString("r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10"),
			89890,
		},
		{
			6,
			FromFENString("3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1"),
			1134888,
		},
		{
			6,
			FromFENString("8/8/4k3/8/2p5/8/B2P2K1/8 w - - 0 1"),
			1015133,
		},
		{
			6,
			FromFENString("8/8/1k6/2b5/2pP4/8/5K2/8 b - d3 0 1"),
			1440467,
		},
		{
			6,
			FromFENString("5k2/8/8/8/8/8/8/4K2R w K - 0 1"),
			661072,
		},
		{
			6,
			FromFENString("3k4/8/8/8/8/8/8/R3K3 w Q - 0 1"),
			803711,
		},
		{
			4,
			FromFENString("r3k2r/1b4bq/8/8/8/8/7B/R3K2R w KQkq - 0 1"),
			1274206,
		},
		{
			4,
			FromFENString("r3k2r/8/3Q4/8/8/5q2/8/R3K2R b KQkq - 0 1"),
			1720476,
		},
		{
			6,
			FromFENString("2K2r2/4P3/8/8/8/8/8/3k4 w - - 0 1"),
			3821001,
		},
		{
			5,
			FromFENString("8/8/1P2K3/8/2n5/1q6/8/5k2 b - - 0 1"),
			1004658,
		},
		{
			6,
			FromFENString("4k3/1P6/8/8/8/8/K7/8 w - - 0 1"),
			217342,
		},
		{
			6,
			FromFENString("K1k5/8/P7/8/8/8/8/8 w - - 0 1"),
			2217,
		},
		{
			7,
			FromFENString("8/k1P5/8/1K6/8/8/8/8 w - - 0 1"),
			567584,
		},
		{
			4,
			FromFENString("8/8/2k5/5q2/5n2/8/5K2/8 b - - 0 1"),
			23527,
		},
	}

	for i, tt := range tests {
		result := BBperft(&tt.board, tt.depth)
		if result != tt.expected {
			t.Errorf("BBPerft Err! Depth: %d, Expected: %d, Received: %d, Test: %d", tt.depth, tt.expected, result, i+1)
			tt.board.Print()
		} else {
			t.Logf("Passed depth %d\n", tt.depth)
		}
	}
}

func TestBBPerft(t *testing.T) {
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
		// {
		// 	6,
		// 	NewBoard(),
		// 	119060324,
		// },
	}
	for _, tt := range tests {
		result := BBperft(&tt.board, tt.depth)
		if result != tt.expected {
			t.Errorf("BBPerft Err! Depth: %d, Expected: %d, Received: %d", tt.depth, tt.expected, result)
		} else {
			t.Logf("Passed depth %d\n", tt.depth)
		}
	}
}
