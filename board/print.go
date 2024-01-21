package board

import "fmt"

const BLACK_ESCAPE = "\033[0;30m"
const WHITE_ESCAPE = "\033[0;37m"
const RESET_BGCOLOR_ESCAPE = "\033[0m"

var PIECE_TO_CODE_POINT = map[int]rune{
	BLACK_PAWN:   '\u2659',
	BLACK_KNIGHT: '\u2658',
	BLACK_BISHOP: '\u2657',
	BLACK_ROOK:   '\u2656',
	BLACK_QUEEN:  '\u2655',
	BLACK_KING:   '\u2654',
	WHITE_PAWN:   '\u265F',
	WHITE_KNIGHT: '\u265E',
	WHITE_BISHOP: '\u265D',
	WHITE_ROOK:   '\u265C',
	WHITE_QUEEN:  '\u265B',
	WHITE_KING:   '\u265A',
	EMPTY_SQUARE: ' ',
}

// const CLEAR_ESCAPE = "\033[2J"

func (b Board) Print() {
	// fmt.Printf(CLEAR_ESCAPE)
	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			p := b.PieceAt(IA8_TO_IH1[r][f])
			fmt.Printf("%s %s %s", escapeSequence(r, f), string(PIECE_TO_CODE_POINT[p]), RESET_BGCOLOR_ESCAPE)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n\n\n")
}

func escapeSequence(r, f int) string {
	if r%2 != 0 && f%2 == 0 {
		return "\033[45m"
	} else if r%2 == 0 && f%2 != 0 {
		return "\033[45m"
	} else {
		return "\033[46m"
	}
}
