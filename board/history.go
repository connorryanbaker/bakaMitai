package board

import (
	"fmt"
	"strings"
)

type History struct {
	previousSquareOccupant int
	Move                   Move
	castle                 [4]bool
	ep                     *int
	hply                   int
	ply                    int
	Hash                   uint64
}

func (b *Board) pushHistory(m Move) {
	h := History{
		previousSquareOccupant: b.PieceAt(m.To),
		Move:                   m,
		castle:                 b.Castle,
		ep:                     b.Ep,
		hply:                   b.Hply,
		ply:                    b.Ply,
		Hash:                   b.Hash(),
	}
	b.History = append(b.History, h)
}

func (b *Board) popHistory() {
	b.History[len(b.History)-1] = History{}
	b.History = b.History[:len(b.History)-1]
}

func (b *Board) MoveHistory() []Move {
	m := make([]Move, len(b.History))
	for i := 0; i < len(b.History); i++ {
		m[i] = b.History[i].Move
	}
	return m
}

var PGN_PIECE_TRANSLATION = map[int]string{
	WHITE_KNIGHT: "N",
	BLACK_KNIGHT: "N",
	WHITE_BISHOP: "B",
	BLACK_BISHOP: "B",
	WHITE_ROOK:   "R",
	BLACK_ROOK:   "R",
	WHITE_QUEEN:  "Q",
	BLACK_QUEEN:  "Q",
	WHITE_KING:   "K",
	BLACK_KING:   "K",
}

func ToPGN(h []History) string {
	moves := make([]string, 0)
	mb := [2]string{"", ""}
	b := NewBoard()
	mn := 1

	for i := 0; i < len(h); i++ {
		m := h[i].Move
		p := b.PieceAt(m.From)
		switch p {
		case WHITE_PAWN, BLACK_PAWN:
			mb[i%2] = parsePawnMove(b, m)
		case WHITE_BISHOP, BLACK_BISHOP, WHITE_QUEEN, BLACK_QUEEN:
			mb[i%2] = parseQBMove(b, m)
		case WHITE_KNIGHT, BLACK_KNIGHT:
			mb[i%2] = parseQBMove(b, m)
		case WHITE_ROOK, BLACK_ROOK:
			mb[i%2] = parseQBMove(b, m)
		case WHITE_KING, BLACK_KING:
			mb[i%2] = parseKingMove(b, m)
		}
		b.MakeMove(m)

		if i%2 != 0 || i == len(h)-1 {
			moves = append(moves, fmt.Sprintf("%d. %s %s ", mn, mb[0], mb[1]))
			mn += 1
		}
	}

	return strings.Join(moves, " ")
}

func parsePawnMove(b Board, m Move) string {
	if m.Promote {
		if m.Capture {
			return fmt.Sprintf("%bx%s%s", fileName(m.From), SQ_NUM_TO_NAME[m.To], PGN_PIECE_TRANSLATION[m.PromotionPiece])
		}
		return fmt.Sprintf("%s%s", SQ_NUM_TO_NAME[m.To], PGN_PIECE_TRANSLATION[m.PromotionPiece])
	}
	if m.Capture {
		return fmt.Sprintf("%sx%s", string(fileName(m.From)), SQ_NUM_TO_NAME[m.To])
	}
	return SQ_NUM_TO_NAME[m.To]
}

func parseQBMove(b Board, m Move) string {
	if m.Capture {
		return fmt.Sprintf("%sx%s", PGN_PIECE_TRANSLATION[b.PieceAt(m.From)], SQ_NUM_TO_NAME[m.To])
	}
	return fmt.Sprintf("%s%s", PGN_PIECE_TRANSLATION[b.PieceAt(m.From)], SQ_NUM_TO_NAME[m.To])
}

func parseKingMove(b Board, m Move) string {
	if m.CastleKingside {
		return "0-0"
	}
	if m.CastleQueenside {
		return "0-0-0"
	}
	if m.Capture {
		return fmt.Sprintf("%sx%s", PGN_PIECE_TRANSLATION[b.PieceAt(m.From)], SQ_NUM_TO_NAME[m.To])
	}
	return fmt.Sprintf("%s%s", PGN_PIECE_TRANSLATION[b.PieceAt(m.From)], SQ_NUM_TO_NAME[m.To])
}

// TODO: update knight and rook fns to check for
// situation where move needs to be qualified with file/rank
// for now, these functions are just duplicated

// func parseKnightMove(b Board, m Move) string {
// TODO
// }

// func parseRookMove(b Board, m Move) string {
// TODO
// }
