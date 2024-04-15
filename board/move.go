package board

import "fmt"

type Move struct {
	From            int
	To              int
	Capture         bool
	CastleKingside  bool
	CastleQueenside bool
	Promote         bool
	PromotionPiece  int
	DoublePawnPush  bool
}

func (m Move) IsNull() bool {
	if m.From == 0 && m.To == 0 {
		return true
	}
	return false
}

func (m Move) See(b *Board) int {
	return see(m.To, b)
}

func (m Move) Score(b *Board, bb bitboard) int {
	s := 0
	if m.Promote {
		s += 1000
	}
	if m.IsCheck(*b, bb) {
		s += 10000
	}
	if m.Capture {
		s += 100
		mp := b.PieceAt(m.From)
		cp := b.PieceAt(m.To)
		s += CAPTURE_SCORE[mp][cp] * 100
	}
	if m.CastleKingside || m.CastleQueenside {
		s += 100
	}
	if m.DoublePawnPush && 2 < file(m.To) && file(m.To) < 6 {
		s += 100
	}
	return s
}

func (m Move) IsCheck(b Board, bb bitboard) bool {
	piece := m.PromotionPiece
	startSq := m.To
	var kingSq int
	if b.Side == WHITE {
		kingSq = BB_TO_BOARDSQUARE[deBruijnLSB(bb.blackking)]
	} else {
		kingSq = BB_TO_BOARDSQUARE[deBruijnLSB(bb.whiteking)]
	}
	return attacksSquare(bb, b.Side, piece, startSq, kingSq)
}

func (m Move) Print() {
	fmt.Printf("%s", m.ToString())
}

func (m Move) ToString() string {
	return fmt.Sprintf("%s - %s\n", SQ_NUM_TO_NAME[m.From], SQ_NUM_TO_NAME[m.To])
}

func EqualMoves(m1, m2 Move) bool {
	if m1.From != m2.From {
		return false
	}
	if m1.To != m2.To {
		return false
	}
	if m1.Capture != m2.Capture {
		return false
	}
	if m1.CastleKingside != m2.CastleKingside {
		return false
	}
	if m1.CastleQueenside != m2.CastleQueenside {
		return false
	}
	if m1.Promote != m2.Promote {
		return false
	}
	if m1.PromotionPiece != m2.PromotionPiece {
		return false
	}
	if m1.DoublePawnPush != m2.DoublePawnPush {
		return false
	}
	return true
}

var CAPTURE_SCORE = map[int]map[int]int{
	WHITE_PAWN: map[int]int{
		BLACK_QUEEN:  1000,
		BLACK_ROOK:   900,
		BLACK_BISHOP: 800,
		BLACK_KNIGHT: 800,
		BLACK_PAWN:   100,
	},
	BLACK_PAWN: map[int]int{
		WHITE_QUEEN:  1000,
		WHITE_ROOK:   900,
		WHITE_BISHOP: 800,
		WHITE_KNIGHT: 800,
		WHITE_PAWN:   100,
	},
	WHITE_KNIGHT: map[int]int{
		BLACK_QUEEN:  1000,
		BLACK_ROOK:   900,
		BLACK_BISHOP: 100,
		BLACK_KNIGHT: 100,
		BLACK_PAWN:   50,
	},
	BLACK_KNIGHT: map[int]int{
		WHITE_QUEEN:  1000,
		WHITE_ROOK:   900,
		WHITE_BISHOP: 100,
		WHITE_KNIGHT: 100,
		WHITE_PAWN:   50,
	},
	WHITE_BISHOP: map[int]int{
		BLACK_QUEEN:  1000,
		BLACK_ROOK:   900,
		BLACK_BISHOP: 100,
		BLACK_KNIGHT: 100,
		BLACK_PAWN:   50,
	},
	BLACK_BISHOP: map[int]int{
		WHITE_QUEEN:  1000,
		WHITE_ROOK:   900,
		WHITE_BISHOP: 100,
		WHITE_KNIGHT: 100,
		WHITE_PAWN:   50,
	},
	WHITE_ROOK: map[int]int{
		BLACK_QUEEN:  1000,
		BLACK_ROOK:   100,
		BLACK_BISHOP: 50,
		BLACK_KNIGHT: 50,
		BLACK_PAWN:   25,
	},
	BLACK_ROOK: map[int]int{
		WHITE_QUEEN:  1000,
		WHITE_ROOK:   100,
		WHITE_BISHOP: 50,
		WHITE_KNIGHT: 50,
		WHITE_PAWN:   25,
	},
	WHITE_QUEEN: map[int]int{
		BLACK_QUEEN:  100,
		BLACK_ROOK:   10,
		BLACK_BISHOP: 10,
		BLACK_KNIGHT: 10,
		BLACK_PAWN:   5,
	},
	BLACK_QUEEN: map[int]int{
		WHITE_QUEEN:  100,
		WHITE_ROOK:   10,
		WHITE_BISHOP: 10,
		WHITE_KNIGHT: 10,
		WHITE_PAWN:   5,
	},
	WHITE_KING: map[int]int{
		BLACK_QUEEN:  100,
		BLACK_ROOK:   10,
		BLACK_BISHOP: 10,
		BLACK_KNIGHT: 10,
		BLACK_PAWN:   5,
	},
	BLACK_KING: map[int]int{
		WHITE_QUEEN:  100,
		WHITE_ROOK:   10,
		WHITE_BISHOP: 10,
		WHITE_KNIGHT: 10,
		WHITE_PAWN:   5,
	},
}
