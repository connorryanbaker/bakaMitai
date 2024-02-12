package board

import "fmt"

type BB uint64
const INIT_BLACK_PAWN_BB BB = 0b0000000011111111000000000000000000000000000000000000000000000000
const INIT_WHITE_PAWN_BB BB = 0b0000000000000000000000000000000000000000000000001111111100000000
const BOARDMASK BB = 0xFFFFFFFFFFFFFFFF

// LEFT SHIFTS
const NORTHEAST = 9
const NORTH = 8
const NORTHWEST = 7
const EAST = 1
// RIGHT SHIFTS
const WEST = -1
const SOUTHEAST = -7
const SOUTH = -8
const SOUTHWEST = -9

var AFILE BB = 0b0000000100000001000000010000000100000001000000010000000100000001
var BFILE BB = 0b0000001000000010000000100000001000000010000000100000001000000010
var CFILE BB = 0b0000010000000100000001000000010000000100000001000000010000000100
var DFILE BB = 0b0000100000001000000010000000100000001000000010000000100000001000
var EFILE BB = 0b0001000000010000000100000001000000010000000100000001000000010000
var FFILE BB = 0b0010000000100000001000000010000000100000001000000010000000100000
var GFILE BB = 0b0100000001000000010000000100000001000000010000000100000001000000
var HFILE BB = 0b1000000010000000100000001000000010000000100000001000000010000000
var RANK1 BB = 0b0000000000000000000000000000000000000000000000000000000011111111
var RANK2 BB = 0b0000000000000000000000000000000000000000000000001111111100000000
var RANK3 BB = 0b0000000000000000000000000000000000000000111111110000000000000000
var RANK4 BB = 0b0000000000000000000000000000000011111111000000000000000000000000
var RANK5 BB = 0b0000000000000000000000001111111100000000000000000000000000000000
var RANK6 BB = 0b0000000000000000111111110000000000000000000000000000000000000000
var RANK7 BB = 0b0000000011111111000000000000000000000000000000000000000000000000
var RANK8 BB = 0b1111111100000000000000000000000000000000000000000000000000000000

type bitboard struct {
  whitepawns BB
  blackpawns BB
}

func NewBitboard() bitboard {
  return bitboard{
    whitepawns: INIT_WHITE_PAWN_BB,
    blackpawns: INIT_BLACK_PAWN_BB,
  }
}

func (bb bitboard) emptySquares() BB {
  return BOARDMASK ^ bb.allPieces()
}

// update below as more pieces added 
func (bb bitboard) whitePieces() BB {
  return bb.whitepawns
}

func (bb bitboard) blackPieces() BB {
  return bb.blackpawns
}

func (bb bitboard) allPieces() BB {
  return bb.whitePieces() | bb.blackPieces()
}

func (bb bitboard) pushOneWhitePawns() BB {
  return shiftBB(bb.emptySquares(), SOUTH) & bb.whitepawns
}

func (bb bitboard) pushOneBlackPawns() BB {
  return shiftBB(bb.emptySquares(), NORTH) & bb.blackpawns
}

func (bb bitboard) pushTwoWhitePawns() BB {
  var fourthRank BB = 0x00000000FF000000
  emptyThirdRank := shiftBB(fourthRank & bb.emptySquares(), SOUTH) & bb.emptySquares()
  return shiftBB(emptyThirdRank, SOUTH) & bb.whitepawns
}

func (bb bitboard) pushTwoBlackPawns() BB {
  var fifthRank BB = 0x000000FF00000000
  emptySixthRank := shiftBB(fifthRank & bb.emptySquares(), NORTH) & bb.emptySquares()
  return shiftBB(emptySixthRank, NORTH) & bb.blackpawns
}

func (bb bitboard) whitePawnAttacks() BB {
  return bb.whitePawnWestAttacks() | bb.whitePawnEastAttacks()
}

func (bb bitboard) whitePawnWestAttacks() BB {
  return shiftBB(bb.whitepawns, NORTHWEST) & ^HFILE
}

func (bb bitboard) whitePawnEastAttacks() BB {
  return shiftBB(bb.whitepawns, NORTHEAST) & ^AFILE
}

func (bb bitboard) blackPawnAttacks() BB {
  return bb.blackPawnWestAttacks() | bb.blackPawnEastAttacks()
}

func (bb bitboard) blackPawnWestAttacks() BB {
  return shiftBB(bb.blackpawns, SOUTHWEST) & ^HFILE
}

func (bb bitboard) blackPawnEastAttacks() BB {
  return shiftBB(bb.blackpawns, SOUTHEAST) & ^AFILE
}

func shiftBB(bb BB, d int) BB {
  if d < 0 {
    s := d * -1
    return bb >> s
  }
  return bb << d
}

func printBB(b BB) {
  var sqs [8][8]int

  m := BB(1)
  f := 0
  r := 7

  for i := 0; i < 64; i++ {
    if b & (m << i) != BB(0) {
      sqs[r][f] = 1
    }
    f += 1
    if f == 8 {
      r -= 1
      f = 0
    }
  }

  for i := 0; i < 8; i++ {
    fmt.Println(sqs[i])
  }
  fmt.Printf("\n")
}
