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

func TestMakeMoveCastleKingsideBlack(t *testing.T) {
	b := FromFENString("r1bqk2r/pppp1ppp/2n2n2/2b1p3/2B1P3/3P1N2/PPP2PPP/RNBQ1RK1 b kq - 0 1")
	m := Move{
		IE8,
		IG8,
		false,
		true,
		false,
		false,
		BLACK_KING,
		false,
	}
	res := b.MakeMove(m)
	if res != true {
		t.Errorf("makemove returned false for legal castle kingside")
	}

	if b.castle[2] == true || b.castle[3] == true {
		t.Errorf("castle kingside should disable castle permissions, %v", b.castle)
	}

	if b.ep != nil {
		t.Errorf("castle kingside should nullify ep, %d", *b.ep)
	}

	if b.side != WHITE {
		t.Errorf("castle kingside should flip side to play, %d", b.side)
	}

	if b.hply != 1 {
		t.Errorf("castle kingside should increment hply, %d", b.ply)
	}

	if b.ply != 2 {
		t.Errorf("castle kingside should increment ply after black move, %d", b.ply)
	}

	for i := WHITE_PAWN; i <= BLACK_KING; i++ {
		sqs := b.pieceSquares[i]
		var expected []int
		switch i {
		case WHITE_PAWN:
			expected = []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2}
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
			expected = []int{IC6, IF6}
		case BLACK_BISHOP:
			expected = []int{IC8, IC5}
		case BLACK_ROOK:
			expected = []int{IA8, IF8}
		case BLACK_QUEEN:
			expected = []int{ID8}
		case BLACK_KING:
			expected = []int{IG8}
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

func TestMakeMoveCastleQueensideWhite(t *testing.T) {
	b := FromFENString("r2qkb1r/pppbpppp/2n2n2/3p4/3P4/2N1B3/PPPQPPPP/R3KBNR w KQkq - 0 1")
	m := Move{
		IE1,
		IC1,
		false,
		false,
		true,
		false,
		WHITE_KING,
		false,
	}
	res := b.MakeMove(m)
	if res != true {
		t.Errorf("makemove returned false for legal castle queenside")
	}

	if b.castle[0] == true || b.castle[1] == true {
		t.Errorf("castle queenside should disable castle permissions, %v", b.castle)
	}

	if b.ep != nil {
		t.Errorf("castle queenside should nullify ep, %d", *b.ep)
	}

	if b.side != BLACK {
		t.Errorf("castle queenside should flip side to play, %d", b.side)
	}

	if b.hply != 1 {
		t.Errorf("castle queenside should increment hply, %d", b.ply)
	}

	if b.ply != 0 {
		t.Errorf("castle queenside should not increment ply after white move, %d", b.ply)
	}
	for i := WHITE_PAWN; i <= BLACK_KING; i++ {
		sqs := b.pieceSquares[i]
		var expected []int
		switch i {
		case WHITE_PAWN:
			expected = []int{ID4, IA2, IB2, IC2, IE2, IF2, IG2, IH2}
		case WHITE_KNIGHT:
			expected = []int{IC3, IG1}
		case WHITE_BISHOP:
			expected = []int{IE3, IF1}
		case WHITE_ROOK:
			expected = []int{ID1, IH1}
		case WHITE_QUEEN:
			expected = []int{ID2}
		case WHITE_KING:
			expected = []int{IC1}
		case BLACK_PAWN:
			expected = []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7, ID5}
		case BLACK_KNIGHT:
			expected = []int{IC6, IF6}
		case BLACK_BISHOP:
			expected = []int{IF8, ID7}
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

func TestMakeMoveCastleQueensideBlack(t *testing.T) {
	b := FromFENString("r3kb1r/pppbqppp/2n1pn2/3p4/3P3P/2N1BN2/PPPQPPP1/2KR1B1R b kq - 0 1")
	m := Move{
		IE8,
		IC8,
		false,
		false,
		true,
		false,
		BLACK_KING,
		false,
	}
	res := b.MakeMove(m)
	if res != true {
		t.Errorf("makemove returned false for legal castle queenside")
	}

	if b.castle[2] == true || b.castle[3] == true {
		t.Errorf("castle queenside should disable castle permissions, %v", b.castle)
	}

	if b.ep != nil {
		t.Errorf("castle queenside should nullify ep, %d", *b.ep)
	}

	if b.side != WHITE {
		t.Errorf("castle queenside should flip side to play, %d", b.side)
	}

	if b.hply != 1 {
		t.Errorf("castle queenside should increment hply, %d", b.hply)
	}

	if b.ply != 2 {
		t.Errorf("castle queenside should increment ply after black move, %d", b.ply)
	}
	for i := WHITE_PAWN; i <= BLACK_KING; i++ {
		sqs := b.pieceSquares[i]
		var expected []int
		switch i {
		case WHITE_PAWN:
			expected = []int{ID4, IH4, IA2, IB2, IC2, IE2, IF2, IG2}
		case WHITE_KNIGHT:
			expected = []int{IC3, IF3}
		case WHITE_BISHOP:
			expected = []int{IE3, IF1}
		case WHITE_ROOK:
			expected = []int{ID1, IH1}
		case WHITE_QUEEN:
			expected = []int{ID2}
		case WHITE_KING:
			expected = []int{IC1}
		case BLACK_PAWN:
			expected = []int{IA7, IB7, IC7, IF7, IG7, IH7, IE6, ID5}
		case BLACK_KNIGHT:
			expected = []int{IC6, IF6}
		case BLACK_BISHOP:
			expected = []int{IF8, ID7}
		case BLACK_ROOK:
			expected = []int{ID8, IH8}
		case BLACK_QUEEN:
			expected = []int{IE7}
		case BLACK_KING:
			expected = []int{IC8}
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

func TestMakeMoveEPCapture(t *testing.T) {
	var tests = []struct {
		b            Board
		m            Move
		res          bool
		castle       [4]bool
		ep           *int
		side         int
		hply         int
		ply          int
		pieceSquares map[int][]int
		d            string
	}{
		{
			FromFENString("rnbqkbnr/p1p1pppp/1p6/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 1"),
			Move{
				IE5,
				ID6,
				true,
				false,
				false,
				false,
				WHITE_PAWN,
				false,
			},
			true,
			[4]bool{true, true, true, true},
			nil,
			BLACK,
			0,
			0,
			map[int][]int{
				WHITE_PAWN:   []int{ID6, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IC1, IF1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IA7, IC7, IE7, IF7, IG7, IH7, IB6},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			"legal white ep capture d5->e6",
		},
	}

	for _, tt := range tests {
		res := tt.b.MakeMove(tt.m)
		if res != tt.res {
			t.Errorf("%s MakeMove returned unexpected result: %t, expected %t", tt.d, res, tt.res)
		}
		for i, v := range tt.castle {
			if tt.b.castle[i] != v {
				t.Errorf("%s MakeMove produced unexpected castle permission: %v, expected %v", tt.d, tt.b.castle, tt.castle)
			}
		}
		if tt.b.ep != tt.ep {
			t.Errorf("%s MakeMove resulted in unexpected ep: %p, expected %p", tt.d, tt.b.ep, tt.ep)
		}
		if tt.b.side != tt.side {
			t.Errorf("%s MakeMove resulted in unexpected side: %d, expected %d", tt.d, tt.b.side, tt.side)
		}
		if tt.b.hply != tt.hply {
			t.Errorf("%s MakeMove resulted in unexpected hply: %d, expected %d", tt.d, tt.b.hply, tt.hply)
		}
		if tt.b.ply != tt.ply {
			t.Errorf("%s MakeMove resulted in unexpected ply: %d, expected %d", tt.d, tt.b.ply, tt.ply)
		}
		for i := WHITE_PAWN; i <= BLACK_KING; i++ {
			sqs := tt.b.pieceSquares[i]
			if len(sqs) != len(tt.pieceSquares[i]) {
				t.Errorf("p: %d, expected and pieceSquares have different lengths: %v %v %d", i, sqs, tt.pieceSquares[i], tt.m.to)
			}
			for j, _ := range sqs {
				if sqs[j] != tt.pieceSquares[i][j] {
					t.Errorf("p: %d, expected and pieceSquares have different values: %v %v", i, sqs, tt.pieceSquares[i])
				}
			}
		}
	}
}

// todo:
// test legal & illegal ep capture (when capture moves pinned piece)
// test legal & illegal promotion
// test legal & illegal capture
// test legal & illegal quiet move
// test moving rook updates castle permissions
// test moving king updates castle permissions
// test double pawn push
