package board

import "testing"

func TestWhitePawnMovesOpeningMoveNoCapture(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IA2, []Move{
			{IA2, IA3, false, false, false, false, 1, false},
			{IA2, IA4, false, false, false, false, 1, true},
		},
		},
		{IB2, []Move{
			{IB2, IB3, false, false, false, false, 1, false},
			{IB2, IB4, false, false, false, false, 1, true},
		},
		},
		{IC2, []Move{
			{IC2, IC3, false, false, false, false, 1, false},
			{IC2, IC4, false, false, false, false, 1, true},
		},
		},
		{ID2, []Move{
			{ID2, ID3, false, false, false, false, 1, false},
			{ID2, ID4, false, false, false, false, 1, true},
		},
		},
		{IE2, []Move{
			{IE2, IE3, false, false, false, false, 1, false},
			{IE2, IE4, false, false, false, false, 1, true},
		},
		},
		{IF2, []Move{
			{IF2, IF3, false, false, false, false, 1, false},
			{IF2, IF4, false, false, false, false, 1, true},
		},
		},
		{IG2, []Move{
			{IG2, IG3, false, false, false, false, 1, false},
			{IG2, IG4, false, false, false, false, 1, true},
		},
		},
		{IH2, []Move{
			{IH2, IH3, false, false, false, false, 1, false},
			{IH2, IH4, false, false, false, false, 1, true},
		},
		},
	}

	for _, tt := range tests {
		moves := b.WhitePawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBlackPawnMovesOpeningMoveNoCapture(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IA7, []Move{
			{IA7, IA6, false, false, false, false, 7, false},
			{IA7, IA5, false, false, false, false, 7, true},
		},
		},
		{IB7, []Move{
			{IB7, IB6, false, false, false, false, 7, false},
			{IB7, IB5, false, false, false, false, 7, true},
		},
		},
		{IC7, []Move{
			{IC7, IC6, false, false, false, false, 7, false},
			{IC7, IC5, false, false, false, false, 7, true},
		},
		},
		{ID7, []Move{
			{ID7, ID6, false, false, false, false, 7, false},
			{ID7, ID5, false, false, false, false, 7, true},
		},
		},
		{IE7, []Move{
			{IE7, IE6, false, false, false, false, 7, false},
			{IE7, IE5, false, false, false, false, 7, true},
		},
		},
		{IF7, []Move{
			{IF7, IF6, false, false, false, false, 7, false},
			{IF7, IF5, false, false, false, false, 7, true},
		},
		},
		{IG7, []Move{
			{IG7, IG6, false, false, false, false, 7, false},
			{IG7, IG5, false, false, false, false, 7, true},
		},
		},
		{IH7, []Move{
			{IH7, IH6, false, false, false, false, 7, false},
			{IH7, IH5, false, false, false, false, 7, true},
		},
		},
	}

	for _, tt := range tests {
		moves := b.BlackPawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestWhitePawnCapturesOriginalSquare(t *testing.T) {
	b := FromFENString("rnbqkbnr/8/8/8/8/pppppppp/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IA2, []Move{
			{IA2, IB3, true, false, false, false, 1, false},
		},
		},
		{IB2, []Move{
			{IB2, IA3, true, false, false, false, 1, false},
			{IB2, IC3, true, false, false, false, 1, false},
		},
		},
		{IC2, []Move{
			{IC2, IB3, true, false, false, false, 1, false},
			{IC2, ID3, true, false, false, false, 1, false},
		},
		},
		{ID2, []Move{
			{ID2, IC3, true, false, false, false, 1, false},
			{ID2, IE3, true, false, false, false, 1, false},
		},
		},
		{IE2, []Move{
			{IE2, ID3, true, false, false, false, 1, false},
			{IE2, IF3, true, false, false, false, 1, false},
		},
		},
		{IF2, []Move{
			{IF2, IE3, true, false, false, false, 1, false},
			{IF2, IG3, true, false, false, false, 1, false},
		},
		},
		{IG2, []Move{
			{IG2, IF3, true, false, false, false, 1, false},
			{IG2, IH3, true, false, false, false, 1, false},
		},
		},
		{IH2, []Move{
			{IH2, IG3, true, false, false, false, 1, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.WhitePawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBlackPawnCapturesOriginalSquare(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/PPPPPPPP/8/8/8/8/RNBQKBNR w KQkq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IA7, []Move{
			{IA7, IB6, true, false, false, false, 7, false},
		},
		},
		{IB7, []Move{
			{IB7, IC6, true, false, false, false, 7, false},
			{IB7, IA6, true, false, false, false, 7, false},
		},
		},
		{IC7, []Move{
			{IC7, ID6, true, false, false, false, 7, false},
			{IC7, IB6, true, false, false, false, 7, false},
		},
		},
		{ID7, []Move{
			{ID7, IE6, true, false, false, false, 7, false},
			{ID7, IC6, true, false, false, false, 7, false},
		},
		},
		{IE7, []Move{
			{IE7, IF6, true, false, false, false, 7, false},
			{IE7, ID6, true, false, false, false, 7, false},
		},
		},
		{IF7, []Move{
			{IF7, IG6, true, false, false, false, 7, false},
			{IF7, IE6, true, false, false, false, 7, false},
		},
		},
		{IG7, []Move{
			{IG7, IH6, true, false, false, false, 7, false},
			{IG7, IF6, true, false, false, false, 7, false},
		},
		},
		{IH7, []Move{
			{IH7, IG6, true, false, false, false, 7, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.BlackPawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestWhiteCaptureAndPushPromotions(t *testing.T) {
	b := FromFENString("1n1q1bn1/P1P4P/5k2/8/8/8/1P1PPPP1/RNBQKBNR w KQ - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IA7, []Move{
			{IA7, IB8, true, false, false, true, WHITE_QUEEN, false},
			{IA7, IB8, true, false, false, true, WHITE_ROOK, false},
			{IA7, IB8, true, false, false, true, WHITE_BISHOP, false},
			{IA7, IB8, true, false, false, true, WHITE_KNIGHT, false},
			{IA7, IA8, false, false, false, true, WHITE_QUEEN, false},
			{IA7, IA8, false, false, false, true, WHITE_ROOK, false},
			{IA7, IA8, false, false, false, true, WHITE_BISHOP, false},
			{IA7, IA8, false, false, false, true, WHITE_KNIGHT, false},
		},
		},
		{IC7, []Move{
			{IC7, IB8, true, false, false, true, WHITE_QUEEN, false},
			{IC7, IB8, true, false, false, true, WHITE_ROOK, false},
			{IC7, IB8, true, false, false, true, WHITE_BISHOP, false},
			{IC7, IB8, true, false, false, true, WHITE_KNIGHT, false},
			{IC7, ID8, true, false, false, true, WHITE_QUEEN, false},
			{IC7, ID8, true, false, false, true, WHITE_ROOK, false},
			{IC7, ID8, true, false, false, true, WHITE_BISHOP, false},
			{IC7, ID8, true, false, false, true, WHITE_KNIGHT, false},
			{IC7, IC8, false, false, false, true, WHITE_QUEEN, false},
			{IC7, IC8, false, false, false, true, WHITE_ROOK, false},
			{IC7, IC8, false, false, false, true, WHITE_BISHOP, false},
			{IC7, IC8, false, false, false, true, WHITE_KNIGHT, false},
		},
		},
		{IH7, []Move{
			{IH7, IG8, true, false, false, true, WHITE_QUEEN, false},
			{IH7, IG8, true, false, false, true, WHITE_ROOK, false},
			{IH7, IG8, true, false, false, true, WHITE_BISHOP, false},
			{IH7, IG8, true, false, false, true, WHITE_KNIGHT, false},
			{IH7, IH8, false, false, false, true, WHITE_QUEEN, false},
			{IH7, IH8, false, false, false, true, WHITE_ROOK, false},
			{IH7, IH8, false, false, false, true, WHITE_BISHOP, false},
			{IH7, IH8, false, false, false, true, WHITE_KNIGHT, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.WhitePawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBlackCaptureAndPushPromotions(t *testing.T) {
	b := FromFENString("rnbqkbnr/1p1pppp1/8/8/8/4K3/p1pPP2p/1N1Q1BN1 w kq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IA2, []Move{
			{IA2, IB1, true, false, false, true, BLACK_QUEEN, false},
			{IA2, IB1, true, false, false, true, BLACK_ROOK, false},
			{IA2, IB1, true, false, false, true, BLACK_BISHOP, false},
			{IA2, IB1, true, false, false, true, BLACK_KNIGHT, false},
			{IA2, IA1, false, false, false, true, BLACK_QUEEN, false},
			{IA2, IA1, false, false, false, true, BLACK_ROOK, false},
			{IA2, IA1, false, false, false, true, BLACK_BISHOP, false},
			{IA2, IA1, false, false, false, true, BLACK_KNIGHT, false},
		},
		},
		{IC2, []Move{
			{IC2, ID1, true, false, false, true, BLACK_QUEEN, false},
			{IC2, ID1, true, false, false, true, BLACK_ROOK, false},
			{IC2, ID1, true, false, false, true, BLACK_BISHOP, false},
			{IC2, ID1, true, false, false, true, BLACK_KNIGHT, false},
			{IC2, IB1, true, false, false, true, BLACK_QUEEN, false},
			{IC2, IB1, true, false, false, true, BLACK_ROOK, false},
			{IC2, IB1, true, false, false, true, BLACK_BISHOP, false},
			{IC2, IB1, true, false, false, true, BLACK_KNIGHT, false},
			{IC2, IC1, false, false, false, true, BLACK_QUEEN, false},
			{IC2, IC1, false, false, false, true, BLACK_ROOK, false},
			{IC2, IC1, false, false, false, true, BLACK_BISHOP, false},
			{IC2, IC1, false, false, false, true, BLACK_KNIGHT, false},
		},
		},
		{IH2, []Move{
			{IH2, IG1, true, false, false, true, BLACK_QUEEN, false},
			{IH2, IG1, true, false, false, true, BLACK_ROOK, false},
			{IH2, IG1, true, false, false, true, BLACK_BISHOP, false},
			{IH2, IG1, true, false, false, true, BLACK_KNIGHT, false},
			{IH2, IH1, false, false, false, true, BLACK_QUEEN, false},
			{IH2, IH1, false, false, false, true, BLACK_ROOK, false},
			{IH2, IH1, false, false, false, true, BLACK_BISHOP, false},
			{IH2, IH1, false, false, false, true, BLACK_KNIGHT, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.BlackPawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestWhiteCaptureEP(t *testing.T) {
	b := FromFENString("rnbqkbnr/p1p1pppp/1p6/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 3")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IE5, []Move{
			{IE5, ID6, true, false, false, false, WHITE_PAWN, false},
			{IE5, IE6, false, false, false, false, WHITE_PAWN, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.WhitePawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBlackCaptureEP(t *testing.T) {
	b := FromFENString("rnbqkbnr/ppp1pppp/8/8/3pP3/1P3N2/P1PP1PPP/RNBQKB1R b KQkq e3 0 3")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{ID4, []Move{
			{ID4, IE3, true, false, false, false, BLACK_PAWN, false},
			{ID4, ID3, false, false, false, false, BLACK_PAWN, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.BlackPawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestWhitePawnCannotCaptureKing(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppPpp/8/8/8/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IF7, []Move{
			{IF7, IG8, true, false, false, true, WHITE_QUEEN, false},
			{IF7, IG8, true, false, false, true, WHITE_ROOK, false},
			{IF7, IG8, true, false, false, true, WHITE_BISHOP, false},
			{IF7, IG8, true, false, false, true, WHITE_KNIGHT, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.WhitePawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBlackPawnCannotCaptureKing(t *testing.T) {
	b := FromFENString("rnbqkbnr/ppppp2p/8/8/8/8/PPPPPpPP/RNBQKBNR w KQkq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IF2, []Move{
			{IF2, IG1, true, false, false, true, BLACK_QUEEN, false},
			{IF2, IG1, true, false, false, true, BLACK_ROOK, false},
			{IF2, IG1, true, false, false, true, BLACK_BISHOP, false},
			{IF2, IG1, true, false, false, true, BLACK_KNIGHT, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.BlackPawnMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestWhiteKnightMovesStartPosition(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IB1, []Move{
			{IB1, IA3, false, false, false, false, WHITE_KNIGHT, false},
			{IB1, IC3, false, false, false, false, WHITE_KNIGHT, false},
		},
		},
		{IG1, []Move{
			{IG1, IF3, false, false, false, false, WHITE_KNIGHT, false},
			{IG1, IH3, false, false, false, false, WHITE_KNIGHT, false},
		},
		},
	}

	for _, tt := range tests {
		moves := b.WhiteKnightMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBlackKnightMovesStartPosition(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IB8, []Move{
			{IB8, IC6, false, false, false, false, BLACK_KNIGHT, false},
			{IB8, IA6, false, false, false, false, BLACK_KNIGHT, false},
		},
		},
		{IG8, []Move{
			{IG8, IH6, false, false, false, false, BLACK_KNIGHT, false},
			{IG8, IF6, false, false, false, false, BLACK_KNIGHT, false},
		},
		},
	}

	for _, tt := range tests {
		moves := b.BlackKnightMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestWhiteKnightEightCaptures(t *testing.T) {
	b := FromFENString("r3k3/pppppppp/2n3n1/4N3/2q3r1/3b1b2/PPPPPPPP/RNBQKB1R w KQq - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IE5, []Move{
			{IE5, ID7, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IF7, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IC6, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IG6, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IG4, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IF3, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, ID3, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IC4, true, false, false, false, WHITE_KNIGHT, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.WhiteKnightMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBlackKnightEightCaptures(t *testing.T) {
	b := FromFENString("r1bqkbnr/pppppppp/2N1B3/1R3N2/3n4/1P3P2/P1PPP1PP/2BQK2R w Kkq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{ID4, []Move{
			{ID4, IC6, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IE6, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IB5, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IF5, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IF3, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IE2, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IC2, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IB3, true, false, false, false, BLACK_KNIGHT, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.BlackKnightMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestWhiteKnightCannotCaptureKing(t *testing.T) {
	b := FromFENString("r7/pppppkpp/2n3n1/4N3/2q3r1/3b1b2/PPPPPPPP/RNBQKB1R w KQq - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IE5, []Move{
			{IE5, ID7, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IC6, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IG6, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IG4, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IF3, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, ID3, true, false, false, false, WHITE_KNIGHT, false},
			{IE5, IC4, true, false, false, false, WHITE_KNIGHT, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.WhiteKnightMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBlackKnightCannotCaptureKing(t *testing.T) {
	b := FromFENString("r1bqkbnr/pppppppp/2N1B3/1R3N2/3n4/1P3P2/P1PPK1PP/2BQ3R w kq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{ID4, []Move{
			{ID4, IC6, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IE6, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IB5, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IF5, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IF3, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IC2, true, false, false, false, BLACK_KNIGHT, false},
			{ID4, IB3, true, false, false, false, BLACK_KNIGHT, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.BlackKnightMoves(tt.sq)
		for i, m := range tt.moves {
			if !equalMoves(moves[i], m) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestWhiteBishopInitialSquare(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	c1Moves := b.WhiteBishopMoves(IC1)
	f1Moves := b.WhiteBishopMoves(IF1)
	if len(c1Moves) != 0 {
		t.Errorf("C1: expected 0 moves, received: %v", c1Moves)
	}
	if len(f1Moves) != 0 {
		t.Errorf("F1: expected 0 moves, received: %v", f1Moves)
	}
}

func TestBlackBishopInitialSquare(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	c8Moves := b.BlackBishopMoves(IC8)
	f8Moves := b.BlackBishopMoves(IF8)
	if len(c8Moves) != 0 {
		t.Errorf("C8: expected 0 moves, received: %v", c8Moves)
	}
	if len(f8Moves) != 0 {
		t.Errorf("F8: expected 0 moves, received: %v", f8Moves)
	}
}

func TestWhiteBishopInitialSquareCenterPawnsMoved(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/3PP3/8/PPP2PPP/RNBQKBNR w KQkq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IC1, []Move{
			{IC1, ID2, false, false, false, false, WHITE_BISHOP, false},
			{IC1, IE3, false, false, false, false, WHITE_BISHOP, false},
			{IC1, IF4, false, false, false, false, WHITE_BISHOP, false},
			{IC1, IG5, false, false, false, false, WHITE_BISHOP, false},
			{IC1, IH6, false, false, false, false, WHITE_BISHOP, false},
		},
		},
		{IF1, []Move{
			{IF1, IE2, false, false, false, false, WHITE_BISHOP, false},
			{IF1, ID3, false, false, false, false, WHITE_BISHOP, false},
			{IF1, IC4, false, false, false, false, WHITE_BISHOP, false},
			{IF1, IB5, false, false, false, false, WHITE_BISHOP, false},
			{IF1, IA6, false, false, false, false, WHITE_BISHOP, false},
		},
		},
	}

	for _, tt := range tests {
		moves := b.WhiteBishopMoves(tt.sq)
		for i, _ := range moves {
			if !equalMoves(moves[i], tt.moves[i]) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBlackBishopInitialSquareCenterPawnsMoved(t *testing.T) {
	b := FromFENString("rnbqkbnr/ppp2ppp/8/3pp3/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IC8, []Move{
			{IC8, ID7, false, false, false, false, BLACK_BISHOP, false},
			{IC8, IE6, false, false, false, false, BLACK_BISHOP, false},
			{IC8, IF5, false, false, false, false, BLACK_BISHOP, false},
			{IC8, IG4, false, false, false, false, BLACK_BISHOP, false},
			{IC8, IH3, false, false, false, false, BLACK_BISHOP, false},
		},
		},
		{IF8, []Move{
			{IF8, IE7, false, false, false, false, BLACK_BISHOP, false},
			{IF8, ID6, false, false, false, false, BLACK_BISHOP, false},
			{IF8, IC5, false, false, false, false, BLACK_BISHOP, false},
			{IF8, IB4, false, false, false, false, BLACK_BISHOP, false},
			{IF8, IA3, false, false, false, false, BLACK_BISHOP, false},
		},
		},
	}

	for _, tt := range tests {
		moves := b.BlackBishopMoves(tt.sq)
		for i, _ := range moves {
			if !equalMoves(moves[i], tt.moves[i]) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestWhiteBishopFourCaptures(t *testing.T) {
	b := FromFENString("r3k3/pp5p/8/2nbpn2/3BB3/2qpbr2/PPPPPPPP/RN1QK1NR w KQq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IE4, []Move{
			{IE4, ID5, true, false, false, false, WHITE_BISHOP, false},
			{IE4, IF5, true, false, false, false, WHITE_BISHOP, false},
			{IE4, ID3, true, false, false, false, WHITE_BISHOP, false},
			{IE4, IF3, true, false, false, false, WHITE_BISHOP, false},
		},
		},
		{ID4, []Move{
			{ID4, IC5, true, false, false, false, WHITE_BISHOP, false},
			{ID4, IE5, true, false, false, false, WHITE_BISHOP, false},
			{ID4, IC3, true, false, false, false, WHITE_BISHOP, false},
			{ID4, IE3, true, false, false, false, WHITE_BISHOP, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.WhiteBishopMoves(tt.sq)
		for i, _ := range moves {
			if !equalMoves(moves[i], tt.moves[i]) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBlackBishopFourCaptures(t *testing.T) {
	b := FromFENString("rn1qk1nr/pppppppp/8/2QNBR2/3bb3/2NPPB2/PPP2PPP/R3K3 w Qkq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IE4, []Move{
			{IE4, ID5, true, false, false, false, BLACK_BISHOP, false},
			{IE4, IF5, true, false, false, false, BLACK_BISHOP, false},
			{IE4, ID3, true, false, false, false, BLACK_BISHOP, false},
			{IE4, IF3, true, false, false, false, BLACK_BISHOP, false},
		},
		},
		{ID4, []Move{
			{ID4, IC5, true, false, false, false, BLACK_BISHOP, false},
			{ID4, IE5, true, false, false, false, BLACK_BISHOP, false},
			{ID4, IC3, true, false, false, false, BLACK_BISHOP, false},
			{ID4, IE3, true, false, false, false, BLACK_BISHOP, false},
		},
		},
	}
	for _, tt := range tests {
		moves := b.BlackBishopMoves(tt.sq)
		for i, _ := range moves {
			if !equalMoves(moves[i], tt.moves[i]) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestBishopsCannotCaptureKing(t *testing.T) {
	b := FromFENString("rnbqk1nr/1pp2ppp/p3p3/1B1p4/1b1P4/P3P3/1PP2PPP/RNBQK1NR w KQkq - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
	}{
		{IB5, []Move{
			{IB5, IA6, true, false, false, false, WHITE_BISHOP, false},
			{IB5, IC6, false, false, false, false, WHITE_BISHOP, false},
			{IB5, ID7, false, false, false, false, WHITE_BISHOP, false},
			{IB5, IA4, false, false, false, false, WHITE_BISHOP, false},
			{IB5, IC4, false, false, false, false, WHITE_BISHOP, false},
			{IB5, ID3, false, false, false, false, WHITE_BISHOP, false},
			{IB5, IE2, false, false, false, false, WHITE_BISHOP, false},
			{IB5, IF1, false, false, false, false, WHITE_BISHOP, false},
		},
		},
		{IB4, []Move{
			{IB4, IA5, false, false, false, false, BLACK_BISHOP, false},
			{IB4, IC5, false, false, false, false, BLACK_BISHOP, false},
			{IB4, ID6, false, false, false, false, BLACK_BISHOP, false},
			{IB4, IE7, false, false, false, false, BLACK_BISHOP, false},
			{IB4, IF8, false, false, false, false, BLACK_BISHOP, false},
			{IB4, IA3, true, false, false, false, BLACK_BISHOP, false},
			{IB4, IC3, false, false, false, false, BLACK_BISHOP, false},
			{IB4, ID2, false, false, false, false, BLACK_BISHOP, false},
		},
		},
	}

	whiteMoves := b.WhiteBishopMoves(tests[0].sq)
	blackMoves := b.BlackBishopMoves(tests[1].sq)
	for i, _ := range whiteMoves {
		if !equalMoves(whiteMoves[i], tests[0].moves[i]) {
			t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[IB5], whiteMoves[i], tests[0].moves[i])
		}
	}
	for i, _ := range blackMoves {
		if !equalMoves(blackMoves[i], tests[1].moves[i]) {
			t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[IB4], blackMoves[i], tests[1].moves[i])
		}
	}
}

func TestRookMovesInitialSquares(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	sqs := []int{IA1, IH1, IA8, IH8}
	for _, sq := range sqs[:2] {
		if len(b.WhiteRookMoves(sq)) != 0 {
			t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[sq], b.WhiteRookMoves(sq), []Move{})
		}
	}
	for _, sq := range sqs[2:] {
		if len(b.BlackRookMoves(sq)) != 0 {
			t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[sq], b.BlackRookMoves(sq), []Move{})
		}
	}
}

func TestQueenMovesInitialSquares(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if len(b.WhiteQueenMoves(ID1)) != 0 {
		t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[ID1], b.WhiteQueenMoves(ID1), []Move{})
	}
	if len(b.BlackQueenMoves(ID8)) != 0 {
		t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[ID8], b.BlackQueenMoves(ID8), []Move{})
	}
}

func TestRookMovesCentralSquares(t *testing.T) {
	b := FromFENString("1nbqkbnr/pppppppp/8/3r4/4R3/8/PPPPPPPP/RNBQKBN1 w Qk - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
		fn    func(int) []Move
	}{
		{
			IE4,
			[]Move{
				{IE4, IE5, false, false, false, false, WHITE_ROOK, false},
				{IE4, IE6, false, false, false, false, WHITE_ROOK, false},
				{IE4, IE7, true, false, false, false, WHITE_ROOK, false},
				{IE4, ID4, false, false, false, false, WHITE_ROOK, false},
				{IE4, IC4, false, false, false, false, WHITE_ROOK, false},
				{IE4, IB4, false, false, false, false, WHITE_ROOK, false},
				{IE4, IA4, false, false, false, false, WHITE_ROOK, false},
				{IE4, IF4, false, false, false, false, WHITE_ROOK, false},
				{IE4, IG4, false, false, false, false, WHITE_ROOK, false},
				{IE4, IH4, false, false, false, false, WHITE_ROOK, false},
				{IE4, IE3, false, false, false, false, WHITE_ROOK, false},
			},
			b.WhiteRookMoves,
		},
		{
			ID5,
			[]Move{
				{ID5, ID6, false, false, false, false, BLACK_ROOK, false},
				{ID5, IC5, false, false, false, false, BLACK_ROOK, false},
				{ID5, IB5, false, false, false, false, BLACK_ROOK, false},
				{ID5, IA5, false, false, false, false, BLACK_ROOK, false},
				{ID5, IE5, false, false, false, false, BLACK_ROOK, false},
				{ID5, IF5, false, false, false, false, BLACK_ROOK, false},
				{ID5, IG5, false, false, false, false, BLACK_ROOK, false},
				{ID5, IH5, false, false, false, false, BLACK_ROOK, false},
				{ID5, ID4, false, false, false, false, BLACK_ROOK, false},
				{ID5, ID3, false, false, false, false, BLACK_ROOK, false},
				{ID5, ID2, true, false, false, false, BLACK_ROOK, false},
			},
			b.BlackRookMoves,
		},
	}

	for _, tt := range tests {
		moves := tt.fn(tt.sq)
		for i, _ := range moves {
			if !equalMoves(moves[i], tt.moves[i]) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestRooksCannotCaptureKings(t *testing.T) {
	b := FromFENString("1nbqkbnr/pppp1ppp/8/4R3/4r3/8/PPPP1PPP/RNBQKBN1 w Qk - 0 1")
	var tests = []struct {
		sq    int
		moves []Move
		fn    func(int) []Move
	}{
		{
			IE5,
			[]Move{
				{IE5, IE6, false, false, false, false, WHITE_ROOK, false},
				{IE5, IE7, false, false, false, false, WHITE_ROOK, false},
				{IE5, ID5, false, false, false, false, WHITE_ROOK, false},
				{IE5, IC5, false, false, false, false, WHITE_ROOK, false},
				{IE5, IB5, false, false, false, false, WHITE_ROOK, false},
				{IE5, IA5, false, false, false, false, WHITE_ROOK, false},
				{IE5, IF5, false, false, false, false, WHITE_ROOK, false},
				{IE5, IG5, false, false, false, false, WHITE_ROOK, false},
				{IE5, IH5, false, false, false, false, WHITE_ROOK, false},
				{IE5, IE4, true, false, false, false, WHITE_ROOK, false},
			},
			b.WhiteRookMoves,
		},
		{
			IE4,
			[]Move{
				{IE4, IE5, true, false, false, false, BLACK_ROOK, false},
				{IE4, ID4, false, false, false, false, BLACK_ROOK, false},
				{IE4, IC4, false, false, false, false, BLACK_ROOK, false},
				{IE4, IB4, false, false, false, false, BLACK_ROOK, false},
				{IE4, IA4, false, false, false, false, BLACK_ROOK, false},
				{IE4, IF4, false, false, false, false, BLACK_ROOK, false},
				{IE4, IG4, false, false, false, false, BLACK_ROOK, false},
				{IE4, IH4, false, false, false, false, BLACK_ROOK, false},
				{IE4, IE3, false, false, false, false, BLACK_ROOK, false},
				{IE4, IE2, false, false, false, false, BLACK_ROOK, false},
			},
			b.BlackRookMoves,
		},
	}

	for _, tt := range tests {
		moves := tt.fn(tt.sq)
		for i, _ := range moves {
			if !equalMoves(moves[i], tt.moves[i]) {
				t.Errorf("SQ: %s, received: %v, expected: %v", SQ_NUM_TO_NAME[tt.sq], moves[i], tt.moves[i])
			}
		}
	}
}

func TestWhiteQueenCentralSquares(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppp1ppp/8/8/4Q3/8/PPPPPPPP/RNB1KBNR w KQkq - 0 1")
	moves := b.WhiteQueenMoves(IE4)
	expectedMoves := []Move{
		{IE4, ID5, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IC6, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IB7, true, false, false, false, WHITE_QUEEN, false},
		{IE4, IF5, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IG6, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IH7, true, false, false, false, WHITE_QUEEN, false},
		{IE4, ID3, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IF3, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IE5, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IE6, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IE7, false, false, false, false, WHITE_QUEEN, false},
		{IE4, ID4, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IC4, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IB4, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IA4, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IF4, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IG4, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IH4, false, false, false, false, WHITE_QUEEN, false},
		{IE4, IE3, false, false, false, false, WHITE_QUEEN, false},
	}
	for i, _ := range moves {
		if !equalMoves(moves[i], expectedMoves[i]) {
			t.Errorf("SQ: %s, received: %v, expected: %v", "E4", moves[i], expectedMoves[i])
		}
	}
}

func TestBlackQueenCentralSquares(t *testing.T) {
	b := FromFENString("rnb1kbnr/pppp1ppp/8/4q3/8/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1")
	moves := b.BlackQueenMoves(IE5)
	expectedMoves := []Move{
		{IE5, ID6, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IF6, false, false, false, false, BLACK_QUEEN, false},
		{IE5, ID4, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IC3, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IB2, true, false, false, false, BLACK_QUEEN, false},
		{IE5, IF4, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IG3, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IH2, true, false, false, false, BLACK_QUEEN, false},
		{IE5, IE6, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IE7, false, false, false, false, BLACK_QUEEN, false},
		{IE5, ID5, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IC5, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IB5, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IA5, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IF5, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IG5, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IH5, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IE4, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IE3, false, false, false, false, BLACK_QUEEN, false},
		{IE5, IE2, false, false, false, false, BLACK_QUEEN, false},
	}
	for i, _ := range moves {
		if !equalMoves(moves[i], expectedMoves[i]) {
			t.Errorf("SQ: %s, received: %v, expected: %v", "E4", moves[i], expectedMoves[i])
		}
	}
}

func TestKingMovesInitialPosition(t *testing.T) {
	b := FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if len(b.WhiteKingMoves(IE1)) != 0 {
		t.Errorf("White king is unable to move in intial pos")
	}
	if len(b.BlackKingMoves(IE8)) != 0 {
		t.Errorf("Black king is unable to move in intial pos")
	}
}

func TestKingsFourCaptures(t *testing.T) {
	b := FromFENString("3q4/pppp4/N1N5/1k2n1n1/N1N2K2/4n1n1/PP2PPPP/3Q4 w - - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
		fn    func(int) []Move
	}{
		{
			IF4,
			[]Move{
				{IF4, IE5, true, false, false, false, WHITE_KING, false},
				{IF4, IE3, true, false, false, false, WHITE_KING, false},
				{IF4, IG3, true, false, false, false, WHITE_KING, false},
				{IF4, IG5, true, false, false, false, WHITE_KING, false},
			},
			b.WhiteKingMoves,
		},
		{
			IB5,
			[]Move{
				{IB5, IA6, true, false, false, false, BLACK_KING, false},
				{IB5, IC6, true, false, false, false, BLACK_KING, false},
				{IB5, IC4, true, false, false, false, BLACK_KING, false},
				{IB5, IA4, true, false, false, false, BLACK_KING, false},
			},
			b.BlackKingMoves,
		},
	}

	for _, tt := range tests {
		moves := tt.fn(tt.sq)
		for i, _ := range moves {
			if !equalMoves(moves[i], tt.moves[i]) {
				t.Errorf("SQ: %s, expected: %v, received: %v", SQ_NUM_TO_NAME[tt.sq], tt.moves[i], moves[i])
			}
		}
	}
}

func TestKingsEightMoves(t *testing.T) {
	b := FromFENString("rnbq1bnr/pppppppp/8/1k6/5K2/8/PPPPPPPP/RNBQ1BNR w - - 0 1")

	var tests = []struct {
		sq    int
		moves []Move
		fn    func(int) []Move
	}{
		{
			IF4,
			[]Move{
				{IF4, IE5, false, false, false, false, WHITE_KING, false},
				{IF4, IF5, false, false, false, false, WHITE_KING, false},
				{IF4, IG5, false, false, false, false, WHITE_KING, false},
				{IF4, IE4, false, false, false, false, WHITE_KING, false},
				{IF4, IG4, false, false, false, false, WHITE_KING, false},
				{IF4, IE3, false, false, false, false, WHITE_KING, false},
				{IF4, IF3, false, false, false, false, WHITE_KING, false},
				{IF4, IG3, false, false, false, false, WHITE_KING, false},
			},
			b.WhiteKingMoves,
		},
		{
			IB5,
			[]Move{
				{IB5, IA6, false, false, false, false, BLACK_KING, false},
				{IB5, IB6, false, false, false, false, BLACK_KING, false},
				{IB5, IC6, false, false, false, false, BLACK_KING, false},
				{IB5, IA5, false, false, false, false, BLACK_KING, false},
				{IB5, IC5, false, false, false, false, BLACK_KING, false},
				{IB5, IA4, false, false, false, false, BLACK_KING, false},
				{IB5, IB4, false, false, false, false, BLACK_KING, false},
				{IB5, IC4, false, false, false, false, BLACK_KING, false},
			},
			b.BlackKingMoves,
		},
	}

	for _, tt := range tests {
		moves := tt.fn(tt.sq)
		for i, _ := range moves {
			if !equalMoves(moves[i], tt.moves[i]) {
				t.Errorf("SQ: %s, expected: %v, received: %v", SQ_NUM_TO_NAME[tt.sq], tt.moves[i], moves[i])
			}
		}
	}
}

func TestSquaresAttackedByWhitePieces(t *testing.T) {
	var tests = []struct {
		attacks     []int
		b           Board
		description string
	}{
		{
			[]int{
				IB3, // pawns
				IA3,
				IC3,
				IB3,
				ID3,
				IC3,
				IE3,
				ID3,
				IF3,
				IE3,
				IG3,
				IF3,
				IH3,
				IG3,
				IA3, // knights
				IC3,
				ID2,
				IF3,
				IH3,
				IE2,
				IB2, // bishops
				ID2,
				IE2,
				IG2,
				IA2, // rooks
				IB1,
				IH2,
				IG1,
				IC2, // queen
				IE2,
				ID2,
				IC1,
				IE1,
				ID2, // king
				IE2,
				IF2,
				ID1,
				IF1,
			},
			FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			"initial position",
		},
		{
			[]int{
				ID5,
				IF5,
				IG3,
				IF7,
				IH7,
				IE6,
				IH3,
				IF3,
				IE4,
				IA6,
				IC6,
				ID7,
				IE8,
				IA4,
				IC4,
				ID3,
				IE2,
				IF1,
				IB4,
				IC5,
				ID6,
				IE7,
				IF8,
				IB2,
				IC1,
				IA2,
				IA3,
				IB1,
				IC1,
				ID1,
				IH2,
				IG1,
				IF1,
				IE1,
				IC2,
				IB3,
				IA4,
				IE2,
				IF3,
				IG4,
				IH5,
				ID2,
				ID3,
				ID4,
				ID5,
				ID6,
				ID7,
				ID8,
				IC1,
				IB1,
				IA1,
				IE1,
				ID2,
				IE2,
				IF2,
				ID1,
				IF1,
			},
			FromFENString("rnbqkbnr/5ppp/8/1B2p1N1/4P3/B7/7P/R2QK2R w KQkq - 0 1"),
			"ls bishop check",
		},
	}

	for _, tt := range tests {
		res := tt.b.SquaresAttackedByWhitePieces()
		if !equalIntSlices(res, tt.attacks) {
			diffIdx := firstDiffIdx(res, tt.attacks)
			t.Errorf("%s: received: %v, expected: %v, diffIdx: %d", tt.description, mapSqIdxToSqNames(res), mapSqIdxToSqNames(tt.attacks), diffIdx)
		}
	}
}

// in the future the individual piece methods should be tested should
// regressions arise
func TestSquaresAttackedByBlackPieces(t *testing.T) {
	var tests = []struct {
		attacks     []int
		b           Board
		description string
	}{
		{
			[]int{
				IB6, // pawns
				IC6,
				IA6,
				ID6,
				IB6,
				IE6,
				IC6,
				IF6,
				ID6,
				IG6,
				IE6,
				IH6,
				IF6,
				IG6,
				ID7, // knights
				IC6,
				IA6,
				IH6,
				IF6,
				IE7,
				IB7, // bishops
				ID7,
				IE7,
				IG7,
				IB8, // rooks
				IA7,
				IG8,
				IH7,
				IC7, // queen
				IE7,
				IC8,
				IE8,
				ID7,
				ID8, // king
				IF8,
				ID7,
				IE7,
				IF7,
			},
			FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			"initial position",
		},
		{
			[]int{
				IB6,
				IH6,
				IF6,
				IG6,
				IE4,
				IC4,
				IE8,
				IG8,
				ID7,
				IH7,
				IH5,
				IG4,
				IE4,
				ID5,
				ID6,
				IF6,
				IC5,
				IG5,
				IG3,
				IF2,
				ID2,
				IC3,
				IB7,
				IA6,
				ID7,
				IE6,
				IF5,
				IG4,
				IH3,
				IE7,
				ID6,
				IC5,
				IB4,
				IA3,
				IG7,
				IB8,
				IC8,
				IA7,
				IG8,
				IF8,
				IH7,
				IB6,
				IC7,
				ID8,
				IB4,
				IC3,
				ID2,
				IE1,
				IA6,
				IA7,
				IB5,
				IC5,
				ID5,
				IA4,
				IA3,
				IA2,
				ID8,
				IF8,
				ID7,
				IE7,
				IF7,
			},
			FromFENString("r1b1kb1r/p5pp/5n2/q2p4/3Pn3/7P/PP2PPP1/RNBQKBNR w KQkq - 0 1"),
			"queen a5 check",
		},
	}

	for _, tt := range tests {
		res := tt.b.SquaresAttackedByBlackPieces()
		if !equalIntSlices(res, tt.attacks) {
			diffIdx := firstDiffIdx(res, tt.attacks)
			t.Errorf("%s: received: %v, expected: %v, diffIdx: %d", tt.description, mapSqIdxToSqNames(res), mapSqIdxToSqNames(tt.attacks), diffIdx)
		}
	}
}

func TestCastling(t *testing.T) {
	var tests = []struct {
		b           Board
		side        int
		moves       []Move
		description string
	}{
		{
			FromFENString("4k3/8/8/8/8/8/8/RN2K2R w KQ - 0 1"),
			WHITE,
			[]Move{
				{IE1, ID2, false, false, false, false, WHITE_KING, false},
				{IE1, IE2, false, false, false, false, WHITE_KING, false},
				{IE1, IF2, false, false, false, false, WHITE_KING, false},
				{IE1, ID1, false, false, false, false, WHITE_KING, false},
				{IE1, IF1, false, false, false, false, WHITE_KING, false},
				{IE1, IG1, false, true, false, false, WHITE_KING, false},
			},
			"check white ck, no attacks, qs blocked",
		},
		{
			FromFENString("r3k2r/8/8/1Q6/8/8/8/4K3 b kq - 0 1"),
			BLACK,
			[]Move{
				{IE8, ID8, false, false, false, false, BLACK_KING, false},
				{IE8, IF8, false, false, false, false, BLACK_KING, false},
				{IE8, IE7, false, false, false, false, BLACK_KING, false},
				{IE8, IF7, false, false, false, false, BLACK_KING, false},
			},
			"black king in check Qb5",
		},
		{
			FromFENString("4k3/8/2q5/8/8/8/8/R3K2R w KQ - 0 1"),
			WHITE,
			[]Move{
				{IE1, ID2, false, false, false, false, WHITE_KING, false},
				{IE1, IE2, false, false, false, false, WHITE_KING, false},
				{IE1, IF2, false, false, false, false, WHITE_KING, false},
				{IE1, ID1, false, false, false, false, WHITE_KING, false},
				{IE1, IF1, false, false, false, false, WHITE_KING, false},
				{IE1, IG1, false, true, false, false, WHITE_KING, false},
			},
			"Qc5 attacks h1 rook white castle qs",
		},
		{
			FromFENString("r3k2r/8/8/8/5Q2/8/8/4K3 b kq - 0 1"),
			BLACK,
			[]Move{
				{IE8, ID8, false, false, false, false, BLACK_KING, false},
				{IE8, ID7, false, false, false, false, BLACK_KING, false},
				{IE8, IE7, false, false, false, false, BLACK_KING, false},
				{IE8, IC8, false, false, true, false, BLACK_KING, false},
			},
			"wQf4 prevents black kc, can cq",
		},
	}

	for _, tt := range tests {
		var moves []Move
		if tt.side == WHITE {
			moves = tt.b.WhiteKingMoves(tt.b.pieceSquares[WHITE_KING][0])
		} else {
			moves = tt.b.BlackKingMoves(tt.b.pieceSquares[BLACK_KING][0])
		}

		if len(moves) != len(tt.moves) {
			t.Errorf("%s: generated %d moves, expected %d", tt.description, len(moves), len(tt.moves))
		}

		for i, _ := range moves {
			if !equalMoves(moves[i], tt.moves[i]) {
				t.Errorf("%s: received: %v, expected: %v", tt.description, moves[i], tt.moves[i])
			}
		}
	}
}

func mapSqIdxToSqNames(s []int) []string {
	r := make([]string, len(s), len(s))
	i := 0

	for i < len(s) {
		r[i] = SQ_NUM_TO_NAME[s[i]]
		i += 1
	}
	return r
}

func firstDiffIdx(s1, s2 []int) int {
	var l int
	if len(s1) < len(s2) {
		l = len(s1)
	} else {
		l = len(s2)
	}
	i := 0
	for i < l {
		if s1[i] != s2[i] {
			return i
		}
		i += 1
	}
	return i
}
