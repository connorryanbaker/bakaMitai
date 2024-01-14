package board

import "testing"

func TestPieceAtNewBoard(t *testing.T) {
	var tests = []struct {
		sq       int
		expected int
	}{
		{IA8, BLACK_ROOK},
		{IB8, BLACK_KNIGHT},
		{IC8, BLACK_BISHOP},
		{ID8, BLACK_QUEEN},
		{IE8, BLACK_KING},
		{IF8, BLACK_BISHOP},
		{IG8, BLACK_KNIGHT},
		{IH8, BLACK_ROOK},
		{IA7, BLACK_PAWN},
		{IB7, BLACK_PAWN},
		{IC7, BLACK_PAWN},
		{ID7, BLACK_PAWN},
		{IE7, BLACK_PAWN},
		{IF7, BLACK_PAWN},
		{IG7, BLACK_PAWN},
		{IH7, BLACK_PAWN},
		{IA6, EMPTY_SQUARE},
		{IB6, EMPTY_SQUARE},
		{IC6, EMPTY_SQUARE},
		{ID6, EMPTY_SQUARE},
		{IE6, EMPTY_SQUARE},
		{IF6, EMPTY_SQUARE},
		{IG6, EMPTY_SQUARE},
		{IH6, EMPTY_SQUARE},
		{IA5, EMPTY_SQUARE},
		{IB5, EMPTY_SQUARE},
		{IC5, EMPTY_SQUARE},
		{ID5, EMPTY_SQUARE},
		{IE5, EMPTY_SQUARE},
		{IF5, EMPTY_SQUARE},
		{IG5, EMPTY_SQUARE},
		{IH5, EMPTY_SQUARE},
		{IA4, EMPTY_SQUARE},
		{IB4, EMPTY_SQUARE},
		{IC4, EMPTY_SQUARE},
		{ID4, EMPTY_SQUARE},
		{IE4, EMPTY_SQUARE},
		{IF4, EMPTY_SQUARE},
		{IG4, EMPTY_SQUARE},
		{IH4, EMPTY_SQUARE},
		{IA3, EMPTY_SQUARE},
		{IB3, EMPTY_SQUARE},
		{IC3, EMPTY_SQUARE},
		{ID3, EMPTY_SQUARE},
		{IE3, EMPTY_SQUARE},
		{IF3, EMPTY_SQUARE},
		{IG3, EMPTY_SQUARE},
		{IH3, EMPTY_SQUARE},
		{IA2, WHITE_PAWN},
		{IB2, WHITE_PAWN},
		{IC2, WHITE_PAWN},
		{ID2, WHITE_PAWN},
		{IE2, WHITE_PAWN},
		{IF2, WHITE_PAWN},
		{IG2, WHITE_PAWN},
		{IH2, WHITE_PAWN},
		{IA1, WHITE_ROOK},
		{IB1, WHITE_KNIGHT},
		{IC1, WHITE_BISHOP},
		{ID1, WHITE_QUEEN},
		{IE1, WHITE_KING},
		{IF1, WHITE_BISHOP},
		{IG1, WHITE_KNIGHT},
		{IH1, WHITE_ROOK},
	}

	b := NewBoard()

	for _, tt := range tests {
		received := b.PieceAt(tt.sq)
		if received != tt.expected {
			t.Errorf("received %d, expected %d", received, tt.expected)
		}
	}
}

