package board

import (
	"regexp"
	"strconv"
	"strings"
)

// components:
// piece placement
// side to move
// castling ability
// ep square
// halfmove clock (reset to 0 after capture / pawn move) for 50 move draw
// fullmove counter

// example:
// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
// r1bqkb1r/pppppppp/1n3n2/8/8/1N3N2/PPPPPPPP/R1BQKB1R w KQkq - 14 8
// rn1qk1nr/pp2ppbp/2pp2p1/8/2PP4/2N2P2/PP2BPPP/R1BQ1RK1 b kq - 3 7
// 3k4/8/8/8/8/8/8/R3K3 w Q - 0 1
// 8/8/1k6/2b5/2pP4/8/5K2/8 b - d3 0 1

var FEN_TO_PIECE = map[string]int{
	"r": BLACK_ROOK,
	"n": BLACK_KNIGHT,
	"b": BLACK_BISHOP,
	"q": BLACK_QUEEN,
	"k": BLACK_KING,
	"p": BLACK_PAWN,
	"R": WHITE_ROOK,
	"N": WHITE_KNIGHT,
	"B": WHITE_BISHOP,
	"Q": WHITE_QUEEN,
	"K": WHITE_KING,
	"P": WHITE_PAWN,
}

func FromFENString(f string) Board {
	b := Board{}
	components := strings.Split(f, " ")
	pieces, pieceSquares := parsePieceString(components[0])
	b.pieces = pieces
	b.pieceSquares = pieceSquares
	b.side = parseSideToMove(components[1])
	b.castle = parseCastlePermissions(components[2])
	b.ep = parseEnPassant(components[3])
	b.hply = parseHply(components[4])
	b.ply = parsePly(components[5])
	return b
}

func parsePieceString(s string) ([120]int, map[int][]int) {
	b := emptyPiecesArray()
	pieceSquares := make(map[int][]int)
	p := strings.Split(s, "/")
	i := 0
	re := regexp.MustCompile(`\d`)

	for _, rank := range p {
		for _, c := range rank {
			if re.MatchString(string(c)) { // TODO: cleanup coercion
				v, err := strconv.Atoi(string(c))
				if err != nil {
					panic(err)
				}
				i += v
			} else {
				p := FEN_TO_PIECE[string(c)]
				sq := MAILBOX_64[i]
				b[sq] = p
				sqs, ok := pieceSquares[p]
				if !ok {
					a := make([]int, 1)
					a[0] = sq
					pieceSquares[p] = a
				} else {
					pieceSquares[p] = append(sqs, sq)
				}
				i += 1
			}
		}
	}

	return b, pieceSquares
}

func parseSideToMove(s string) int {
	if s == "w" {
		return 0
	}
	return 1
}

func parseCastlePermissions(s string) [4]bool {
	b := [4]bool{false, false, false, false}
	keys := "KQkq"
	for i, c := range keys {
		if strings.ContainsRune(s, c) {
			b[i] = true
		}
	}
	return b
}

func parseEnPassant(s string) *int {
	if s == "-" {
		return nil
	}

	var rankToMailboxStart = map[byte]int{
		'1': 91,
		'2': 81,
		'3': 71,
		'4': 61,
		'5': 51,
		'6': 41,
		'7': 31,
		'8': 21,
	}

	v := rankToMailboxStart[s[1]] + strings.IndexByte("abcdefgh", s[0])
	return &v
}

func parseHply(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parsePly(move string) int {
	n, err := strconv.Atoi(move)
	if err != nil {
		panic(err)
	}

	return n
}
