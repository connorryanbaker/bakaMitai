package board

func Mirror(b Board) *Board {
	nb := Board{}
	for i := 0; i < 120; i++ {
		nb.pieces[i] = OFF_BOARD
	}

	for i := 0; i < 64; i++ {
		sq := MAILBOX_64[i]
		flippedSq := FLIPPED_SQUARES[sq]
		nb.pieces[sq] = FLIPPED_PIECES[b.PieceAt(flippedSq)]
	}
	nb.Castle[0] = nb.Castle[2]
	nb.Castle[1] = nb.Castle[3]
	nb.Castle[2] = nb.Castle[0]
	nb.Castle[3] = nb.Castle[1]
	if b.Ep != nil {
		newEp := FLIPPED_SQUARES[*b.Ep]
		nb.Ep = &newEp
	}
	nb.Side = b.Side ^ 1
	nb.updatePieceSquares()
	return &nb
}

var FLIPPED_PIECES = map[int]int{
	WHITE_PAWN:   BLACK_PAWN,
	BLACK_PAWN:   WHITE_PAWN,
	WHITE_KNIGHT: BLACK_KNIGHT,
	BLACK_KNIGHT: WHITE_KNIGHT,
	WHITE_BISHOP: BLACK_BISHOP,
	BLACK_BISHOP: WHITE_BISHOP,
	WHITE_ROOK:   BLACK_ROOK,
	BLACK_ROOK:   WHITE_ROOK,
	WHITE_QUEEN:  BLACK_QUEEN,
	BLACK_QUEEN:  WHITE_QUEEN,
	WHITE_KING:   BLACK_KING,
	BLACK_KING:   WHITE_KING,
}

var FLIPPED_SQUARES = map[int]int{
	IA8: IA1,
	IA7: IA2,
	IA6: IA3,
	IA5: IA4,
	IA4: IA5,
	IA3: IA6,
	IA2: IA7,
	IA1: IA8,
	IB8: IB1,
	IB7: IB2,
	IB6: IB3,
	IB5: IB4,
	IB4: IB5,
	IB3: IB6,
	IB2: IB7,
	IB1: IB8,
	IC8: IC1,
	IC7: IC2,
	IC6: IC3,
	IC5: IC4,
	IC4: IC5,
	IC3: IC6,
	IC2: IC7,
	IC1: IC8,
	ID8: ID1,
	ID7: ID2,
	ID6: ID3,
	ID5: ID4,
	ID4: ID5,
	ID3: ID6,
	ID2: ID7,
	ID1: ID8,
	IE8: IE1,
	IE7: IE2,
	IE6: IE3,
	IE5: IE4,
	IE4: IE5,
	IE3: IE6,
	IE2: IE7,
	IE1: IE8,
	IF8: IF1,
	IF7: IF2,
	IF6: IF3,
	IF5: IF4,
	IF4: IF5,
	IF3: IF6,
	IF2: IF7,
	IF1: IF8,
	IG8: IG1,
	IG7: IG2,
	IG6: IG3,
	IG5: IG4,
	IG4: IG5,
	IG3: IG6,
	IG2: IG7,
	IG1: IG8,
	IH8: IH1,
	IH7: IH2,
	IH6: IH3,
	IH5: IH4,
	IH4: IH5,
	IH3: IH6,
	IH2: IH7,
	IH1: IH8,
}
