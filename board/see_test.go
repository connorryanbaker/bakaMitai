package board

import "testing"

func TestGetSmallestAttacker(t *testing.T) {
	var tests = []struct {
		b             Board
		expectedPiece int
		expectedScore int
		targetSquare  int
	}{
		{
			FromFENString("rnbqkb1r/ppp1pppp/5n2/3p4/4P3/2N5/PPPP1PPP/R1BQKBNR w KQkq - 0 1"),
			WHITE_PAWN,
			pawnCaptureScore,
			ID5,
		},
		{
			FromFENString("rnbqkb1r/ppp1pppp/5n2/3p4/2B5/2N2Q2/PPPP1PPP/R1B1K1NR w KQkq - 0 1"),
			WHITE_KNIGHT,
			knightCaptureScore,
			ID5,
		},
		{
			FromFENString("rnbqkb1r/ppp1pppp/5n2/3p4/2B5/5Q2/PPP2PPP/2BRK1NR w Kkq - 0 1"),
			WHITE_BISHOP,
			bishopCaptureScore,
			ID5,
		},
		{
			FromFENString("rnbqkb1r/ppp1pppp/5n2/3p4/8/5Q2/PPP2PPP/2BRK1NR w Kkq - 0 1"),
			WHITE_ROOK,
			rookCaptureScore,
			ID5,
		},
		{
			FromFENString("rnbqkb1r/ppp1pppp/5n2/3p4/8/5Q2/PPP2PPP/2B1K1NR w KAkq - 0 1"),
			WHITE_QUEEN,
			queenCaptureScore,
			ID5,
		},
		{
			FromFENString("rnbqkb1r/ppp1pppp/5n2/8/8/6QN/PPP1pPPP/2B1K2R w Kkq - 0 1"),
			WHITE_KING,
			kingCaptureScore,
			IE2,
		},
		{
			FromFENString("r1bqk1n1/ppp2ppp/2n5/2b1p3/3P3r/8/PPP1PPPP/RNBQKBNR b KQq - 0 1"),
			BLACK_PAWN,
			pawnCaptureScore,
			ID4,
		},
		{
			FromFENString("r1bqk1n1/ppp2ppp/2n5/2b5/3P3r/8/PPP1PPPP/RNBQKBNR b KQq - 0 1"),
			BLACK_KNIGHT,
			knightCaptureScore,
			ID4,
		},
		{
			FromFENString("r1bqk1n1/ppp2ppp/8/2b5/3P3r/8/PPP1PPPP/RNBQKBNR b KQq - 0 1"),
			BLACK_BISHOP,
			bishopCaptureScore,
			ID4,
		},
		{
			FromFENString("r1bqk1n1/ppp2ppp/8/8/3P3r/8/PPP1PPPP/RNBQKBNR b KQq - 0 1"),
			BLACK_ROOK,
			rookCaptureScore,
			ID4,
		},
		{
			FromFENString("r1bqk1n1/ppp2ppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQq - 0 1"),
			BLACK_QUEEN,
			queenCaptureScore,
			ID4,
		},
		{
			FromFENString("r1b1k3/ppp1Pppp/2q4n/8/8/8/PPP1PPPP/RNBQKBNR b KQq - 0 1"),
			BLACK_KING,
			kingCaptureScore,
			IE7,
		},
	}

	for _, tt := range tests {
		score, move := getSmallestAttackerAndMove(tt.targetSquare, tt.b)
		if score != tt.expectedScore {
			tt.b.Print()
			move.Print()
			t.Errorf("Received unexpected score - expected %d, received %d", tt.expectedScore, score)
		}

		p := tt.b.PieceAt(move.From)
		if p != tt.expectedPiece {
			tt.b.Print()
			move.Print()
			t.Errorf("Received unexpected piece - expected %d, received %d", tt.expectedPiece, p)
		}
	}
}

func TestSee(t *testing.T) {
	var tests = []struct {
		b     Board
		sq    int
		score int
	}{
		{
			FromFENString("1k1r4/1pp4p/p7/4p3/8/P5P1/1PP4P/2K1R3 w - - 0 1"),
			IE5,
			100,
		},
		{
			FromFENString("1k2r3/1pp4p/p7/4p3/8/P5P1/1PP4P/2K1R3 w - - 0 1"),
			IE5,
			-400,
		},
		{
			FromFENString("1k1r3q/1ppn3p/p4b2/4p3/8/P2N2P1/1PP1R1BP/2K1Q3 w HAha - 0 1"),
			IE5,
			-950,
		},
	}

	for _, tt := range tests {
		res := see(tt.sq, &tt.b)
		if res != tt.score {
			tt.b.Print()
			t.Errorf("expected %d, received %d", tt.score, res)
		}
	}
}
