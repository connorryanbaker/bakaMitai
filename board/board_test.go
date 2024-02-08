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
			FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1"),
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
			tt.b.Print()
			t.Errorf("dsc: %s, side: %d, received %t, expected %t", tt.description, tt.side, received, tt.inCheck)
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

	if b.Castle[0] == true || b.Castle[1] == true {
		t.Errorf("castle kingside should disable castle permissions, %v", b.Castle)
	}

	if b.Ep != nil {
		t.Errorf("castle kingside should nullify ep, %d", *b.Ep)
	}

	if b.Side != BLACK {
		t.Errorf("castle kingside should flip side to play, %d", b.Side)
	}

	if b.Hply != 1 {
		t.Errorf("castle kingside should increment hply, %d", b.Ply)
	}

	if b.Ply != 1 {
		t.Errorf("castle kingside should not increment ply after white move, %d", b.Ply)
	}

	if len(b.History) != 1 {
		t.Errorf("history length should be 1")
	}

	h := b.History[0]
	if h.Move != m {
		t.Errorf("history move does not match; received: %v, expected: %v", h.Move, m)
	}
	if h.previousSquareOccupant != EMPTY_SQUARE {
		t.Errorf("history pso does not match; received: %d, expected: %d", h.previousSquareOccupant, EMPTY_SQUARE)
	}
	if h.ep != nil {
		t.Errorf("history ep does not match; received: %p, expected: nil", h.ep)
	}
	if h.hply != 0 {
		t.Errorf("history hply does not match; received: %d, expected: %d", h.hply, 0)
	}
	if h.ply != 1 {
		t.Errorf("history ply does not match; received: %d, expected: %d", h.ply, 1)
	}
	for i := 0; i < 4; i++ {
		if h.castle[i] != true {
			t.Errorf("history castle does not match; received: %v, expected: %v", h.castle, [4]bool{true, true, true, true})
		}
	}

	for i := WHITE_PAWN; i <= BLACK_KING; i++ {
		sqs := b.PieceSquares[i]
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

	if b.Castle[2] == true || b.Castle[3] == true {
		t.Errorf("castle kingside should disable castle permissions, %v", b.Castle)
	}

	if b.Ep != nil {
		t.Errorf("castle kingside should nullify ep, %d", *b.Ep)
	}

	if b.Side != WHITE {
		t.Errorf("castle kingside should flip side to play, %d", b.Side)
	}

	if b.Hply != 1 {
		t.Errorf("castle kingside should increment hply, %d", b.Ply)
	}

	if b.Ply != 2 {
		t.Errorf("castle kingside should increment ply after black move, %d", b.Ply)
	}

	if len(b.History) != 1 {
		t.Errorf("history length should be 1")
	}

	h := b.History[0]
	if h.Move != m {
		t.Errorf("history move does not match; received: %v, expected: %v", h.Move, m)
	}
	if h.previousSquareOccupant != EMPTY_SQUARE {
		t.Errorf("history pso does not match; received: %d, expected: %d", h.previousSquareOccupant, EMPTY_SQUARE)
	}
	if h.ep != nil {
		t.Errorf("history ep does not match; received: %p, expected: nil", h.ep)
	}
	if h.hply != 0 {
		t.Errorf("history hply does not match; received: %d, expected: %d", h.hply, 0)
	}
	if h.ply != 1 {
		t.Errorf("history ply does not match; received: %d, expected: %d", h.ply, 1)
	}
	expected := [4]bool{false, false, true, true}
	for i := 0; i < 4; i++ {
		if h.castle[i] != expected[i] {
			t.Errorf("history castle does not match; received: %v, expected: %v", h.castle, [4]bool{true, true, true, true})
		}
	}

	for i := WHITE_PAWN; i <= BLACK_KING; i++ {
		sqs := b.PieceSquares[i]
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

	if b.Castle[0] == true || b.Castle[1] == true {
		t.Errorf("castle queenside should disable castle permissions, %v", b.Castle)
	}

	if b.Ep != nil {
		t.Errorf("castle queenside should nullify ep, %d", *b.Ep)
	}

	if b.Side != BLACK {
		t.Errorf("castle queenside should flip side to play, %d", b.Side)
	}

	if b.Hply != 1 {
		t.Errorf("castle queenside should increment hply, %d", b.Ply)
	}

	if b.Ply != 1 {
		t.Errorf("castle queenside should not increment ply after white move, %d", b.Ply)
	}

	if len(b.History) != 1 {
		t.Errorf("history length should be 1")
	}

	h := b.History[0]
	if h.Move != m {
		t.Errorf("history move does not match; received: %v, expected: %v", h.Move, m)
	}
	if h.previousSquareOccupant != EMPTY_SQUARE {
		t.Errorf("history pso does not match; received: %d, expected: %d", h.previousSquareOccupant, EMPTY_SQUARE)
	}
	if h.ep != nil {
		t.Errorf("history ep does not match; received: %p, expected: nil", h.ep)
	}
	if h.hply != 0 {
		t.Errorf("history hply does not match; received: %d, expected: %d", h.hply, 0)
	}
	if h.ply != 1 {
		t.Errorf("history ply does not match; received: %d, expected: %d", h.ply, 1)
	}
	expected := [4]bool{true, true, true, true}
	for i := 0; i < 4; i++ {
		if h.castle[i] != expected[i] {
			t.Errorf("history castle does not match; received: %v, expected: %v", h.castle, [4]bool{true, true, true, true})
		}
	}
	for i := WHITE_PAWN; i <= BLACK_KING; i++ {
		sqs := b.PieceSquares[i]
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

	if b.Castle[2] == true || b.Castle[3] == true {
		t.Errorf("castle queenside should disable castle permissions, %v", b.Castle)
	}

	if b.Ep != nil {
		t.Errorf("castle queenside should nullify ep, %d", *b.Ep)
	}

	if b.Side != WHITE {
		t.Errorf("castle queenside should flip side to play, %d", b.Side)
	}

	if b.Hply != 1 {
		t.Errorf("castle queenside should increment hply, %d", b.Hply)
	}

	if b.Ply != 2 {
		t.Errorf("castle queenside should increment ply after black move, %d", b.Ply)
	}
	if len(b.History) != 1 {
		t.Errorf("history length should be 1")
	}

	h := b.History[0]
	if h.Move != m {
		t.Errorf("history move does not match; received: %v, expected: %v", h.Move, m)
	}
	if h.previousSquareOccupant != EMPTY_SQUARE {
		t.Errorf("history pso does not match; received: %d, expected: %d", h.previousSquareOccupant, EMPTY_SQUARE)
	}
	if h.ep != nil {
		t.Errorf("history ep does not match; received: %p, expected: nil", h.ep)
	}
	if h.hply != 0 {
		t.Errorf("history hply does not match; received: %d, expected: %d", h.hply, 0)
	}
	if h.ply != 1 {
		t.Errorf("history ply does not match; received: %d, expected: %d", h.ply, 1)
	}
	expected := [4]bool{false, false, true, true}
	for i := 0; i < 4; i++ {
		if h.castle[i] != expected[i] {
			t.Errorf("history castle does not match; received: %v, expected: %v", h.castle, [4]bool{true, true, true, true})
		}
	}
	for i := WHITE_PAWN; i <= BLACK_KING; i++ {
		sqs := b.PieceSquares[i]
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
		h            *History
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
			1,
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
			&History{
				EMPTY_SQUARE,
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
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal white ep capture d5->e6",
		},
		{
			FromFENString("rnbqkbnr/ppp1pppp/8/8/3pP3/6P1/PPPP1PBP/RNBQK1NR b KQkq e3 0 1"),
			Move{
				ID4,
				IE3,
				true,
				false,
				false,
				false,
				BLACK_PAWN,
				false,
			},
			true,
			[4]bool{true, true, true, true},
			nil,
			WHITE,
			0,
			2,
			map[int][]int{
				WHITE_PAWN:   []int{IG3, IA2, IB2, IC2, ID2, IF2, IH2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IG2, IC1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7, IE3},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				EMPTY_SQUARE,
				Move{
					ID4,
					IE3,
					true,
					false,
					false,
					false,
					BLACK_PAWN,
					false,
				},
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal black ep capture d4->e3",
		},
		{
			FromFENString("4k3/6b1/8/3pP3/8/2K5/8/8 w - d6 0 1"),
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
			false,
			[4]bool{false, false, false, false},
			&ID6,
			WHITE,
			0,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IE5},
				WHITE_KING:   []int{IC3},
				BLACK_PAWN:   []int{ID5},
				BLACK_BISHOP: []int{IG7},
				BLACK_KING:   []int{IE8},
			},
			nil,
			"illegal white ep capture wp e5 pinned",
		},
		{
			FromFENString("8/5k2/8/8/2pP4/8/B3K3/8 b - d3 0 1"),
			Move{
				IC4,
				ID3,
				true,
				false,
				false,
				false,
				BLACK_PAWN,
				false,
			},
			false,
			[4]bool{false, false, false, false},
			&ID3,
			BLACK,
			0,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{ID4},
				WHITE_BISHOP: []int{IA2},
				WHITE_KING:   []int{IE2},
				BLACK_PAWN:   []int{IC4},
				BLACK_KING:   []int{IF7},
			},
			nil,
			"illegal black ep capture bp c4 pinned",
		},
	}

	for _, tt := range tests {
		res := tt.b.MakeMove(tt.m)
		if res != tt.res {
			t.Errorf("%s MakeMove returned unexpected result: %t, expected %t", tt.d, res, tt.res)
		}
		for i, v := range tt.castle {
			if tt.b.Castle[i] != v {
				t.Errorf("%s MakeMove produced unexpected castle permission: %v, expected %v", tt.d, tt.b.Castle, tt.castle)
			}
		}
		if tt.b.Ep != nil && tt.ep != nil && *tt.b.Ep != *tt.ep {
			t.Errorf("%s MakeMove resulted in unexpected ep: %d, expected %d", tt.d, *tt.b.Ep, *tt.ep)
		}
		if tt.b.Side != tt.side {
			t.Errorf("%s MakeMove resulted in unexpected side: %d, expected %d", tt.d, tt.b.Side, tt.side)
		}
		if tt.b.Hply != tt.hply {
			t.Errorf("%s MakeMove resulted in unexpected hply: %d, expected %d", tt.d, tt.b.Hply, tt.hply)
		}
		if tt.b.Ply != tt.ply {
			t.Errorf("%s MakeMove resulted in unexpected ply: %d, expected %d", tt.d, tt.b.Ply, tt.ply)
		}
		if tt.h == nil && len(tt.b.History) != 0 {
			t.Errorf("%s MakeMove resulted in history when no history was expected: %v", tt.d, tt.b.History)
		} else if tt.h != nil {
			if len(tt.b.History) != 1 {
				t.Errorf("%s len MakeMove resulted in unexpected history: %v", tt.d, tt.b.History)
			}
			h := tt.b.History[0]
			if h.previousSquareOccupant != tt.h.previousSquareOccupant {
				t.Errorf("%s pso MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if !EqualMoves(tt.m, h.Move) {
				t.Errorf("%s move MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.ep != nil && tt.h.ep != nil && *h.ep != *tt.h.ep {
				t.Errorf("%s ep MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.ply != tt.h.ply {
				t.Errorf("%s ply MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.hply != tt.h.hply {
				t.Errorf("%s hply MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			for i := 0; i < 4; i++ {
				if h.castle[i] != tt.h.castle[i] {
					t.Errorf("%s castle MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
				}
			}
		}
		for i := WHITE_PAWN; i <= BLACK_KING; i++ {
			sqs := tt.b.PieceSquares[i]
			if len(sqs) != len(tt.pieceSquares[i]) {
				t.Errorf("p: %d, expected and pieceSquares have different lengths: %v %v %d", i, sqs, tt.pieceSquares[i], tt.m.To)
			}
			for j, _ := range sqs {
				if sqs[j] != tt.pieceSquares[i][j] {
					t.Errorf("p: %d, expected and pieceSquares have different values: %v %v", i, sqs, tt.pieceSquares[i])
				}
			}
		}
	}
}

func TestMakeMovePromotion(t *testing.T) {
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
		h            *History
		d            string
	}{
		{
			FromFENString("8/3P1k2/8/8/2p5/8/B3K3/8 w - - 0 1"),
			Move{
				ID7,
				ID8,
				false,
				false,
				false,
				true,
				WHITE_QUEEN,
				false,
			},
			true,
			[4]bool{false, false, false, false},
			nil,
			BLACK,
			0,
			1,
			map[int][]int{
				WHITE_BISHOP: []int{IA2},
				WHITE_QUEEN:  []int{ID8},
				WHITE_KING:   []int{IE2},
				BLACK_PAWN:   []int{IC4},
				BLACK_KING:   []int{IF7},
			},
			&History{
				EMPTY_SQUARE,
				Move{
					ID7,
					ID8,
					false,
					false,
					false,
					true,
					WHITE_QUEEN,
					false,
				},
				[4]bool{false, false, false, false},
				nil,
				0,
				1,
				uint64(0),
			},
			"wp legal promotion d8Q",
		},
		{
			FromFENString("8/5k2/8/8/8/8/2p1K3/1B6 b - - 0 1"),
			Move{
				IC2,
				IB1,
				true,
				false,
				false,
				true,
				BLACK_QUEEN,
				false,
			},
			true,
			[4]bool{false, false, false, false},
			nil,
			WHITE,
			0,
			2,
			map[int][]int{
				WHITE_KING:  []int{IE2},
				BLACK_QUEEN: []int{IB1},
				BLACK_KING:  []int{IF7},
			},
			&History{
				WHITE_BISHOP,
				Move{
					IC2,
					IB1,
					true,
					false,
					false,
					true,
					BLACK_QUEEN,
					false,
				},
				[4]bool{false, false, false, false},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal promotion capture cxb1Q",
		},
		{
			FromFENString("8/8/8/8/4K3/8/R1p1k3/1B6 b - - 0 1"),
			Move{
				IC2,
				IB1,
				true,
				false,
				false,
				true,
				BLACK_QUEEN,
				false,
			},
			false,
			[4]bool{false, false, false, false},
			nil,
			BLACK,
			0,
			1,
			map[int][]int{
				WHITE_BISHOP: []int{IB1},
				WHITE_ROOK:   []int{IA2},
				WHITE_KING:   []int{IE4},
				BLACK_PAWN:   []int{IC2},
				BLACK_KING:   []int{IE2},
			},
			nil,
			"illegal promotion capture cxb1Q cpawn pinned",
		},
		{
			FromFENString("8/3P4/8/8/r3K3/8/4k3/8 w - - 0 1"),
			Move{
				ID7,
				ID8,
				false,
				false,
				false,
				true,
				WHITE_QUEEN,
				false,
			},
			false,
			[4]bool{false, false, false, false},
			nil,
			WHITE,
			0,
			1,
			map[int][]int{
				WHITE_PAWN: []int{ID7},
				WHITE_KING: []int{IE4},
				BLACK_ROOK: []int{IA4},
				BLACK_KING: []int{IE2},
			},
			nil,
			"illegal promotion white king in check",
		},
	}

	for _, tt := range tests {
		res := tt.b.MakeMove(tt.m)
		if res != tt.res {
			t.Errorf("%s MakeMove returned unexpected result: %t, expected %t", tt.d, res, tt.res)
		}
		for i, v := range tt.castle {
			if tt.b.Castle[i] != v {
				t.Errorf("%s MakeMove produced unexpected castle permission: %v, expected %v", tt.d, tt.b.Castle, tt.castle)
			}
		}
		if tt.b.Ep != nil && tt.ep != nil && *tt.b.Ep != *tt.ep {
			t.Errorf("%s MakeMove resulted in unexpected ep: %d, expected %d", tt.d, *tt.b.Ep, *tt.ep)
		}
		if tt.b.Side != tt.side {
			t.Errorf("%s MakeMove resulted in unexpected side: %d, expected %d", tt.d, tt.b.Side, tt.side)
		}
		if tt.b.Hply != tt.hply {
			t.Errorf("%s MakeMove resulted in unexpected hply: %d, expected %d", tt.d, tt.b.Hply, tt.hply)
		}
		if tt.b.Ply != tt.ply {
			t.Errorf("%s MakeMove resulted in unexpected ply: %d, expected %d", tt.d, tt.b.Ply, tt.ply)
		}
		if tt.h == nil && len(tt.b.History) != 0 {
			t.Errorf("%s MakeMove resulted in history when no history was expected: %v", tt.d, tt.b.History)
		} else if tt.h != nil {
			if len(tt.b.History) != 1 {
				t.Errorf("%s len MakeMove resulted in unexpected history: %v", tt.d, tt.b.History)
			}
			h := tt.b.History[0]
			if h.previousSquareOccupant != tt.h.previousSquareOccupant {
				t.Errorf("%s pso MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if !EqualMoves(tt.m, h.Move) {
				t.Errorf("%s move MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.ep != nil && tt.h.ep != nil && *h.ep != *tt.h.ep {
				t.Errorf("%s ep MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.ply != tt.h.ply {
				t.Errorf("%s ply MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.hply != tt.h.hply {
				t.Errorf("%s hply MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			for i := 0; i < 4; i++ {
				if h.castle[i] != tt.h.castle[i] {
					t.Errorf("%s castle MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
				}
			}
		}
		for i := WHITE_PAWN; i <= BLACK_KING; i++ {
			sqs := tt.b.PieceSquares[i]
			if len(sqs) != len(tt.pieceSquares[i]) {
				t.Errorf("p: %d, expected and pieceSquares have different lengths: %v %v %d", i, sqs, tt.pieceSquares[i], tt.m.To)
			}
			for j, _ := range sqs {
				if sqs[j] != tt.pieceSquares[i][j] {
					t.Errorf("p: %d, expected and pieceSquares have different values: %v %v", i, sqs, tt.pieceSquares[i])
				}
			}
		}
	}
}

func TestMakeMoveCapture(t *testing.T) {
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
		h            *History
		d            string
	}{
		{
			FromFENString("rnbqkb1r/pppp1ppp/8/8/8/4n3/PPPPPpPP/RNBQKBNR w KQkq - 0 1"),
			Move{
				IE1,
				IF2,
				true,
				false,
				false,
				false,
				WHITE_KING,
				false,
			},
			true,
			[4]bool{false, false, true, true},
			nil,
			BLACK,
			0,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IA2, IB2, IC2, ID2, IE2, IG2, IH2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IC1, IF1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IF2},
				BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7},
				BLACK_KNIGHT: []int{IB8, IE3},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				BLACK_PAWN,
				Move{
					IE1,
					IF2,
					true,
					false,
					false,
					false,
					WHITE_KING,
					false,
				},
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"wk legal capture pf7",
		},
		{
			FromFENString("rnbqkb1r/pppp1ppp/8/8/8/4n3/PPPPPpPP/RNBQKBNR w KQkq - 0 1"),
			Move{
				ID2,
				IE3,
				true,
				false,
				false,
				false,
				WHITE_PAWN,
				false,
			},
			false,
			[4]bool{true, true, true, true},
			nil,
			WHITE,
			0,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IA2, IB2, IC2, ID2, IE2, IG2, IH2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IC1, IF1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7, IF2},
				BLACK_KNIGHT: []int{IB8, IE3},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			nil,
			"illegal capture dxe3 wk in check",
		},
		{
			FromFENString("rnbqkb1r/pppp1ppp/8/8/8/8/PPPPPPPn/RNBQKBNR w KQkq - 0 1"),
			Move{
				IH1,
				IH2,
				true,
				false,
				false,
				false,
				WHITE_ROOK,
				false,
			},
			true,
			[4]bool{false, true, true, true},
			nil,
			BLACK,
			0,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IA2, IB2, IC2, ID2, IE2, IF2, IG2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IC1, IF1},
				WHITE_ROOK:   []int{IH2, IA1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7},
				BLACK_KNIGHT: []int{IB8},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				BLACK_KNIGHT,
				Move{
					IH1,
					IH2,
					true,
					false,
					false,
					false,
					WHITE_ROOK,
					false,
				},
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal rook capture, should rm ck perm",
		},
		{
			FromFENString("rnbqkbnr/Rppp1ppp/8/8/8/8/1PPPPPPP/1NBQKBNR b Kkq - 0 1"),
			Move{
				IA8,
				IA7,
				true,
				false,
				false,
				false,
				BLACK_ROOK,
				false,
			},
			true,
			[4]bool{true, false, true, false},
			nil,
			WHITE,
			0,
			2,
			map[int][]int{
				WHITE_PAWN:   []int{IB2, IC2, ID2, IE2, IF2, IG2, IH2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IC1, IF1},
				WHITE_ROOK:   []int{IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IB7, IC7, ID7, IF7, IG7, IH7},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IH8, IA7},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				WHITE_ROOK,
				Move{
					IA8,
					IA7,
					true,
					false,
					false,
					false,
					BLACK_ROOK,
					false,
				},
				[4]bool{true, false, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal black rook capture a7, rm bcq perm",
		},
		{
			FromFENString("rnbqkbnr/pppp1ppp/4P3/1B6/8/8/PPPP1PPP/RNBQK1NR b KQkq - 0 1"),
			Move{
				ID7,
				IE6,
				true,
				false,
				false,
				false,
				WHITE_PAWN,
				false,
			},
			false,
			[4]bool{true, true, true, true},
			nil,
			BLACK,
			0,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IE6, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IB5, IC1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			nil,
			"illegal dxe6, pawn pinned",
		},
	}

	for _, tt := range tests {
		res := tt.b.MakeMove(tt.m)
		if res != tt.res {
			t.Errorf("%s MakeMove returned unexpected result: %t, expected %t", tt.d, res, tt.res)
		}
		for i, v := range tt.castle {
			if tt.b.Castle[i] != v {
				t.Errorf("%s MakeMove produced unexpected castle permission: %v, expected %v", tt.d, tt.b.Castle, tt.castle)
			}
		}
		if tt.b.Ep != nil && tt.ep != nil && *tt.b.Ep != *tt.ep {
			t.Errorf("%s MakeMove resulted in unexpected ep: %d, expected %d", tt.d, *tt.b.Ep, *tt.ep)
		}
		if tt.b.Side != tt.side {
			t.Errorf("%s MakeMove resulted in unexpected side: %d, expected %d", tt.d, tt.b.Side, tt.side)
		}
		if tt.b.Hply != tt.hply {
			t.Errorf("%s MakeMove resulted in unexpected hply: %d, expected %d", tt.d, tt.b.Hply, tt.hply)
		}
		if tt.b.Ply != tt.ply {
			t.Errorf("%s MakeMove resulted in unexpected ply: %d, expected %d", tt.d, tt.b.Ply, tt.ply)
		}
		if tt.h == nil && len(tt.b.History) != 0 {
			t.Errorf("%s MakeMove resulted in history when no history was expected: %v", tt.d, tt.b.History)
		} else if tt.h != nil {
			if len(tt.b.History) != 1 {
				t.Errorf("%s len MakeMove resulted in unexpected history: %v", tt.d, tt.b.History)
			}
			h := tt.b.History[0]
			if h.previousSquareOccupant != tt.h.previousSquareOccupant {
				t.Errorf("%s pso MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if !EqualMoves(tt.m, h.Move) {
				t.Errorf("%s move MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.ep != nil && tt.h.ep != nil && *h.ep != *tt.h.ep {
				t.Errorf("%s ep MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.ply != tt.h.ply {
				t.Errorf("%s ply MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.hply != tt.h.hply {
				t.Errorf("%s hply MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			for i := 0; i < 4; i++ {
				if h.castle[i] != tt.h.castle[i] {
					t.Errorf("%s castle MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
				}
			}
		}
		for i := WHITE_PAWN; i <= BLACK_KING; i++ {
			sqs := tt.b.PieceSquares[i]
			if len(sqs) != len(tt.pieceSquares[i]) {
				t.Errorf("p: %d, expected and pieceSquares have different lengths: %v %v %d", i, sqs, tt.pieceSquares[i], tt.m.To)
			}
			for j, _ := range sqs {
				if sqs[j] != tt.pieceSquares[i][j] {
					t.Errorf("p: %d, expected and pieceSquares have different values: %v %v", i, sqs, tt.pieceSquares[i])
				}
			}
		}
	}
}

func TestMakeMoveQuietMoves(t *testing.T) {
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
		h            *History
		d            string
	}{
		{
			FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			Move{
				IE2,
				IE4,
				false,
				false,
				false,
				false,
				WHITE_PAWN,
				true,
			},
			true,
			[4]bool{true, true, true, true},
			&IE3,
			BLACK,
			0,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IE4, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IC1, IF1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IE7, IF7, IG7, IH7},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				EMPTY_SQUARE,
				Move{
					IE2,
					IE4,
					false,
					false,
					false,
					false,
					WHITE_PAWN,
					true,
				},
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal opening move updates ep square",
		},
		{
			FromFENString("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"),
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
			true,
			[4]bool{true, true, true, true},
			&IE6,
			WHITE,
			0,
			2,
			map[int][]int{
				WHITE_PAWN:   []int{IE4, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IC1, IF1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7, IE5},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				EMPTY_SQUARE,
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
				[4]bool{true, true, true, true},
				&IE3,
				0,
				1,
				uint64(0),
			},
			"legal black opening move updates ep square e6",
		},
		{
			FromFENString("rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1"),
			Move{
				IE1,
				IE2,
				false,
				false,
				false,
				false,
				WHITE_KING,
				false,
			},
			true,
			[4]bool{false, false, true, true},
			nil,
			BLACK,
			1,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IE4, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IC1, IF1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE2},
				BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7, IE5},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				EMPTY_SQUARE,
				Move{
					IE1,
					IE2,
					false,
					false,
					false,
					false,
					WHITE_KING,
					false,
				},
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal bongcloud updates castle permissions",
		},
		{
			FromFENString("rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPPKPPP/RNBQ1BNR b kq - 0 1"),
			Move{
				IE8,
				IE7,
				false,
				false,
				false,
				false,
				BLACK_KING,
				false,
			},
			true,
			[4]bool{false, false, false, false},
			nil,
			WHITE,
			1,
			2,
			map[int][]int{
				WHITE_PAWN:   []int{IE4, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IC1, IF1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE2},
				BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7, IE5},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IF8},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE7},
			},
			&History{
				EMPTY_SQUARE,
				Move{
					IE8,
					IE7,
					false,
					false,
					false,
					false,
					BLACK_KING,
					false,
				},
				[4]bool{false, false, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal bongcloud updates castle permissions black ke7",
		},
		{
			FromFENString("rnbqk1nr/1ppp1pp1/8/pB2p2p/Pb2P2P/8/1PPP1PP1/RNBQK1NR w KQkq - 0 1"),
			Move{
				IH1,
				IH2,
				false,
				false,
				false,
				false,
				WHITE_ROOK,
				false,
			},
			true,
			[4]bool{false, true, true, true},
			nil,
			BLACK,
			1,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IA4, IE4, IH4, IB2, IC2, ID2, IF2, IG2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IB5, IC1},
				WHITE_ROOK:   []int{IH2, IA1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IB7, IC7, ID7, IF7, IG7, IA5, IE5, IH5},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IB4},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				EMPTY_SQUARE,
				Move{
					IH1,
					IH2,
					false,
					false,
					false,
					false,
					WHITE_ROOK,
					false,
				},
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal kings rook move updates castle permissions",
		},
		{
			FromFENString("rnbqk1nr/1ppp1pp1/8/pB2p2p/Pb2P2P/8/1PPP1PP1/RNBQK1NR w KQkq - 0 1"),
			Move{
				IA1,
				IA2,
				false,
				false,
				false,
				false,
				WHITE_ROOK,
				false,
			},
			true,
			[4]bool{true, false, true, true},
			nil,
			BLACK,
			1,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IA4, IE4, IH4, IB2, IC2, ID2, IF2, IG2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IB5, IC1},
				WHITE_ROOK:   []int{IA2, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IB7, IC7, ID7, IF7, IG7, IA5, IE5, IH5},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IB4},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				EMPTY_SQUARE,
				Move{
					IA1,
					IA2,
					false,
					false,
					false,
					false,
					WHITE_ROOK,
					false,
				},
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal queens rook move updates castle permissions",
		},
		{
			FromFENString("rnbqk1nr/1ppp1pp1/8/pB2p2p/Pb2P2P/8/1PPP1PP1/RNBQK1NR b KQkq - 0 1"),
			Move{
				IH8,
				IH7,
				false,
				false,
				false,
				false,
				BLACK_ROOK,
				false,
			},
			true,
			[4]bool{true, true, false, true},
			nil,
			WHITE,
			1,
			2,
			map[int][]int{
				WHITE_PAWN:   []int{IA4, IE4, IH4, IB2, IC2, ID2, IF2, IG2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IB5, IC1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IB7, IC7, ID7, IF7, IG7, IA5, IE5, IH5},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IB4},
				BLACK_ROOK:   []int{IA8, IH7},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				EMPTY_SQUARE,
				Move{
					IH8,
					IH7,
					false,
					false,
					false,
					false,
					BLACK_ROOK,
					false,
				},
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal kings rook move updates castle permissions",
		},
		{
			FromFENString("rnbqk1nr/1ppp1pp1/8/pB2p2p/Pb2P2P/8/1PPP1PP1/RNBQK1NR b KQkq - 0 1"),
			Move{
				IA8,
				IA7,
				false,
				false,
				false,
				false,
				BLACK_ROOK,
				false,
			},
			true,
			[4]bool{true, true, true, false},
			nil,
			WHITE,
			1,
			2,
			map[int][]int{
				WHITE_PAWN:   []int{IA4, IE4, IH4, IB2, IC2, ID2, IF2, IG2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IB5, IC1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IB7, IC7, ID7, IF7, IG7, IA5, IE5, IH5},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IB4},
				BLACK_ROOK:   []int{IH8, IA7},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			&History{
				EMPTY_SQUARE,
				Move{
					IA8,
					IA7,
					false,
					false,
					false,
					false,
					BLACK_ROOK,
					false,
				},
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				uint64(0),
			},
			"legal queens rook move updates castle permissions",
		},
		{
			FromFENString("rnbqk1nr/1ppp1pp1/8/pB2p2p/Pb2P2P/8/1PPP1PP1/RNBQK1NR w KQkq - 0 1"),
			Move{
				ID2,
				ID3,
				false,
				false,
				false,
				false,
				WHITE_PAWN,
				false,
			},
			false,
			[4]bool{true, true, true, true},
			nil,
			WHITE,
			0,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IA4, IE4, IH4, IB2, IC2, ID2, IF2, IG2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IB5, IC1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IB7, IC7, ID7, IF7, IG7, IA5, IE5, IH5},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IB4},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			nil,
			"illegal quiet move returns false",
		},
		{
			FromFENString("rnbqk1nr/1ppp1pp1/8/pB2p2p/Pb2P2P/8/1PPP1PP1/RNBQK1NR b KQkq - 0 1"),
			Move{
				ID7,
				ID6,
				false,
				false,
				false,
				false,
				BLACK_PAWN,
				false,
			},
			false,
			[4]bool{true, true, true, true},
			nil,
			BLACK,
			0,
			1,
			map[int][]int{
				WHITE_PAWN:   []int{IA4, IE4, IH4, IB2, IC2, ID2, IF2, IG2},
				WHITE_KNIGHT: []int{IB1, IG1},
				WHITE_BISHOP: []int{IB5, IC1},
				WHITE_ROOK:   []int{IA1, IH1},
				WHITE_QUEEN:  []int{ID1},
				WHITE_KING:   []int{IE1},
				BLACK_PAWN:   []int{IB7, IC7, ID7, IF7, IG7, IA5, IE5, IH5},
				BLACK_KNIGHT: []int{IB8, IG8},
				BLACK_BISHOP: []int{IC8, IB4},
				BLACK_ROOK:   []int{IA8, IH8},
				BLACK_QUEEN:  []int{ID8},
				BLACK_KING:   []int{IE8},
			},
			nil,
			"illegal quiet move returns false",
		},
	}
	for _, tt := range tests {
		res := tt.b.MakeMove(tt.m)
		if res != tt.res {
			t.Errorf("%s MakeMove returned unexpected result: %t, expected %t", tt.d, res, tt.res)
		}
		for i, v := range tt.castle {
			if tt.b.Castle[i] != v {
				t.Errorf("%s MakeMove produced unexpected castle permission: %v, expected %v", tt.d, tt.b.Castle, tt.castle)
			}
		}
		if tt.b.Ep != nil && tt.ep != nil && *tt.b.Ep != *tt.ep {
			t.Errorf("%s MakeMove resulted in unexpected ep: %d, expected %d", tt.d, *tt.b.Ep, *tt.ep)
		}
		if tt.b.Side != tt.side {
			t.Errorf("%s MakeMove resulted in unexpected side: %d, expected %d", tt.d, tt.b.Side, tt.side)
		}
		if tt.b.Hply != tt.hply {
			t.Errorf("%s MakeMove resulted in unexpected hply: %d, expected %d", tt.d, tt.b.Hply, tt.hply)
		}
		if tt.b.Ply != tt.ply {
			t.Errorf("%s MakeMove resulted in unexpected ply: %d, expected %d", tt.d, tt.b.Ply, tt.ply)
		}
		if tt.h == nil && len(tt.b.History) != 0 {
			t.Errorf("%s MakeMove resulted in history when no history was expected: %v", tt.d, tt.b.History)
		} else if tt.h != nil {
			if len(tt.b.History) != 1 {
				t.Errorf("%s len MakeMove resulted in unexpected history: %v", tt.d, tt.b.History)
			}
			h := tt.b.History[0]
			if h.previousSquareOccupant != tt.h.previousSquareOccupant {
				t.Errorf("%s pso MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if !EqualMoves(tt.m, h.Move) {
				t.Errorf("%s move MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.ep != nil && tt.h.ep != nil && *h.ep != *tt.h.ep {
				t.Errorf("%s ep MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.ply != tt.h.ply {
				t.Errorf("%s ply MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			if h.hply != tt.h.hply {
				t.Errorf("%s hply MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
			}
			for i := 0; i < 4; i++ {
				if h.castle[i] != tt.h.castle[i] {
					t.Errorf("%s castle MakeMove resulted in unexpected history; received: %v, expected: %v", tt.d, h, tt.h)
				}
			}
		}
		for i := WHITE_PAWN; i <= BLACK_KING; i++ {
			sqs := tt.b.PieceSquares[i]
			if len(sqs) != len(tt.pieceSquares[i]) {
				t.Errorf("p: %d, expected and pieceSquares have different lengths: %v %v %d", i, sqs, tt.pieceSquares[i], tt.m.To)
			}
			for j, _ := range sqs {
				if sqs[j] != tt.pieceSquares[i][j] {
					t.Errorf("p: %d, expected and pieceSquares have different values: %v %v", i, sqs, tt.pieceSquares[i])
				}
			}
		}
	}
}

func TestUnmakeMove(t *testing.T) {
	type state struct {
		castle       [4]bool
		ep           *int
		hply         int
		ply          int
		side         int
		hlen         int
		pieceSquares map[int][]int
	}
	var tests = []struct {
		b           Board
		m           Move
		afterMake   state
		afterUnmake state
		d           string
	}{
		{
			FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			Move{
				IE2,
				IE4,
				false,
				false,
				false,
				false,
				WHITE_PAWN,
				true,
			},
			state{
				[4]bool{true, true, true, true},
				&IE3,
				0,
				1,
				BLACK,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IE7, IF7, IG7, IH7},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				WHITE,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IA2, IB2, IC2, ID2, IE2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IE7, IF7, IG7, IH7},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"opening move e4",
		},
		{
			FromFENString("rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"),
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
			state{
				[4]bool{true, true, true, true},
				&IE6,
				0,
				2,
				WHITE,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7, IE5},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				&IE3,
				0,
				1,
				BLACK,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IE7, IF7, IG7, IH7},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"black opening move e5",
		},
		{
			FromFENString("rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 1"),
			Move{
				ID2,
				ID3,
				false,
				false,
				false,
				false,
				WHITE_PAWN,
				false,
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				BLACK,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7, IE5},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				&IE6,
				0,
				1,
				WHITE,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7, IE5},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"white second move d2d3",
		},
		{
			FromFENString("rnbqkbnr/pppp1ppp/8/4p3/4P3/3P4/PPP2PPP/RNBQKBNR b KQkq - 0 1"),
			Move{
				ID7,
				ID6,
				false,
				false,
				false,
				false,
				BLACK_PAWN,
				false,
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				2,
				WHITE,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IF7, IG7, IH7, ID6, IE5},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				BLACK,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IF7, IG7, IH7, IE5},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"black second move d7d6",
		},
		{
			FromFENString("rnbqkbnr/ppp2ppp/3p4/4p3/4P3/3P4/PPP2PPP/RNBQKBNR w KQkq - 0 1"),
			Move{
				IG1,
				IF3,
				false,
				false,
				false,
				false,
				WHITE_KNIGHT,
				false,
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				1,
				1,
				BLACK,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IF3, IB1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IF7, IG7, IH7, ID6, IE5},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				WHITE,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IF7, IG7, IH7, ID6, IE5},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"white third move nf3",
		},
		{
			FromFENString("rnbqkbnr/ppp2ppp/3p4/4p3/4P3/3P1N2/PPP2PPP/RNBQKB1R b KQkq - 1 1"),
			Move{
				IG8,
				IF6,
				false,
				false,
				false,
				false,
				BLACK_KNIGHT,
				false,
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				2,
				2,
				WHITE,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IF3, IB1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IF7, IG7, IH7, ID6, IE5},
					BLACK_KNIGHT: []int{IB8, IF6},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				1,
				1,
				BLACK,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IF3, IB1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IF7, IG7, IH7, ID6, IE5},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"black third move nf6",
		},
		{
			FromFENString("rnbqk2r/ppp1bppp/3p1n2/4p3/4P3/3P1N2/PPP1BPPP/RNBQK2R w KQkq - 0 1"),
			Move{
				IE1,
				IG1,
				false,
				true,
				false,
				false,
				WHITE_KING,
				false,
			},
			state{
				[4]bool{false, false, true, true},
				nil,
				1,
				1,
				BLACK,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IF3, IB1},
					WHITE_BISHOP: []int{IE2, IC1},
					WHITE_ROOK:   []int{IA1, IF1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IG1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IF7, IG7, IH7, ID6, IE5},
					BLACK_KNIGHT: []int{IB8, IF6},
					BLACK_BISHOP: []int{IC8, IE7},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				WHITE,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IF3, IB1},
					WHITE_BISHOP: []int{IE2, IC1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IF7, IG7, IH7, ID6, IE5},
					BLACK_KNIGHT: []int{IB8, IF6},
					BLACK_BISHOP: []int{IC8, IE7},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"white castle kingside",
		},
		{
			FromFENString("rnbqk2r/ppp1bppp/3p1n2/4p3/4P3/3P1N2/PPP1BPPP/RNBQ1RK1 b kq - 0 1"),
			Move{
				IE8,
				IG8,
				false,
				true,
				false,
				false,
				BLACK_KING,
				false,
			},
			state{
				[4]bool{false, false, false, false},
				nil,
				1,
				2,
				WHITE,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IF3, IB1},
					WHITE_BISHOP: []int{IE2, IC1},
					WHITE_ROOK:   []int{IA1, IF1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IG1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IF7, IG7, IH7, ID6, IE5},
					BLACK_KNIGHT: []int{IB8, IF6},
					BLACK_BISHOP: []int{IC8, IE7},
					BLACK_ROOK:   []int{IA8, IF8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IG8},
				},
			},
			state{
				[4]bool{false, false, true, true},
				nil,
				0,
				1,
				BLACK,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IF3, IB1},
					WHITE_BISHOP: []int{IE2, IC1},
					WHITE_ROOK:   []int{IA1, IF1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IG1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IF7, IG7, IH7, ID6, IE5},
					BLACK_KNIGHT: []int{IB8, IF6},
					BLACK_BISHOP: []int{IC8, IE7},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"black castle kingside",
		},
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
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				BLACK,
				1,
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
			},
			state{
				[4]bool{true, true, true, true},
				&ID6,
				0,
				1,
				WHITE,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IE5, IA2, IB2, IC2, ID2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IC7, IE7, IF7, IG7, IH7, IB6, ID5},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"white ep capture",
		},
		{
			FromFENString("rnbqkbnr/ppp1pppp/8/8/3pP3/3P1N2/PPP2PPP/RNBQKB1R b KQkq e3 0 1"),
			Move{
				ID4,
				IE3,
				true,
				false,
				false,
				false,
				BLACK_PAWN,
				false,
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				2,
				WHITE,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IF3, IB1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7, IE3},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				&IE3,
				0,
				1,
				BLACK,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IE4, ID3, IA2, IB2, IC2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IF3, IB1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7, ID4},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"black ep capture",
		},
		{
			FromFENString("r3kbnr/pppqpppp/2n1b3/3p4/3P4/2N1B3/PPPQPPPP/R3KBNR w KQkq - 0 1"),
			Move{
				IE1,
				IC1,
				false,
				false,
				true,
				false,
				WHITE_KING,
				false,
			},
			state{
				[4]bool{false, false, true, true},
				nil,
				1,
				1,
				BLACK,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{ID4, IA2, IB2, IC2, IE2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IC3, IG1},
					WHITE_BISHOP: []int{IE3, IF1},
					WHITE_ROOK:   []int{ID1, IH1},
					WHITE_QUEEN:  []int{ID2},
					WHITE_KING:   []int{IC1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7, ID5},
					BLACK_KNIGHT: []int{IG8, IC6},
					BLACK_BISHOP: []int{IF8, IE6},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID7},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				WHITE,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{ID4, IA2, IB2, IC2, IE2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IC3, IG1},
					WHITE_BISHOP: []int{IE3, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID2},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7, ID5},
					BLACK_KNIGHT: []int{IG8, IC6},
					BLACK_BISHOP: []int{IF8, IE6},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID7},
					BLACK_KING:   []int{IE8},
				},
			},
			"white castle queenside",
		},
		{
			FromFENString("r3kbnr/pppqpppp/2n1b3/3p4/3P4/2N1B3/PPPQPPPP/2KR1BNR b kq - 0 1"),
			Move{
				IE8,
				IC8,
				false,
				false,
				true,
				false,
				BLACK_KING,
				false,
			},
			state{
				[4]bool{false, false, false, false},
				nil,
				1,
				2,
				WHITE,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{ID4, IA2, IB2, IC2, IE2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IC3, IG1},
					WHITE_BISHOP: []int{IE3, IF1},
					WHITE_ROOK:   []int{ID1, IH1},
					WHITE_QUEEN:  []int{ID2},
					WHITE_KING:   []int{IC1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7, ID5},
					BLACK_KNIGHT: []int{IG8, IC6},
					BLACK_BISHOP: []int{IF8, IE6},
					BLACK_ROOK:   []int{ID8, IH8},
					BLACK_QUEEN:  []int{ID7},
					BLACK_KING:   []int{IC8},
				},
			},
			state{
				[4]bool{false, false, true, true},
				nil,
				0,
				1,
				BLACK,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{ID4, IA2, IB2, IC2, IE2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IC3, IG1},
					WHITE_BISHOP: []int{IE3, IF1},
					WHITE_ROOK:   []int{ID1, IH1},
					WHITE_QUEEN:  []int{ID2},
					WHITE_KING:   []int{IC1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7, ID5},
					BLACK_KNIGHT: []int{IG8, IC6},
					BLACK_BISHOP: []int{IF8, IE6},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID7},
					BLACK_KING:   []int{IE8},
				},
			},
			"black castle queenside",
		},
		{
			FromFENString("rnbqkbnr/ppp1pppp/8/8/8/8/PPPPPpPP/RNBQKBNR w KQkq - 0 1"),
			Move{
				IE1,
				IF2,
				true,
				false,
				false,
				false,
				WHITE_KING,
				false,
			},
			state{
				[4]bool{false, false, true, true},
				nil,
				0,
				1,
				BLACK,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IA2, IB2, IC2, ID2, IE2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IF2},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				WHITE,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IA2, IB2, IC2, ID2, IE2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7, IF2},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"white king xf2",
		},
		{
			FromFENString("rnbqkbnr/pppppppR/8/8/8/8/PPPPPPP1/RNBQKBN1 b Qkq - 0 1"),
			Move{
				IH8,
				IH7,
				true,
				false,
				false,
				false,
				BLACK_ROOK,
				false,
			},
			state{
				[4]bool{false, true, false, true},
				nil,
				0,
				2,
				WHITE,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IA2, IB2, IC2, ID2, IE2, IF2, IG2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IE7, IF7, IG7},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH7},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{false, true, true, true},
				nil,
				0,
				1,
				BLACK,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IA2, IB2, IC2, ID2, IE2, IF2, IG2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IH7, IA1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, ID7, IE7, IF7, IG7},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"black rook xh7",
		},
		{
			FromFENString("rnbqkbnr/ppp1pppp/8/8/8/8/pPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			Move{
				IA1,
				IA2,
				true,
				false,
				false,
				false,
				WHITE_ROOK,
				false,
			},
			state{
				[4]bool{true, false, true, true},
				nil,
				0,
				1,
				BLACK,
				1,
				map[int][]int{
					WHITE_PAWN:   []int{IB2, IC2, ID2, IE2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA2, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			state{
				[4]bool{true, true, true, true},
				nil,
				0,
				1,
				WHITE,
				0,
				map[int][]int{
					WHITE_PAWN:   []int{IB2, IC2, ID2, IE2, IF2, IG2, IH2},
					WHITE_KNIGHT: []int{IB1, IG1},
					WHITE_BISHOP: []int{IC1, IF1},
					WHITE_ROOK:   []int{IA1, IH1},
					WHITE_QUEEN:  []int{ID1},
					WHITE_KING:   []int{IE1},
					BLACK_PAWN:   []int{IA7, IB7, IC7, IE7, IF7, IG7, IH7, IA2},
					BLACK_KNIGHT: []int{IB8, IG8},
					BLACK_BISHOP: []int{IC8, IF8},
					BLACK_ROOK:   []int{IA8, IH8},
					BLACK_QUEEN:  []int{ID8},
					BLACK_KING:   []int{IE8},
				},
			},
			"white rookxa2",
		},
		{
			FromFENString("r3k3/1P6/8/8/8/8/8/R3K3 w Qq - 0 1"),
			Move{
				IB7,
				IA8,
				true,
				false,
				false,
				true,
				WHITE_QUEEN,
				false,
			},
			state{
				[4]bool{false, true, false, true},
				nil,
				0,
				1,
				BLACK,
				1,
				map[int][]int{
					WHITE_ROOK:  []int{IA1},
					WHITE_QUEEN: []int{IA8},
					WHITE_KING:  []int{IE1},
					BLACK_KING:  []int{IE8},
				},
			},
			state{
				[4]bool{false, true, false, true},
				nil,
				0,
				1,
				WHITE,
				0,
				map[int][]int{
					WHITE_PAWN: []int{IB7},
					WHITE_ROOK: []int{IA1},
					WHITE_KING: []int{IE1},
					BLACK_ROOK: []int{IA8},
					BLACK_KING: []int{IE8},
				},
			},
			"bxA8Q",
		},
	}

	for _, tt := range tests {
		res := tt.b.MakeMove(tt.m)
		if !res {
			t.Errorf("%s: MakeMove returned false: %v", tt.d, tt.m)
		}
		for i, v := range tt.b.Castle {
			if v != tt.afterMake.castle[i] {
				t.Errorf("%s: MakeMove resulted in unexpected castle; received: %v, expected: %v", tt.d, tt.b.Castle, tt.afterMake.castle)
			}
		}
		if (tt.b.Ep != nil && tt.afterMake.ep == nil) || (tt.b.Ep == nil && tt.afterMake.ep != nil) {
			t.Errorf("%s: MakeMove resulted in unexpected ep; received: %v, expected: %v", tt.d, tt.b.Ep, tt.afterMake.ep)
		} else if tt.b.Ep != nil && tt.afterMake.ep != nil {
			if *tt.b.Ep != *tt.afterMake.ep {
				t.Errorf("%s: MakeMove resulted in unexpected ep; received: %d, expected: %d", tt.d, *tt.b.Ep, *tt.afterMake.ep)
			}
		}
		if tt.b.Hply != tt.afterMake.hply {
			t.Errorf("%s: MakeMove resulted in unexpected hply; received: %d, expected: %d", tt.d, tt.b.Hply, tt.afterMake.hply)
		}
		if tt.b.Ply != tt.afterMake.ply {
			t.Errorf("%s: MakeMove resulted in unexpected ply; received: %d, expected: %d", tt.d, tt.b.Ply, tt.afterMake.ply)
		}
		if tt.b.Side != tt.afterMake.side {
			t.Errorf("%s: MakeMove resulted in unexpected side; received: %d, expected: %d", tt.d, tt.b.Side, tt.afterMake.side)
		}
		if len(tt.b.History) != tt.afterMake.hlen {
			t.Errorf("%s: MakeMove resulted in unexpected history; received: %v, expected len: %d", tt.d, tt.b.History, tt.afterMake.hlen)
		}
		for i := WHITE_PAWN; i <= BLACK_KING; i++ {
			sqs := tt.b.PieceSquares[i]
			if len(sqs) != len(tt.afterMake.pieceSquares[i]) {
				t.Errorf("%s p: %d, MakeMove expected and pieceSquares have different lengths: %v %v %d", tt.d, i, sqs, tt.afterMake.pieceSquares[i], tt.m.To)
			}
			for j, _ := range sqs {
				if sqs[j] != tt.afterMake.pieceSquares[i][j] {
					t.Errorf("%s p: %d, MakeMove expected and pieceSquares have different values: %v %v", tt.d, i, sqs, tt.afterMake.pieceSquares[i])
				}
			}
		}
		tt.b.UnmakeMove()
		for i, v := range tt.b.Castle {
			if v != tt.afterUnmake.castle[i] {
				t.Errorf("%s: UnmakeMove resulted in unexpected castle; received: %v, expected: %v", tt.d, tt.b.Castle, tt.afterUnmake.castle)
			}
		}
		if (tt.b.Ep != nil && tt.afterUnmake.ep == nil) || (tt.b.Ep == nil && tt.afterUnmake.ep != nil) {
			t.Errorf("%s: UnmakeMove resulted in unexpected ep; received: %v, expected: %v", tt.d, tt.b.Ep, tt.afterUnmake.ep)
		} else if tt.b.Ep != nil && tt.afterUnmake.ep != nil {
			if *tt.b.Ep != *tt.afterUnmake.ep {
				t.Errorf("%s: UnmakeMove resulted in unexpected ep; received: %d, expected: %d", tt.d, *tt.b.Ep, *tt.afterUnmake.ep)
			}
		}
		if tt.b.Hply != tt.afterUnmake.hply {
			t.Errorf("%s: UnmakeMove resulted in unexpected hply; received: %d, expected: %d", tt.d, tt.b.Hply, tt.afterUnmake.hply)
		}
		if tt.b.Ply != tt.afterUnmake.ply {
			t.Errorf("%s: UnmakeMove resulted in unexpected ply; received: %d, expected: %d", tt.d, tt.b.Ply, tt.afterUnmake.ply)
		}
		if tt.b.Side != tt.afterUnmake.side {
			t.Errorf("%s: UnmakeMove resulted in unexpected side; received: %d, expected: %d", tt.d, tt.b.Side, tt.afterUnmake.side)
		}
		if len(tt.b.History) != tt.afterUnmake.hlen {
			t.Errorf("%s: UnmakeMove resulted in unexpected history; received: %v, expected len: %d", tt.d, tt.b.History, tt.afterUnmake.hlen)
		}
		for i := WHITE_PAWN; i <= BLACK_KING; i++ {
			sqs := tt.b.PieceSquares[i]
			if len(sqs) != len(tt.afterUnmake.pieceSquares[i]) {
				t.Errorf("p: %d, UnmakeMove expected and pieceSquares have different lengths: %v %v %d", i, sqs, tt.afterUnmake.pieceSquares[i], tt.m.To)
			}
			for j, _ := range sqs {
				if sqs[j] != tt.afterUnmake.pieceSquares[i][j] {
					t.Errorf("p: %d, UnmakeMove expected and pieceSquares have different values: %v %v", i, sqs, tt.afterUnmake.pieceSquares[i])
				}
			}
		}
	}
}

func TestCheckmate(t *testing.T) {
	var tests = []struct {
		b         Board
		checkmate bool
		d         string
	}{
		{
			FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			false,
			"start position",
		},
		{
			FromFENString("rnb1kbnr/pppp1ppp/8/4p3/6Pq/5P2/PPPPP2P/RNBQKBNR w KQkq - 0 1"),
			true,
			"wk checkmate pawns f3g4 Qh4",
		},
		{
			FromFENString("rnbqkbnr/pppppQpp/8/8/2B5/8/PPPPPPPP/RNB1K1NR b KQkq - 0 1"),
			true,
			"bk checkmate on f7 queen supported",
		},
		{
			FromFENString("rnbqkb1r/pppppQpp/7n/8/2B5/8/PPPPPPPP/RNB1K1NR b KQkq - 0 1"),
			false,
			"bk check on f7 queen supported but attacked by knight",
		},
		{
			FromFENString("rnbqkbnr/ppp1pQpp/3p4/8/2B5/8/PPPPPPPP/RNB1K1NR b KQkq - 0 1"),
			false,
			"bk check on f7 queen supported king has escape square",
		},
		{
			FromFENString("4k3/4P3/4K3/8/8/8/8/8 b - - 0 1"),
			false,
			"white stalemates black w/ king and pawn",
		},
	}

	for _, tt := range tests {
		if tt.b.Checkmate() != tt.checkmate {
			t.Errorf("%s: received unexpected checkmate; expected %t, received: %t", tt.d, tt.checkmate, tt.b.Checkmate())
		}
	}
}

func TestStalemate(t *testing.T) {
	var tests = []struct {
		b         Board
		stalemate bool
		d         string
	}{
		{
			FromFENString("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"),
			false,
			"start position",
		},
		{
			FromFENString("rnb1kbnr/pppp1ppp/8/4p3/6Pq/5P2/PPPPP2P/RNBQKBNR w KQkq - 0 1"),
			false,
			"wk checkmate pawns f3g4 Qh4",
		},
		{
			FromFENString("rnbqkbnr/pppppQpp/8/8/2B5/8/PPPPPPPP/RNB1K1NR b KQkq - 0 1"),
			false,
			"bk checkmate on f7 queen supported",
		},
		{
			FromFENString("rnbqkb1r/pppppQpp/7n/8/2B5/8/PPPPPPPP/RNB1K1NR b KQkq - 0 1"),
			false,
			"bk check on f7 queen supported but attacked by knight",
		},
		{
			FromFENString("rnbqkbnr/ppp1pQpp/3p4/8/2B5/8/PPPPPPPP/RNB1K1NR b KQkq - 0 1"),
			false,
			"bk check on f7 queen supported king has escape square",
		},
		{
			FromFENString("4k3/4P3/4K3/8/8/8/8/8 b - - 0 1"),
			true,
			"white stalemates black w/ king and pawn",
		},
		{
			FromFENString("3r1r2/8/8/8/8/4k3/8/4K3 w - - 0 1"),
			true,
			"black stalemates white w/ king and two rooks",
		},
		{
			FromFENString("3r1r2/4P3/8/8/8/4k3/8/4K3 w - - 0 1"),
			false,
			"w/o pawn, black stalemates white w/ king and two rooks, pawn has 3 moves",
		},
	}

	for _, tt := range tests {
		if tt.b.Stalemate() != tt.stalemate {
			t.Errorf("%s: received unexpected stalemate; expected %t, received: %t", tt.d, tt.stalemate, tt.b.Stalemate())
		}
	}
}

func TestThreefoldRepetition(t *testing.T) {
	b := FromFENString("7q/8/8/8/8/4k3/8/7K w - - 0 1")
	b.MakeMove(Move{
		IH1,
		IG1,
		false,
		false,
		false,
		false,
		WHITE_KING,
		false,
	})
	b.MakeMove(Move{
		IH8,
		IG8,
		false,
		false,
		false,
		false,
		BLACK_QUEEN,
		false,
	})
	b.MakeMove(Move{
		IG1,
		IH1,
		false,
		false,
		false,
		false,
		WHITE_KING,
		false,
	})
	b.MakeMove(Move{
		IG8,
		IH8,
		false,
		false,
		false,
		false,
		BLACK_QUEEN,
		false,
	})
	b.MakeMove(Move{
		IH1,
		IG1,
		false,
		false,
		false,
		false,
		WHITE_KING,
		false,
	})
	b.MakeMove(Move{
		IH8,
		IG8,
		false,
		false,
		false,
		false,
		BLACK_QUEEN,
		false,
	})
	b.MakeMove(Move{
		IG1,
		IH1,
		false,
		false,
		false,
		false,
		WHITE_KING,
		false,
	})
	b.MakeMove(Move{
		IG8,
		IH8,
		false,
		false,
		false,
		false,
		BLACK_QUEEN,
		false,
	})
	if b.ThreefoldRepetition() != true {
		t.Errorf("Unexpected threefold repetition; received: %t, expected: %t", false, true)
	}
}

func TestFiftyMoveDraw(t *testing.T) {
	b := FromFENString("7q/8/8/8/8/4k3/8/7K w - - 99 100")
	b.MakeMove(Move{
		IH1,
		IG1,
		false,
		false,
		false,
		false,
		WHITE_KING,
		false,
	})
	if b.FiftyMoveDraw() != true {
		t.Errorf("hply of %d should indicate 50 move draw; received: false, expected true", b.Hply)
	}
}

func TestInsufficientMaterial(t *testing.T) {
	var tests = []struct {
		b Board
		r bool
	}{
		{
			FromFENString("8/8/8/8/8/4k3/8/7K w - - 30 100"),
			true,
		},
		{
			FromFENString("5k2/8/8/8/8/8/6B1/4K3 w - - 0 1"),
			true,
		},
		{
			FromFENString("5k2/8/8/8/8/8/6N1/4K3 w - - 0 1"),
			true,
		},
		{
			FromFENString("5k2/8/8/8/8/8/6b1/4K3 w - - 0 1"),
			true,
		},
		{
			FromFENString("5k2/8/8/8/8/8/6b1/4K3 w - - 0 1"),
			true,
		},
		{
			FromFENString("5k2/8/8/8/8/8/6r1/4K3 w - - 0 1"),
			false,
		},
		{
			FromFENString("5k2/8/8/8/8/8/6R1/4K3 w - - 0 1"),
			false,
		},
		{
			FromFENString("5k2/8/8/8/8/8/6bn/4K3 w - - 0 1"),
			false,
		},
		{
			FromFENString("5k2/8/8/8/8/8/6BN/4K3 w - - 0 1"),
			false,
		},
	}
	for _, tt := range tests {
		if tt.b.InsufficientMaterial() != tt.r {
			tt.b.Print()
			t.Errorf("Unexpected result for insufficient material call")
		}
	}
}

func TestUnmakeMoveRegression(t *testing.T) {
	// this fen resulted in bq disappearing
	var tests = []struct {
		f string
	}{
		{
			"r2qk2r/pppb1ppp/8/1B2N3/1b2p3/8/PPPBPPPP/R2QK2R w KQkq - 1 1",
		},
		{
			"r2qkbnr/ppp1pppp/2n5/1B6/4p1P1/5b1P/PPPP1P2/RNBQK2R b KQkq - 0 6",
		},
		{
			"r2qk1nr/ppp1bpp1/2n1p2p/5b2/4N3/PP1B1N2/1BPP1PPP/R2Q1RK1 b kq - 0 1",
		},
	}
	for _, tt := range tests {
		b := FromFENString(tt.f)
		for _, m := range b.LegalMoves() {
			bh := b.Hash()
			r := b.MakeMove(m)
			if r != true {
				t.Errorf("legal move %v returned false", m)
			}

			for _, m2 := range b.LegalMoves() {
				bh2 := b.Hash()
				r2 := b.MakeMove(m2)
				if r2 != true {
					t.Errorf("legal move %v returned false", m2)
				}
				if m2.CastleQueenside {
					t.Errorf("illegal move castle queenside :%v", m2)
				}
				b.UnmakeMove()
				if b.Hash() != bh2 {
					t.Errorf("Unmake move after %v resulted in changed hash", m2)
				}
			}
			b.UnmakeMove()
			if b.Hash() != bh {
				t.Errorf("Unmake move after %v resulted in changed hash", m)
			}
		}
	}
}
