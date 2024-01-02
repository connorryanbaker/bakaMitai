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
  b.pieces = parsePieceString(components[0])
  // b.side = parseSideToMove(components[1])
  // b.castle = parseCastlePermissions(components[2])
  // b.ep = parseEnPassant(components[3])
  // b.hply = parseHply(components[4])
  // b.ply = parsePly(components[1], components[5])
  return b
}

func parsePieceString(s string) [120]int {
  b := emptyPiecesArray()
  p := strings.Split(s, "/")
  i := 0
  re := regexp.MustCompile(`\d`)

  for _, rank := range p {
    for _, c := range rank {
      if re.MatchString(string(c)) {
        v, err := strconv.Atoi(string(c))
        if err != nil {
          panic(err)
        }
        i += v
      } else {
        b[MAILBOX_64[i]] = FEN_TO_PIECE[string(c)]
        i += 1
      }
    }
  }

  return b
}