func TestInCheck(t *testing.T) {
	var tests = []struct {
		b           Board
		side        int
		inCheck     bool
		description string
	}{
		{
			FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			WHITE,
			false,
			"init",
		},
		{
			FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			BLACK,
			false,
			"init",
		},
		{
			FromFENString("rnbqk1nr/ppp2ppp/4p3/3p4/1b1P4/4PN2/PPP2PPP/RNBQKB1R w KQkq - 0 1"),
			WHITE,
			true,
			"bishop b4 check",
		},
		{
			FromFENString("rnbqkbnr/ppp2ppp/4p3/1B1p4/3P4/4P3/PPP2PPP/RNBQK1NR w KQkq - 0 1"),
			BLACK,
			true,
			"bishop b5 check",
		},
		{
			FromFENString("rnbqkb1r/ppp2ppp/4p3/3p4/3P4/4Pn2/PPP2PPP/RNBQKBNR w KQkq - 0 1"),
			WHITE,
			true,
			"knight f3 check",
		},
		{
			FromFENString("rnbqkbnr/ppN2ppp/4p3/3p4/3P4/4P3/PPP2PPP/R1BQKBNR b KQkq - 0 1"),
			BLACK,
			true,
			"knight c7 check",
		},
		{
			FromFENString("1nbqkbnr/pp2rppp/8/3p4/3P4/8/PPP2PPP/RNBQKBNR w KQk - 0 1"),
			WHITE,
			true,
			"rook e7 check",
		},
		{
			FromFENString("rnbqkbnr/pp3ppp/8/3p4/3P4/8/PPP1RPPP/RNBQKBN1 b Qk - 0 1"),
			BLACK,
			true,
			"rook e2 check",
		},
		{
			FromFENString("rnb1kbnr/pp3ppp/8/3p4/3P3q/5P2/PPP3PP/RNBQKBNR b KQkq - 0 1"),
			WHITE,
			true,
			"queen h4 check",
		},
		{
			FromFENString("rnbqkpnr/pp4pp/8/3p3Q/3P4/8/PPP2PPP/RNB1KBNR b KQkq - 0 1"),
			BLACK,
			true,
			"queen h5 check",
		},
	}

	for _, tt := range tests {
		received := tt.b.InCheck(tt.side)
		if received != tt.inCheck {
			t.Errorf("dsc: %s, received %t, expected %t", tt.description, received, tt.inCheck)
		}
	}
}

func TestMakeMoveWrongColor(t *testing.T) {
	var tests = []struct {
		b        Board
		m        Move
		expected bool
	}{
		{
			FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			Move{
				IE7,
				IE5,
				false,
				false,
				false,
				false,
				BLACK_PAWN,
				true,
			},
			false,
		},
		{
			FromFENString("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1"),
			Move{
				ID2,
				ID4,
				false,
				false,
				false,
				false,
				WHITE_PAWN,
				true,
			},
			false,
		},
	}

	for _, tt := range tests {
		received := tt.b.MakeMove(tt.m)
		if received != tt.expected {
			t.Errorf("illegal move allowed, expected: %t, received: %t", tt.expected, received)
		}
	}
}

func TestMakeMoveCastleKingsideWhite(t *testing.T) {
	b := FromFENString("r1bqk1nr/pppp1ppp/2n5/2b1p3/2B1P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 0 1")
	m := Move{
		IE1,
		IG1,
		false,
		true,
		false,
		false,
		WHITE_KING,
		false,
	}
	res := b.MakeMove(m)
	if res != true {
		t.Errorf("makemove returned false for legal castle kingside")
	}

	if b.castle[0] == true || b.castle[1] == true {
		t.Errorf("castle kingside should disable castle permissions, %v", b.castle)
	}

	if b.ep != nil {
		t.Errorf("castle kingside should nullify ep, %d", *b.ep)
	}

	if b.side != BLACK {
		t.Errorf("castle kingside should flip side to play, %d", b.side)
	}

	if b.hply != 1 {
		t.Errorf("castle kingside should increment hply, %d", b.ply)
	}

	if b.ply != 0 {
		t.Errorf("castle kingside should not increment ply after white move, %d", b.ply)
	}

	for i := WHITE_PAWN; i <= BLACK_KING; i++ {
		sqs := b.pieceSquares[i]
		var expected []int
		switch i {
		case WHITE_PAWN:
			expected = []int{IE4, IA2, IB2, IC2, ID2, IF2, IG2, IH2}
		case WHITE_KNIGHT:
			expected = []int{IF3, IB1}
		case WHITE_BISHOP:
			expected = []int{IC4, IC1}
		case WHITE_ROOK:
			expected = []int{IA1, IF1}
		case WHITE_QUEEN:
			expected = []int{ID1}
		case WHITE_KING:
			expected = []int{IG1}
		case BLACK_PAWN:
			expected = []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7, IE5}
		case BLACK_KNIGHT:
			expected = []int{IG8, IC6}
		case BLACK_BISHOP:
			expected = []int{IC8, IC5}
		case BLACK_ROOK:
			expected = []int{IA8, IH8}
		case BLACK_QUEEN:
			expected = []int{ID8}
		case BLACK_KING:
			expected = []int{IE8}
		}
		if len(expected) != len(sqs) {
			t.Errorf("p: %d, expected and pieceSquares have different lengths: %v %v", i, sqs, expected)
		}
		for j, _ := range expected {
			if expected[j] != sqs[j] {
				t.Errorf("p: %d, expected and pieceSquares have different values: %v %v", i, sqs, expected)
			}
		}
	}
}
