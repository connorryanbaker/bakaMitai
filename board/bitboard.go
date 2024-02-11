package board

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

func shiftBB(bb BB, d int) BB {
  if d < 0 {
    return bb >> (d * -1)
  }
  return bb << d
}

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
