package board

import "strings"

const OFF_BOARD = -1
const EMPTY_SQUARE = 0
const WHITE_PAWN = 1
const WHITE_KNIGHT = 2
const WHITE_BISHOP = 3
const WHITE_ROOK = 4
const WHITE_QUEEN = 5
const WHITE_KING = 6
const BLACK_PAWN = 7
const BLACK_KNIGHT = 8
const BLACK_BISHOP = 9
const BLACK_ROOK = 10
const BLACK_QUEEN = 11
const BLACK_KING = 12

const WHITE = 0
const BLACK = 1

var COLORS = map[int]string{
	WHITE: "WHITE",
	BLACK: "BLACK",
}

var INIT_PIECES = [120]int{
	OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD,
	OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD,
	OFF_BOARD, BLACK_ROOK, BLACK_KNIGHT, BLACK_BISHOP, BLACK_QUEEN, BLACK_KING, BLACK_BISHOP, BLACK_KNIGHT, BLACK_ROOK, OFF_BOARD,
	OFF_BOARD, BLACK_PAWN, BLACK_PAWN, BLACK_PAWN, BLACK_PAWN, BLACK_PAWN, BLACK_PAWN, BLACK_PAWN, BLACK_PAWN, OFF_BOARD,
	OFF_BOARD, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, OFF_BOARD,
	OFF_BOARD, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, OFF_BOARD,
	OFF_BOARD, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, OFF_BOARD,
	OFF_BOARD, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, EMPTY_SQUARE, OFF_BOARD,
	OFF_BOARD, WHITE_PAWN, WHITE_PAWN, WHITE_PAWN, WHITE_PAWN, WHITE_PAWN, WHITE_PAWN, WHITE_PAWN, WHITE_PAWN, OFF_BOARD,
	OFF_BOARD, WHITE_ROOK, WHITE_KNIGHT, WHITE_BISHOP, WHITE_QUEEN, WHITE_KING, WHITE_BISHOP, WHITE_KNIGHT, WHITE_ROOK, OFF_BOARD,
	OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD,
	OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD, OFF_BOARD,
}

var MAILBOX_64 = [64]int{
	21, 22, 23, 24, 25, 26, 27, 28,
	31, 32, 33, 34, 35, 36, 37, 38,
	41, 42, 43, 44, 45, 46, 47, 48,
	51, 52, 53, 54, 55, 56, 57, 58,
	61, 62, 63, 64, 65, 66, 67, 68,
	71, 72, 73, 74, 75, 76, 77, 78,
	81, 82, 83, 84, 85, 86, 87, 88,
	91, 92, 93, 94, 95, 96, 97, 98,
}

func file(sq int) int {
	return sq % 10
}

func fileName(sq int) string {
	return string(FILES[file(sq)-1])
}

const FILES = "abcdefgh"

var INIT_CASTLE = [4]bool{true, true, true, true}

var INIT_PIECE_SQUARES = map[int][]int{
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
}

const (
	A8 = iota
	B8 = iota
	C8 = iota
	D8 = iota
	E8 = iota
	F8 = iota
	G8 = iota
	H8 = iota
	A7 = iota
	B7 = iota
	C7 = iota
	D7 = iota
	E7 = iota
	F7 = iota
	G7 = iota
	H7 = iota
	A6 = iota
	B6 = iota
	C6 = iota
	D6 = iota
	E6 = iota
	F6 = iota
	G6 = iota
	H6 = iota
	A5 = iota
	B5 = iota
	C5 = iota
	D5 = iota
	E5 = iota
	F5 = iota
	G5 = iota
	H5 = iota
	A4 = iota
	B4 = iota
	C4 = iota
	D4 = iota
	E4 = iota
	F4 = iota
	G4 = iota
	H4 = iota
	A3 = iota
	B3 = iota
	C3 = iota
	D3 = iota
	E3 = iota
	F3 = iota
	G3 = iota
	H3 = iota
	A2 = iota
	B2 = iota
	C2 = iota
	D2 = iota
	E2 = iota
	F2 = iota
	G2 = iota
	H2 = iota
	A1 = iota
	B1 = iota
	C1 = iota
	D1 = iota
	E1 = iota
	F1 = iota
	G1 = iota
	H1 = iota
)

