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
