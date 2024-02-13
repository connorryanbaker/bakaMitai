package board

import "fmt"

type BB uint64

const INIT_BLACK_PAWN_BB BB = 0b0000000011111111000000000000000000000000000000000000000000000000
const INIT_WHITE_PAWN_BB BB = 0b0000000000000000000000000000000000000000000000001111111100000000
const INIT_BLACK_KNIGHT_BB BB = 0b0100001000000000000000000000000000000000000000000000000000000000
const INIT_WHITE_KNIGHT_BB BB = 0b0000000000000000000000000000000000000000000000000000000001000010
const INIT_BLACK_KING_BB BB = 0b0001000000000000000000000000000000000000000000000000000000000000
const INIT_WHITE_KING_BB BB = 0b0000000000000000000000000000000000000000000000000000000000001000
const BOARDMASK BB = 0xFFFFFFFFFFFFFFFF

const AFILE BB = 0b0000000100000001000000010000000100000001000000010000000100000001
const BFILE BB = 0b0000001000000010000000100000001000000010000000100000001000000010
const CFILE BB = 0b0000010000000100000001000000010000000100000001000000010000000100
const DFILE BB = 0b0000100000001000000010000000100000001000000010000000100000001000
const EFILE BB = 0b0001000000010000000100000001000000010000000100000001000000010000
const FFILE BB = 0b0010000000100000001000000010000000100000001000000010000000100000
const GFILE BB = 0b0100000001000000010000000100000001000000010000000100000001000000
const HFILE BB = 0b1000000010000000100000001000000010000000100000001000000010000000
const RANK1 BB = 0b0000000000000000000000000000000000000000000000000000000011111111
const RANK2 BB = 0b0000000000000000000000000000000000000000000000001111111100000000
const RANK3 BB = 0b0000000000000000000000000000000000000000111111110000000000000000
const RANK4 BB = 0b0000000000000000000000000000000011111111000000000000000000000000
const RANK5 BB = 0b0000000000000000000000001111111100000000000000000000000000000000
const RANK6 BB = 0b0000000000000000111111110000000000000000000000000000000000000000
const RANK7 BB = 0b0000000011111111000000000000000000000000000000000000000000000000
const RANK8 BB = 0b1111111100000000000000000000000000000000000000000000000000000000

const NORTHEAST = 9
const NORTH = 8
const NORTHWEST = 7
const EAST = 1
const WEST = -1
const SOUTHEAST = -7
const SOUTH = -8
const SOUTHWEST = -9

type bitboard struct {
	whitepawns   BB
	whiteknights BB
	whiteking    BB
	blackpawns   BB
	blackknights BB
	blackking    BB
}

func NewBitboard() bitboard {
	return bitboard{
		whitepawns:   INIT_WHITE_PAWN_BB,
		whiteknights: INIT_WHITE_KNIGHT_BB,
		whiteking:    INIT_WHITE_KING_BB,
		blackpawns:   INIT_BLACK_PAWN_BB,
		blackknights: INIT_BLACK_KNIGHT_BB,
		blackking:    INIT_BLACK_KING_BB,
	}
}

func (bb bitboard) emptySquares() BB {
	return BOARDMASK ^ bb.allPieces()
}

// update below as more pieces added
func (bb bitboard) whitePieces() BB {
	return bb.whitepawns | bb.whiteknights
}

func (bb bitboard) blackPieces() BB {
	return bb.blackpawns | bb.blackknights
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
	emptyThirdRank := shiftBB(fourthRank&bb.emptySquares(), SOUTH) & bb.emptySquares()
	return shiftBB(emptyThirdRank, SOUTH) & bb.whitepawns
}

func (bb bitboard) pushTwoBlackPawns() BB {
	var fifthRank BB = 0x000000FF00000000
	emptySixthRank := shiftBB(fifthRank&bb.emptySquares(), NORTH) & bb.emptySquares()
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

func (bb bitboard) whiteKingMoves() BB {
	return bb.kingMoves(bb.whiteking) & (bb.emptySquares() | bb.blackPieces()) // todo: taboo
}

func (bb bitboard) blackKingMoves() BB {
	return bb.kingMoves(bb.blackking) & (bb.emptySquares() | bb.whitePieces())
}

func (bb bitboard) kingMoves(k BB) BB {
	return shiftBB(k & ^AFILE, NORTHWEST) |
		shiftBB(k & ^AFILE, WEST) |
		shiftBB(k & ^AFILE, SOUTHWEST) |
		shiftBB(k, NORTH) |
		shiftBB(k, SOUTH) |
		shiftBB(k & ^HFILE, NORTHEAST) |
		shiftBB(k & ^HFILE, EAST) |
		shiftBB(k & ^HFILE, SOUTHEAST)
}

// KNIGHT OFFSETS
const NORTHNORTHWEST = 15
const NORTHNORTHEAST = 17
const NORTHWESTWEST = 6
const NORTHEASTEAST = 10
const SOUTHWESTWEST = -10
const SOUTHEASTEAST = -6
const SOUTHSOUTHWEST = -17
const SOUTHSOUTHEAST = -15

func (bb bitboard) whiteKnightMoves() BB {
	return bb.knightMoves(bb.whiteknights) & (bb.emptySquares() | bb.blackPieces())
}

func (bb bitboard) blackKnightMoves() BB {
	return bb.knightMoves(bb.blackknights) & (bb.emptySquares() | bb.whitePieces())
}

func (bb bitboard) knightMoves(k BB) BB {
	return shiftBB(k&(^GFILE & ^HFILE), NORTHEASTEAST) |
		shiftBB(k & ^HFILE, NORTHNORTHEAST) |
		shiftBB(k&(^GFILE & ^HFILE), SOUTHEASTEAST) |
		shiftBB(k & ^HFILE, SOUTHSOUTHEAST) |
		shiftBB(k & ^AFILE, NORTHNORTHWEST) |
		shiftBB(k&(^AFILE & ^BFILE), NORTHWESTWEST) |
		shiftBB(k&(^AFILE & ^BFILE), SOUTHWESTWEST) |
		shiftBB(k & ^AFILE, SOUTHSOUTHWEST)
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
		if b&(m<<i) != BB(0) {
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