// internal square representations
// helpful for testing
var IA8 = MAILBOX_64[A8]
var IB8 = MAILBOX_64[B8]
var IC8 = MAILBOX_64[C8]
var ID8 = MAILBOX_64[D8]
var IE8 = MAILBOX_64[E8]
var IF8 = MAILBOX_64[F8]
var IG8 = MAILBOX_64[G8]
var IH8 = MAILBOX_64[H8]
var IA7 = MAILBOX_64[A7]
var IB7 = MAILBOX_64[B7]
var IC7 = MAILBOX_64[C7]
var ID7 = MAILBOX_64[D7]
var IE7 = MAILBOX_64[E7]
var IF7 = MAILBOX_64[F7]
var IG7 = MAILBOX_64[G7]
var IH7 = MAILBOX_64[H7]
var IA6 = MAILBOX_64[A6]
var IB6 = MAILBOX_64[B6]
var IC6 = MAILBOX_64[C6]
var ID6 = MAILBOX_64[D6]
var IE6 = MAILBOX_64[E6]
var IF6 = MAILBOX_64[F6]
var IG6 = MAILBOX_64[G6]
var IH6 = MAILBOX_64[H6]
var IA5 = MAILBOX_64[A5]
var IB5 = MAILBOX_64[B5]
var IC5 = MAILBOX_64[C5]
var ID5 = MAILBOX_64[D5]
var IE5 = MAILBOX_64[E5]
var IF5 = MAILBOX_64[F5]
var IG5 = MAILBOX_64[G5]
var IH5 = MAILBOX_64[H5]
var IA4 = MAILBOX_64[A4]
var IB4 = MAILBOX_64[B4]
var IC4 = MAILBOX_64[C4]
var ID4 = MAILBOX_64[D4]
var IE4 = MAILBOX_64[E4]
var IF4 = MAILBOX_64[F4]
var IG4 = MAILBOX_64[G4]
var IH4 = MAILBOX_64[H4]
var IA3 = MAILBOX_64[A3]
var IB3 = MAILBOX_64[B3]
var IC3 = MAILBOX_64[C3]
var ID3 = MAILBOX_64[D3]
var IE3 = MAILBOX_64[E3]
var IF3 = MAILBOX_64[F3]
var IG3 = MAILBOX_64[G3]
var IH3 = MAILBOX_64[H3]
var IA2 = MAILBOX_64[A2]
var IB2 = MAILBOX_64[B2]
var IC2 = MAILBOX_64[C2]
var ID2 = MAILBOX_64[D2]
var IE2 = MAILBOX_64[E2]
var IF2 = MAILBOX_64[F2]
var IG2 = MAILBOX_64[G2]
var IH2 = MAILBOX_64[H2]
var IA1 = MAILBOX_64[A1]
var IB1 = MAILBOX_64[B1]
var IC1 = MAILBOX_64[C1]
var ID1 = MAILBOX_64[D1]
var IE1 = MAILBOX_64[E1]
var IF1 = MAILBOX_64[F1]
var IG1 = MAILBOX_64[G1]
var IH1 = MAILBOX_64[H1]

var SQ_NUM_TO_NAME = map[int]string{
	IA1: "a1",
	IB1: "b1",
	IC1: "c1",
	ID1: "d1",
	IE1: "e1",
	IF1: "f1",
	IG1: "g1",
	IH1: "h1",
	IA2: "a2",
	IB2: "b2",
	IC2: "c2",
	ID2: "d2",
	IE2: "e2",
	IF2: "f2",
	IG2: "g2",
	IH2: "h2",
	IA3: "a3",
	IB3: "b3",
	IC3: "c3",
	ID3: "d3",
	IE3: "e3",
	IF3: "f3",
	IG3: "g3",
	IH3: "h3",
	IA4: "a4",
	IB4: "b4",
	IC4: "c4",
	ID4: "d4",
	IE4: "e4",
	IF4: "f4",
	IG4: "g4",
	IH4: "h4",
	IA5: "a5",
	IB5: "b5",
	IC5: "c5",
	ID5: "d5",
	IE5: "e5",
	IF5: "f5",
	IG5: "g5",
	IH5: "h5",
	IA6: "a6",
	IB6: "b6",
	IC6: "c6",
	ID6: "d6",
	IE6: "e6",
	IF6: "f6",
	IG6: "g6",
	IH6: "h6",
	IA7: "a7",
	IB7: "b7",
	IC7: "c7",
	ID7: "d7",
	IE7: "e7",
	IF7: "f7",
	IG7: "g7",
	IH7: "h7",
	IA8: "a8",
	IB8: "b8",
	IC8: "c8",
	ID8: "d8",
	IE8: "e8",
	IF8: "f8",
	IG8: "g8",
	IH8: "h8",
}

var SQ_NAME_TO_SQ_64 = map[int]int{
	IA1: 56,
	IB1: 57,
	IC1: 58,
	ID1: 59,
	IE1: 60,
	IF1: 61,
	IG1: 62,
	IH1: 63,
	IA2: 48,
	IB2: 49,
	IC2: 50,
	ID2: 51,
	IE2: 52,
	IF2: 53,
	IG2: 54,
	IH2: 55,
	IA3: 40,
	IB3: 41,
	IC3: 42,
	ID3: 43,
	IE3: 44,
	IF3: 45,
	IG3: 46,
	IH3: 47,
	IA4: 32,
	IB4: 33,
	IC4: 34,
	ID4: 35,
	IE4: 36,
	IF4: 37,
	IG4: 38,
	IH4: 39,
	IA5: 24,
	IB5: 25,
	IC5: 26,
	ID5: 27,
	IE5: 28,
	IF5: 29,
	IG5: 30,
	IH5: 31,
	IA6: 16,
	IB6: 17,
	IC6: 18,
	ID6: 19,
	IE6: 20,
	IF6: 21,
	IG6: 22,
	IH6: 23,
	IA7: 8,
	IB7: 9,
	IC7: 10,
	ID7: 11,
	IE7: 12,
	IF7: 13,
	IG7: 14,
	IH7: 15,
	IA8: 0,
	IB8: 1,
	IC8: 2,
	ID8: 3,
	IE8: 4,
	IF8: 5,
	IG8: 6,
	IH8: 7,
}

var IA8_TO_IH1 = [8][8]int{
	{
		IA8,
		IB8,
		IC8,
		ID8,
		IE8,
		IF8,
		IG8,
		IH8,
	},
	{
		IA7,
		IB7,
		IC7,
		ID7,
		IE7,
		IF7,
		IG7,
		IH7,
	},
	{
		IA6,
		IB6,
		IC6,
		ID6,
		IE6,
		IF6,
		IG6,
		IH6,
	},
	{
		IA5,
		IB5,
		IC5,
		ID5,
		IE5,
		IF5,
		IG5,
		IH5,
	},
	{
		IA4,
		IB4,
		IC4,
		ID4,
		IE4,
		IF4,
		IG4,
		IH4,
	},
	{
		IA3,
		IB3,
		IC3,
		ID3,
		IE3,
		IF3,
		IG3,
		IH3,
	},
	{
		IA2,
		IB2,
		IC2,
		ID2,
		IE2,
		IF2,
		IG2,
		IH2,
	},
	{
		IA1,
		IB1,
		IC1,
		ID1,
		IE1,
		IF1,
		IG1,
		IH1,
	},
}

func epSquareFile(ep *int) int {
	if ep == nil { // should return error
		return -1
	}

	sqString := SQ_NUM_TO_NAME[*ep]
	return strings.IndexByte(FILES, sqString[0])
}
