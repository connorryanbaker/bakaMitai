package board

import (
	"fmt"
	"sort"
)

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

type bitboard struct {
	whitepawns   BB
	whiteknights BB
	whiteking    BB
	whiterooks   BB
	whitebishops BB
	whitequeens  BB
	blackpawns   BB
	blackknights BB
	blackking    BB
	blackrooks   BB
	blackbishops BB
	blackqueens  BB
}

func (bb bitboard) emptySquares() BB {
	return BOARDMASK ^ bb.allPieces()
}

// update below as more pieces added
func (bb bitboard) whitePieces() BB {
	return bb.whitepawns |
		bb.whiteknights |
		bb.whitebishops |
		bb.whiterooks |
		bb.whitequeens |
		bb.whiteking
}

func (bb bitboard) blackPieces() BB {
	return bb.blackpawns |
		bb.blackknights |
		bb.blackbishops |
		bb.blackrooks |
		bb.blackqueens |
		bb.blackking
}

func (bb bitboard) allPieces() BB {
	return bb.whitePieces() | bb.blackPieces()
}

func (bb bitboard) pushOneWhitePawns(pushMask, pinnedPieces, movesForPinned BB) BB {
	unpinnedPawnsCanPush := shiftBB(bb.emptySquares()&pushMask, SOUTH) & (bb.whitepawns & ^pinnedPieces)
	if pinnedPieces == 0 {
		return unpinnedPawnsCanPush
	}
	pinnedPawns := bb.whitepawns & pinnedPieces
	return unpinnedPawnsCanPush | (shiftBB(bb.emptySquares()&pushMask&movesForPinned, SOUTH) & pinnedPawns)

}

func (bb bitboard) pushOneBlackPawns(pushMask, pinnedPieces, movesForPinned BB) BB {
	unpinnedPawnsCanPush := shiftBB(bb.emptySquares()&pushMask, NORTH) & (bb.blackpawns & ^pinnedPieces)
	if pinnedPieces == 0 {
		return unpinnedPawnsCanPush
	}
	pinnedPawns := bb.blackpawns & pinnedPieces
	return unpinnedPawnsCanPush | (shiftBB(bb.emptySquares()&pushMask&movesForPinned, NORTH) & pinnedPawns)
}

func (bb bitboard) pushTwoWhitePawns(pushMask, pinnedPieces, movesForPinned BB) BB {
	var fourthRank BB = 0x00000000FF000000 & pushMask
	emptyThirdRank := shiftBB(fourthRank&bb.emptySquares(), SOUTH) & bb.emptySquares()
	unpinnedPawnsCanPush := shiftBB(emptyThirdRank, SOUTH) & (bb.whitepawns & ^pinnedPieces)
	if pinnedPieces == 0 {
		return unpinnedPawnsCanPush
	}
	fourthRank = 0x00000000FF000000 & pushMask & movesForPinned
	emptyThirdRank = shiftBB(fourthRank&bb.emptySquares()&movesForPinned, SOUTH) & (bb.emptySquares() & movesForPinned)
	pinnedPawns := bb.whitepawns & pinnedPieces
	return unpinnedPawnsCanPush | (shiftBB(emptyThirdRank, SOUTH) & pinnedPawns)
}

func (bb bitboard) pushTwoBlackPawns(pushMask, pinnedPieces, movesForPinned BB) BB {
	var fifthRank BB = 0x000000FF00000000 & pushMask
	emptySixthRank := shiftBB(fifthRank&bb.emptySquares(), NORTH) & bb.emptySquares()
	unpinnedPawnsCanPush := shiftBB(emptySixthRank, NORTH) & (bb.blackpawns & ^pinnedPieces)
	if pinnedPieces == 0 {
		return unpinnedPawnsCanPush
	}
	fifthRank = 0x000000FF00000000 & pushMask & movesForPinned
	emptySixthRank = shiftBB(fifthRank&bb.emptySquares()&movesForPinned, NORTH) & (bb.emptySquares() & movesForPinned)
	pinnedPawns := bb.blackpawns & pinnedPieces
	return unpinnedPawnsCanPush | (shiftBB(emptySixthRank, NORTH) & pinnedPawns)
}

func (bb bitboard) whitePawnsCaptureWest(captureMask, pinnedPieces, movesForPinned BB) BB {
	targets := shiftBB(bb.whitepawns, NORTHWEST) & (^HFILE & bb.blackPieces() & captureMask)
	unpinnedCaptures := shiftBB(targets, SOUTHEAST) & (bb.whitepawns & ^pinnedPieces)
	pinnedCaptures := shiftBB(targets&movesForPinned, SOUTHEAST) & (bb.whitepawns & pinnedPieces)
	return unpinnedCaptures | pinnedCaptures
}

func (bb bitboard) whitePawnsCaptureEast(captureMask, pinnedPieces, movesForPinned BB) BB {
	targets := shiftBB(bb.whitepawns, NORTHEAST) & (^AFILE & bb.blackPieces() & captureMask)
	unpinnedCaptures := shiftBB(targets, SOUTHWEST) & (bb.whitepawns & ^pinnedPieces)
	pinnedCaptures := shiftBB(targets&movesForPinned, SOUTHWEST) & (bb.whitepawns & pinnedPieces)
	return unpinnedCaptures | pinnedCaptures
}

func (bb bitboard) blackPawnsCaptureWest(captureMask, pinnedPieces, movesForPinned BB) BB {
	targets := shiftBB(bb.blackpawns, SOUTHWEST) & (^HFILE & bb.whitePieces() & captureMask)
	unpinnedCaptures := shiftBB(targets, NORTHEAST) & (bb.blackpawns & ^pinnedPieces)
	pinnedCaptures := shiftBB(targets&movesForPinned, NORTHEAST) & (bb.blackpawns & pinnedPieces)
	return unpinnedCaptures | pinnedCaptures
}

func (bb bitboard) blackPawnsCaptureEast(captureMask, pinnedPieces, movesForPinned BB) BB {
	targets := shiftBB(bb.blackpawns, SOUTHEAST) & (^AFILE & bb.whitePieces() & captureMask)
	unpinnedCaptures := shiftBB(targets, NORTHWEST) & (bb.blackpawns & ^pinnedPieces)
	pinnedCaptures := shiftBB(targets&movesForPinned, NORTHWEST) & (bb.blackpawns & pinnedPieces)
	return unpinnedCaptures | pinnedCaptures
}

func whitePawnAttacks(pawns BB, fpieces BB) BB {
	return (shiftBB(pawns, NORTHWEST) & (^HFILE | fpieces)) |
		(shiftBB(pawns, NORTHEAST) & (^AFILE | fpieces))
}

func blackPawnAttacks(pawns BB, fpieces BB) BB {
	return (shiftBB(pawns, SOUTHWEST) & (^HFILE | fpieces)) |
		(shiftBB(pawns, SOUTHEAST) & (^AFILE | fpieces))
}

func (bb bitboard) whiteKingMoves() BB {
	return bb.kingMoves(bb.whiteking) & (bb.emptySquares() | bb.blackPieces())
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

var KING_ATTACKS [64]BB

func initKingAttacks() {
	bb := bitboard{}
	for i := 0; i < 64; i++ {
		sq := BB(1 << i)
		KING_ATTACKS[i] = bb.kingMoves(sq)
	}
}

const NORTHEAST = 9
const NORTH = 8
const NORTHWEST = 7
const EAST = 1
const WEST = -1
const SOUTHEAST = -7
const SOUTH = -8
const SOUTHWEST = -9

const (
	NORTH_IDX     = iota
	NORTHEAST_IDX = iota
	EAST_IDX      = iota
	SOUTHEAST_IDX = iota
	SOUTH_IDX     = iota
	SOUTHWEST_IDX = iota
	WEST_IDX      = iota
	NORTHWEST_IDX = iota
)

var DELTA_IDXS = [8]int{NORTH_IDX, NORTHEAST_IDX, EAST_IDX, SOUTHEAST_IDX, SOUTH_IDX, SOUTHWEST_IDX, WEST_IDX, NORTHWEST_IDX}

var RAY_ATTACKS [8][64]BB

func initRayAttacks() {
	bb := bitboard{} // empty board

	for i := 0; i < 64; i++ {
		sq := BB(1 << i)
		for j := 0; j < 8; j++ {
			switch j {
			case 0:
				RAY_ATTACKS[j][i] = bb.fillNorth(sq) ^ sq
			case 1:
				RAY_ATTACKS[j][i] = bb.fillNorthEast(sq) ^ sq
			case 2:
				RAY_ATTACKS[j][i] = bb.fillEast(sq) ^ sq
			case 3:
				RAY_ATTACKS[j][i] = bb.fillSouthEast(sq) ^ sq
			case 4:
				RAY_ATTACKS[j][i] = bb.fillSouth(sq) ^ sq
			case 5:
				RAY_ATTACKS[j][i] = bb.fillSouthWest(sq) ^ sq
			case 6:
				RAY_ATTACKS[j][i] = bb.fillWest(sq) ^ sq
			case 7:
				RAY_ATTACKS[j][i] = bb.fillNorthWest(sq) ^ sq
			}
		}
	}
}

// starting w/ dumb7fil now
// todo: these probably shouldn't be methods
// will need to rethink bitboard struct purpose
// this can be xor'd w/ origin square to give east moves, etc.

// TODO: consolidate and generalize, duplication is unneccessary
func (bb bitboard) fillEast(p BB) BB {
	e := bb.emptySquares() & ^AFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, EAST))
	}
	return p
}

func fillEast(emptySquares, piece BB) BB {
	emptySquares &= ^AFILE
	for i := 0; i < 7; i++ {
		piece = piece | (emptySquares & shiftBB(piece, EAST))
	}
	return piece
}

func (bb bitboard) fillNorthEast(p BB) BB {
	e := bb.emptySquares() & ^AFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, NORTHEAST))
	}
	return p
}

func fillNorthEast(emptySquares, piece BB) BB {
	emptySquares &= ^AFILE
	for i := 0; i < 7; i++ {
		piece = piece | (emptySquares & shiftBB(piece, NORTHEAST))
	}
	return piece
}

func (bb bitboard) fillSouthEast(p BB) BB {
	e := bb.emptySquares() & ^AFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, SOUTHEAST))
	}
	return p
}

func fillSouthEast(emptySquares, piece BB) BB {
	emptySquares &= ^AFILE
	for i := 0; i < 7; i++ {
		piece = piece | (emptySquares & shiftBB(piece, SOUTHEAST))
	}
	return piece
}

func (bb bitboard) fillWest(p BB) BB {
	e := bb.emptySquares() & ^HFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, WEST))
	}
	return p
}

func fillWest(emptySquares, piece BB) BB {
	emptySquares &= ^HFILE
	for i := 0; i < 7; i++ {
		piece = piece | (emptySquares & shiftBB(piece, WEST))
	}
	return piece
}

func (bb bitboard) fillNorthWest(p BB) BB {
	e := bb.emptySquares() & ^HFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, NORTHWEST))
	}
	return p
}

func fillNorthWest(emptySquares, piece BB) BB {
	emptySquares &= ^HFILE
	for i := 0; i < 7; i++ {
		piece = piece | (emptySquares & shiftBB(piece, NORTHWEST))
	}
	return piece
}

func (bb bitboard) fillSouthWest(p BB) BB {
	e := bb.emptySquares() & ^HFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, SOUTHWEST))
	}
	return p
}

func fillSouthWest(emptySquares, piece BB) BB {
	emptySquares &= ^HFILE
	for i := 0; i < 7; i++ {
		piece = piece | (emptySquares & shiftBB(piece, SOUTHWEST))
	}
	return piece
}

func (bb bitboard) fillNorth(p BB) BB {
	e := bb.emptySquares()
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, NORTH))
	}
	return p
}

func fillNorth(emptySquares, piece BB) BB {
	for i := 0; i < 7; i++ {
		piece = piece | (emptySquares & shiftBB(piece, NORTH))
	}
	return piece
}

func (bb bitboard) fillSouth(p BB) BB {
	e := bb.emptySquares()
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, SOUTH))
	}
	return p
}

func fillSouth(emptySquares, piece BB) BB {
	for i := 0; i < 7; i++ {
		piece = piece | (emptySquares & shiftBB(piece, SOUTH))
	}
	return piece
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

var KNIGHT_ATTACKS [64]BB

func initKnightAttacks() {
	bb := bitboard{}
	for i := 0; i < 64; i++ {
		sq := BB(1 << i)
		KNIGHT_ATTACKS[i] = bb.knightMoves(sq)
	}
}

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

func init() {
	initKingAttacks()
	initKnightAttacks()
	initRayAttacks()
}

var MAILBOX_TO_BB = [99]int{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	56, 57, 58, 59, 60, 61, 62, 63, -1, -1,
	48, 49, 50, 51, 52, 53, 54, 55, -1, -1,
	40, 41, 42, 43, 44, 45, 46, 47, -1, -1,
	32, 33, 34, 35, 36, 37, 38, 39, -1, -1,
	24, 25, 26, 27, 28, 29, 30, 31, -1, -1,
	16, 17, 18, 19, 20, 21, 22, 23, -1, -1,
	8, 9, 10, 11, 12, 13, 14, 15, -1, -1,
	0, 1, 2, 3, 4, 5, 6, 7,
}

var BB_TO_BOARDSQUARE = [64]int{
	91, 92, 93, 94, 95, 96, 97, 98,
	81, 82, 83, 84, 85, 86, 87, 88,
	71, 72, 73, 74, 75, 76, 77, 78,
	61, 62, 63, 64, 65, 66, 67, 68,
	51, 52, 53, 54, 55, 56, 57, 58,
	41, 42, 43, 44, 45, 46, 47, 48,
	31, 32, 33, 34, 35, 36, 37, 38,
	21, 22, 23, 24, 25, 26, 27, 28,
}

// tranform piecelist to bitboards
func (b Board) newBitboard() *bitboard {
	return &bitboard{
		whitepawns:   pieceSquareToBB(b, WHITE_PAWN),
		whiteknights: pieceSquareToBB(b, WHITE_KNIGHT),
		whitebishops: pieceSquareToBB(b, WHITE_BISHOP),
		whiterooks:   pieceSquareToBB(b, WHITE_ROOK),
		whitequeens:  pieceSquareToBB(b, WHITE_QUEEN),
		whiteking:    pieceSquareToBB(b, WHITE_KING),
		blackpawns:   pieceSquareToBB(b, BLACK_PAWN),
		blackknights: pieceSquareToBB(b, BLACK_KNIGHT),
		blackbishops: pieceSquareToBB(b, BLACK_BISHOP),
		blackrooks:   pieceSquareToBB(b, BLACK_ROOK),
		blackqueens:  pieceSquareToBB(b, BLACK_QUEEN),
		blackking:    pieceSquareToBB(b, BLACK_KING),
	}
}

func pieceSquareToBB(b Board, p int) BB {
	bb := BB(0)
	if sqs, ok := b.PieceSquares[p]; ok {
		for _, sq := range sqs {
			bb |= BB(1 << MAILBOX_TO_BB[sq])
		}
	}
	return bb
}

func popCount(x BB) int {
	count := 0
	for x > 0 {
		count += 1
		x &= (x - 1)
	}
	return count
}

// for each piece, gen moves

func (b Board) inCheck(bb bitboard) bool {
	wp := bb.whitePieces()
	bp := bb.blackPieces()
	em := bb.emptySquares()

	if b.Side == WHITE {
		return attackers(bb, WHITE, bb.whiteking, em, wp, bp) > 0
	}
	return attackers(bb, BLACK, bb.blackking, em, bp, wp) > 0
}

func (b Board) doubleCheck(bb bitboard) bool {
	// TODO: organize to minimize these calls
	// could potentially return two values from inCheck
	wp := bb.whitePieces()
	bp := bb.blackPieces()
	em := bb.emptySquares()

	if b.Side == WHITE {
		return popCount(attackers(bb, WHITE, bb.whiteking, em, wp, bp)) > 1
	}
	return popCount(attackers(bb, BLACK, bb.blackking, em, bp, wp)) > 1
}

func attackers(bb bitboard, side int, piece, empty, fpieces, opieces BB) BB {
	var attackers BB
	sq := deBruijnLSB(piece)
	if side == WHITE {
		attackers |= KNIGHT_ATTACKS[sq] & bb.blackknights
		attackers |= bishopAttacks(sq, empty, fpieces, opieces) & bb.blackbishops
		attackers |= rookAttacks(sq, empty, fpieces, opieces) & bb.blackrooks
		attackers |= queenAttacks(sq, empty, fpieces, opieces) & bb.blackqueens
		attackers |= whitePawnAttacks(piece, fpieces) & bb.blackpawns
	} else {
		attackers |= KNIGHT_ATTACKS[sq] & bb.whiteknights
		attackers |= bishopAttacks(sq, empty, fpieces, opieces) & bb.whitebishops
		attackers |= rookAttacks(sq, empty, fpieces, opieces) & bb.whiterooks
		attackers |= queenAttacks(sq, empty, fpieces, opieces) & bb.whitequeens
		attackers |= blackPawnAttacks(piece, fpieces) & bb.whitepawns
	}
	return attackers
}

func (b Board) checkingPiecesMask(bb bitboard) BB {
	if b.Side == WHITE {
		return attackers(bb, WHITE, bb.whiteking, bb.emptySquares(), bb.whitePieces(), bb.blackPieces())
	}
	return attackers(bb, BLACK, bb.blackking, bb.emptySquares(), bb.blackPieces(), bb.whitePieces())
}

func blockingSquares(attackedSquare, attackingSquare BB) BB {
	attackedSq64 := deBruijnLSB(attackedSquare)
	for i := NORTH_IDX; i <= NORTHWEST_IDX; i++ {
		moveRay := RAY_ATTACKS[i][attackedSq64]
		if moveRay&attackingSquare > 0 {
			return generateBlockingMask(moveRay, attackingSquare, i)
		}
	}
	return BB(0)
}

func generateBlockingMask(moveRay, attackingSquare BB, dir int) BB {
	switch dir {
	case NORTH_IDX:
		oppUnionLsb := cutoffBitboard(moveRay&attackingSquare, true)
		return moveRay & ^fillNorth(BOARDMASK, shiftBB(oppUnionLsb, NORTH))
	case NORTHEAST_IDX:
		oppUnionLsb := cutoffBitboard(moveRay&attackingSquare, true)
		return moveRay & ^fillNorthEast(BOARDMASK, shiftBB(oppUnionLsb, NORTHEAST))
	case NORTHWEST_IDX:
		oppUnionLsb := cutoffBitboard(moveRay&attackingSquare, true)
		return moveRay & ^fillNorthWest(BOARDMASK, shiftBB(oppUnionLsb, NORTHWEST))
	case EAST_IDX:
		oppUnionLsb := cutoffBitboard(moveRay&attackingSquare, true)
		return moveRay & ^fillEast(BOARDMASK, shiftBB(oppUnionLsb, EAST))
	case SOUTH_IDX:
		oppUnionLsb := cutoffBitboard(moveRay&attackingSquare, false)
		return moveRay & ^fillSouth(BOARDMASK, shiftBB(oppUnionLsb, SOUTH))
	case SOUTHEAST_IDX:
		oppUnionLsb := cutoffBitboard(moveRay&attackingSquare, false)
		return moveRay & ^fillSouthEast(BOARDMASK, shiftBB(oppUnionLsb, SOUTHEAST))
	case SOUTHWEST_IDX:
		oppUnionLsb := cutoffBitboard(moveRay&attackingSquare, false)
		return moveRay & ^fillSouthWest(BOARDMASK, shiftBB(oppUnionLsb, SOUTHWEST))
	case WEST_IDX:
		oppUnionLsb := cutoffBitboard(moveRay&attackingSquare, false)
		return moveRay & ^fillWest(BOARDMASK, shiftBB(oppUnionLsb, WEST))
	}
	// TODO: raise error here?
	return BB(0)
}

func (b Board) absolutePinnedPiecesSideToMove(bb bitboard) BB {
	em := bb.emptySquares()
	wp := bb.whitePieces()
	bp := bb.blackPieces()

	// TODO: overlap of rays is incorrect
	// need to find overlap between corresponding rays NORTH - SOUTH, EAST - WEST etc
	// from kingsq, get queenTabooAttacks in each direction
	// get corresponding opposite rook / bishop attack
	// func generateRayTabooAttacks(lsb, rayIdx, shiftDelta int, empty, fpieces, opieces BB, fill func(e, p BB) BB) BB {
	if b.Side == WHITE {
		raysFromKing := slidingRayAttacksFromSquare(deBruijnLSB(bb.whiteking))
		pinningRays := allQueenAttacks(bb.blackqueens&raysFromKing, em, bp, wp) |
			allBishopAttacks(bb.blackbishops&raysFromKing, em, bp, wp) |
			allRookAttacks(bb.blackrooks&raysFromKing, em, bp, wp)
		return pinningRays & wp & raysFromKing
	}

	raysFromKing := slidingRayAttacksFromSquare(deBruijnLSB(bb.blackking))
	pinningRays := allQueenAttacks(bb.whitequeens&raysFromKing, em, wp, bp) |
		allBishopAttacks(bb.whitebishops&raysFromKing, em, wp, bp) |
		allRookAttacks(bb.whiterooks&raysFromKing, em, wp, bp)
	return pinningRays & bp & raysFromKing
}

// Moves for pinned pieces:
// remove pinned pieces from board
// how to differentiate pinning pieces from sliding pieces giving check?
// (replace king with queen?) and calculate moves
// from king square to pinning piece(s)

func (b Board) movesForPinnedPieces(bb bitboard, pinnedPieces BB) BB {
	em := bb.emptySquares()
	wp := bb.whitePieces()
	bp := bb.blackPieces()
	var moveRay BB
	if b.Side == WHITE {
		potentialPinningPieces := bb.blackqueens & bb.blackrooks & bb.blackbishops
		kingSq := deBruijnLSB(bb.whiteking)
		for i := NORTH_IDX; i < 8; i++ {
			if RAY_ATTACKS[i][kingSq]&pinnedPieces > 0 {
				rayAttack := queenAttacks(kingSq, em, wp & ^pinnedPieces, bp)
				if rayAttack&potentialPinningPieces > 0 {
					moveRay |= rayAttack
				}
			}
		}
	} else {
		potentialPinningPieces := bb.whitequeens & bb.whiterooks & bb.whitebishops
		kingSq := deBruijnLSB(bb.blackking)
		for i := NORTH_IDX; i < 8; i++ {
			if RAY_ATTACKS[i][kingSq]&pinnedPieces > 0 {
				rayAttack := queenAttacks(kingSq, em, bp & ^pinnedPieces, wp)
				if rayAttack&potentialPinningPieces > 0 {
					moveRay |= rayAttack
				}
			}
		}
	}
	return moveRay
}

func slidingRayAttacksFromSquare(sq int) BB {
	var ray BB
	for i := 0; i < 8; i++ {
		ray |= RAY_ATTACKS[i][sq]
	}
	return ray
}

func (b Board) GenerateBitboardMoves() []Move {
	bb := b.newBitboard()
	moves := make([]Move, 0)
	// TODO:
	// absolute pins
	// ep capture & promotions
	// castling
	captureMask := BB(0xFFFFFFFFFFFFFFFF)
	pushMask := BB(0xFFFFFFFFFFFFFFFF)
	movesForPinned := BB(0xFFFFFFFFFFFFFFFF)
	pinnedPieces := b.absolutePinnedPiecesSideToMove(*bb)
	if pinnedPieces > 0 {
		printBB(pinnedPieces)
		b.Print()
		movesForPinned = b.movesForPinnedPieces(*bb, pinnedPieces)
	}
	if b.inCheck(*bb) {
		if b.doubleCheck(*bb) {
			return b.generateBitboardKingMoves(*bb)
		}
		captureMask = b.checkingPiecesMask(*bb)
		if b.Side == WHITE {
			if captureMask&(bb.blackbishops|bb.blackqueens|bb.blackrooks) > 0 {
				pushMask = blockingSquares(bb.whiteking, captureMask)
			} else {
				pushMask = BB(0)
			}
		} else {
			if captureMask&(bb.whitebishops|bb.whitequeens|bb.whiterooks) > 0 {
				pushMask = blockingSquares(bb.blackking, captureMask)
			} else {
				pushMask = BB(0)
			}
		}
	}
	moves = append(moves, b.generateBitboardPawnMoves(*bb, captureMask&movesForPinned, pushMask, pinnedPieces, movesForPinned)...)
	moves = append(moves, b.generateBitboardKnightMoves(*bb, captureMask, pushMask, pinnedPieces)...)
	moves = append(moves, b.generateBitboardBishopMoves(*bb, captureMask, pushMask, pinnedPieces, movesForPinned)...)
	moves = append(moves, b.generateBitboardRookMoves(*bb, captureMask, pushMask, pinnedPieces, movesForPinned)...)
	moves = append(moves, b.generateBitboardQueenMoves(*bb, captureMask, pushMask, pinnedPieces, movesForPinned)...)
	moves = append(moves, b.generateBitboardKingMoves(*bb)...)

	sort.Slice(moves, func(i, j int) bool {
		return moves[i].Score(b) > moves[j].Score(b)
	})

	return moves
}

func (bb bitboard) tabooSquares(side int) BB {
	// TODO: atm these taboo squares don't take into account pieces protecting
	// pieces of the same color - will need some generate ray attack modification
	wp := bb.whitePieces()
	bp := bb.blackPieces()
	em := bb.emptySquares()
	if side == WHITE {
		wp ^= bb.whiteking // remove king to allow for xray squares to be marked taboo
		return blackPawnAttacks(bb.blackpawns, bp) |
			rayTabooAttacks(bb.blackqueens, em, bp, wp, queenTabooAttacks) |
			rayTabooAttacks(bb.blackrooks, em, bp, wp, rookTabooAttacks) |
			rayTabooAttacks(bb.blackbishops, em, bp, wp, bishopTabooAttacks) |
			allKnightTabooAttacks(bb.blackknights) |
			KING_ATTACKS[deBruijnLSB(bb.blackking)]
	}
	bp ^= bb.blackking
	return whitePawnAttacks(bb.whitepawns, wp) |
		rayTabooAttacks(bb.whitequeens, em, wp, bp, queenTabooAttacks) |
		rayTabooAttacks(bb.whiterooks, em, wp, bp, rookTabooAttacks) |
		rayTabooAttacks(bb.whitebishops, em, wp, bp, bishopTabooAttacks) |
		allKnightTabooAttacks(bb.whiteknights) |
		KING_ATTACKS[deBruijnLSB(bb.whiteking)]
}

func (b Board) generateBitboardQueenMoves(bb bitboard, captureMask, pushMask, pinnedPieces, movesForPinned BB) []Move {
	if b.Side == WHITE {
		return generateBitboardQueenMovesForSide(bb.whitequeens, bb.emptySquares(), bb.whitePieces(), bb.blackPieces(), captureMask, pushMask, pinnedPieces, movesForPinned, WHITE_QUEEN)
	}
	return generateBitboardQueenMovesForSide(bb.blackqueens, bb.emptySquares(), bb.blackPieces(), bb.whitePieces(), captureMask, pushMask, pinnedPieces, movesForPinned, BLACK_QUEEN)
}

func queenAttacks(sq int, empty, fpieces, opieces BB) BB {
	return bishopAttacks(sq, empty, fpieces, opieces) | rookAttacks(sq, empty, fpieces, opieces)
}

func queenTabooAttacks(sq int, empty, fpieces, opieces BB) BB {
	return bishopTabooAttacks(sq, empty, fpieces, opieces) | rookTabooAttacks(sq, empty, fpieces, opieces)
}

func allQueenAttacks(queens, empty, fpieces, opieces BB) BB {
	var attacks BB
	for queens > 0 {
		lsb := deBruijnLSB(queens)
		attacks |= queenAttacks(lsb, empty, fpieces, opieces)
		queens ^= BB(1 << lsb)
	}
	return attacks
}

func rayTabooAttacks(pieces, empty, fpieces, opieces BB, f func(s int, e, f, o BB) BB) BB {
	var attacks BB
	for pieces > 0 {
		lsb := deBruijnLSB(pieces)
		attacks |= f(lsb, empty, fpieces, opieces)
		pieces ^= BB(1 << lsb)
	}
	return attacks
}

func movesFromAttacks(sq, piece int, attacks, enemies BB) []Move {
	moves := make([]Move, 0)
	captures := attacks & enemies
	quietmoves := attacks & ^enemies
	for captures > 0 {
		lsb := deBruijnLSB(captures)
		dsq := BB_TO_BOARDSQUARE[lsb]
		moves = append(moves, Move{sq, dsq, true, false, false, false, piece, false})
		captures ^= BB(1 << lsb)
	}
	for quietmoves > 0 {
		lsb := deBruijnLSB(quietmoves)
		dsq := BB_TO_BOARDSQUARE[lsb]
		moves = append(moves, Move{sq, dsq, false, false, false, false, piece, false})
		quietmoves ^= BB(1 << lsb)
	}
	return moves
}

func generateBitboardQueenMovesForSide(queens, empty, fpieces, opieces, captureMask, pushMask, pinnedPieces, movesForPinned BB, piece int) []Move {
	moves := make([]Move, 0)
	for queens > 0 {
		lsb := deBruijnLSB(queens)
		attacks := queenAttacks(lsb, empty, fpieces, opieces) & (captureMask | pushMask)
		if BB(1<<lsb)&pinnedPieces > 0 {
			attacks &= movesForPinned
		}
		moves = append(moves, movesFromAttacks(BB_TO_BOARDSQUARE[lsb], piece, attacks, opieces)...)
		queens ^= BB(1 << lsb)
	}
	return moves
}

func (b Board) generateBitboardRookMoves(bb bitboard, captureMask, pushMask, pinnedPieces, movesForPinned BB) []Move {
	if b.Side == WHITE {
		return generateBitboardRookMovesForSide(bb.whiterooks, bb.emptySquares(), bb.whitePieces(), bb.blackPieces(), captureMask, pushMask, pinnedPieces, movesForPinned, WHITE_ROOK)
	}
	return generateBitboardRookMovesForSide(bb.blackrooks, bb.emptySquares(), bb.blackPieces(), bb.whitePieces(), captureMask, pushMask, pinnedPieces, movesForPinned, BLACK_ROOK)
}

func rookAttacks(sq int, empty, fpieces, opieces BB) BB {
	return generateRayAttacks(sq, NORTH_IDX, NORTH, empty, fpieces, opieces, fillNorth) |
		generateRayAttacks(sq, SOUTH_IDX, SOUTH, empty, fpieces, opieces, fillSouth) |
		generateRayAttacks(sq, WEST_IDX, WEST, empty, fpieces, opieces, fillWest) |
		generateRayAttacks(sq, EAST_IDX, EAST, empty, fpieces, opieces, fillEast)
}

func rookTabooAttacks(sq int, empty, fpieces, opieces BB) BB {
	return generateRayTabooAttacks(sq, NORTH_IDX, NORTH, empty, fpieces, opieces, fillNorth) |
		generateRayTabooAttacks(sq, SOUTH_IDX, SOUTH, empty, fpieces, opieces, fillSouth) |
		generateRayTabooAttacks(sq, WEST_IDX, WEST, empty, fpieces, opieces, fillWest) |
		generateRayTabooAttacks(sq, EAST_IDX, EAST, empty, fpieces, opieces, fillEast)
}

func allRookAttacks(rooks, empty, fpieces, opieces BB) BB {
	var attacks BB
	for rooks > 0 {
		lsb := deBruijnLSB(rooks)
		attacks |= rookAttacks(lsb, empty, fpieces, opieces)
		rooks ^= BB(1 << lsb)
	}
	return attacks
}

func generateBitboardRookMovesForSide(rooks, empty, fpieces, opieces, captureMask, pushMask, pinnedPieces, movesForPinned BB, piece int) []Move {
	moves := make([]Move, 0)
	for rooks > 0 {
		lsb := deBruijnLSB(rooks)
		attacks := rookAttacks(lsb, empty, fpieces, opieces) & (captureMask | pushMask)
		if BB(1<<lsb)&pinnedPieces > 0 {
			attacks &= movesForPinned
		}
		moves = append(moves, movesFromAttacks(BB_TO_BOARDSQUARE[lsb], piece, attacks, opieces)...)
		rooks ^= BB(1 << lsb)
	}
	return moves
}

func (b Board) generateBitboardBishopMoves(bb bitboard, captureMask, pushMask, pinnedPieces, movesForPinned BB) []Move {
	if b.Side == WHITE {
		return generateBitboardBishopMovesForSide(bb.whitebishops, bb.emptySquares(), bb.whitePieces(), bb.blackPieces(), captureMask, pushMask, pinnedPieces, movesForPinned, WHITE_BISHOP)
	}
	return generateBitboardBishopMovesForSide(bb.blackbishops, bb.emptySquares(), bb.blackPieces(), bb.whitePieces(), captureMask, pushMask, pinnedPieces, movesForPinned, BLACK_BISHOP)
}

func bishopAttacks(sq int, empty, fpieces, opieces BB) BB {
	return generateRayAttacks(sq, NORTHWEST_IDX, NORTHWEST, empty, fpieces, opieces, fillNorthWest) |
		generateRayAttacks(sq, SOUTHWEST_IDX, SOUTHWEST, empty, fpieces, opieces, fillSouthWest) |
		generateRayAttacks(sq, NORTHEAST_IDX, NORTHEAST, empty, fpieces, opieces, fillNorthEast) |
		generateRayAttacks(sq, SOUTHEAST_IDX, SOUTHEAST, empty, fpieces, opieces, fillSouthEast)
}

func bishopTabooAttacks(sq int, empty, fpieces, opieces BB) BB {
	return generateRayTabooAttacks(sq, NORTHWEST_IDX, NORTHWEST, empty, fpieces, opieces, fillNorthWest) |
		generateRayTabooAttacks(sq, SOUTHWEST_IDX, SOUTHWEST, empty, fpieces, opieces, fillSouthWest) |
		generateRayTabooAttacks(sq, NORTHEAST_IDX, NORTHEAST, empty, fpieces, opieces, fillNorthEast) |
		generateRayTabooAttacks(sq, SOUTHEAST_IDX, SOUTHEAST, empty, fpieces, opieces, fillSouthEast)
}

func allBishopAttacks(bishops, empty, fpieces, opieces BB) BB {
	var attacks BB
	for bishops > 0 {
		lsb := deBruijnLSB(bishops)
		attacks |= bishopAttacks(lsb, empty, fpieces, opieces)
		bishops ^= BB(1 << lsb)
	}
	return attacks
}

func generateBitboardBishopMovesForSide(bishops, empty, fpieces, opieces, captureMask, pushMask, pinnedPieces, movesForPinned BB, piece int) []Move {
	moves := make([]Move, 0)
	for bishops > 0 {
		lsb := deBruijnLSB(bishops)
		attacks := bishopAttacks(lsb, empty, fpieces, opieces) & (captureMask | pushMask)
		if BB(1<<lsb)&pinnedPieces > 0 {
			attacks &= movesForPinned
		}
		moves = append(moves, movesFromAttacks(BB_TO_BOARDSQUARE[lsb], piece, attacks, opieces)...)
		bishops ^= BB(1 << lsb)
	}
	return moves
}

// need piece, friendlyPieces, oppPieces BB, shift delta fill function(BB BB)BB
// yikes this fn should be broken up
func generateRayAttackMoves(lsb int, empty, fpieces, opieces BB, pieceType, rayIdx, shiftDelta int, fill func(e, p BB) BB) []Move {
	moves := make([]Move, 0)

	moveRay := generateRayAttacks(lsb, rayIdx, shiftDelta, empty, fpieces, opieces, fill)
	sq := BB_TO_BOARDSQUARE[lsb]
	for moveRay > 0 {
		moveLsb := deBruijnLSB(moveRay)
		isCapture := (opieces & BB(1<<moveLsb)) > 0
		destSq := BB_TO_BOARDSQUARE[moveLsb]
		moves = append(
			moves,
			Move{sq, destSq, isCapture, false, false, false, pieceType, false},
		)
		moveRay ^= BB(1 << moveLsb)
	}

	return moves
}

// fill need to take same shift approach for taboo squares to take protected
// pieces into account
func generateRayAttacks(lsb, rayIdx, shiftDelta int, empty, fpieces, opieces BB, fill func(e, p BB) BB) BB {
	moveRay := RAY_ATTACKS[rayIdx][lsb]
	friendlyUnion := moveRay & fpieces
	moveRay &= ^fill(BOARDMASK, cutoffBitboard(friendlyUnion, shiftDelta > 0))
	if moveRay == 0 {
		return moveRay
	}

	oppUnionLsb := cutoffBitboard(moveRay&opieces, shiftDelta > 0)
	return moveRay & ^fill(BOARDMASK, shiftBB(oppUnionLsb, shiftDelta))
}

func generateRayTabooAttacks(lsb, rayIdx, shiftDelta int, empty, fpieces, opieces BB, fill func(e, p BB) BB) BB {
	moveRay := RAY_ATTACKS[rayIdx][lsb]
	friendlyUnionLsb := cutoffBitboard(moveRay&fpieces, shiftDelta > 0)
	moveRay &= ^fill(BOARDMASK, shiftBB(friendlyUnionLsb, shiftDelta))
	// rm duplication
	oppUnionLsb := cutoffBitboard(moveRay&opieces, shiftDelta > 0)
	return moveRay & ^fill(BOARDMASK, shiftBB(oppUnionLsb, shiftDelta))
}

func cutoffBitboard(bb BB, fromLSB bool) BB {
	if bb == 0 {
		return 0
	}
	if fromLSB {
		return BB(1 << deBruijnLSB(bb))
	}
	return BB(1 << deBruijnMSB(bb))
}

func (b Board) generateBitboardKingMoves(bb bitboard) []Move {
	// TODO: castling
	taboo := bb.tabooSquares(b.Side)
	if b.Side == WHITE {
		return generateKingMovesForSide(
			bb.whiteking, bb.blackPieces(), bb.emptySquares(), taboo, WHITE_KING,
		)
	}
	return generateKingMovesForSide(
		bb.blackking, bb.whitePieces(), bb.emptySquares(), taboo, BLACK_KING,
	)
}

func kingAttacks(lsb int, fpieces BB) BB {
	return KING_ATTACKS[lsb] & ^fpieces
}

// TODO: update to pass attacks constant to reuse for king / knight
func generateKingMovesForSide(king, oppPieces, emptySqs, taboo BB, piece int) []Move {
	moves := make([]Move, 0)
	for king > 0 {
		lsb := deBruijnLSB(king)
		sq := BB_TO_BOARDSQUARE[lsb]
		captures := KING_ATTACKS[lsb] & oppPieces & ^taboo
		quietmoves := KING_ATTACKS[lsb] & emptySqs & ^taboo
		for captures > 0 {
			clsb := deBruijnLSB(captures)
			csq := BB_TO_BOARDSQUARE[clsb]
			moves = append(
				moves,
				Move{sq, csq, true, false, false, false, piece, false},
			)
			captures ^= BB(1 << clsb)
		}
		for quietmoves > 0 {
			qlsb := deBruijnLSB(quietmoves)
			qsq := BB_TO_BOARDSQUARE[qlsb]
			moves = append(
				moves,
				Move{sq, qsq, false, false, false, false, piece, false},
			)
			quietmoves ^= BB(1 << qlsb)
		}
		king ^= BB(1 << lsb)
	}
	return moves
}

func (b Board) generateBitboardKnightMoves(bb bitboard, captureMask, pushMask, pinnedPieces BB) []Move {
	if b.Side == WHITE {
		return generateKnightMovesForSide(
			bb.whiteknights, bb.blackPieces(), bb.emptySquares(), captureMask, pushMask, pinnedPieces, WHITE_KNIGHT,
		)
	}
	return generateKnightMovesForSide(
		bb.blackknights, bb.whitePieces(), bb.emptySquares(), captureMask, pushMask, pinnedPieces, BLACK_KNIGHT,
	)
}

func knightAttacks(lsb int, fpieces BB) BB {
	return KNIGHT_ATTACKS[lsb] & ^fpieces
}

func allKnightTabooAttacks(knights BB) BB {
	var attacks BB
	for knights > 0 {
		lsb := deBruijnLSB(knights)
		attacks |= KNIGHT_ATTACKS[lsb]
		knights ^= BB(1 << lsb)
	}
	return attacks
}

func allKnightAttacks(knights, fpieces BB) BB {
	var attacks BB
	for knights > 0 {
		lsb := deBruijnLSB(knights)
		attacks |= knightAttacks(lsb, fpieces)
		knights ^= BB(1 << lsb)
	}
	return attacks
}

func generateKnightMovesForSide(knights, oppPieces, emptySqs, captureMask, pushMask, pinnedPieces BB, piece int) []Move {
	moves := make([]Move, 0)
	knights &= ^pinnedPieces
	for knights > 0 {
		lsb := deBruijnLSB(knights & ^pinnedPieces)
		sq := BB_TO_BOARDSQUARE[lsb]
		captures := KNIGHT_ATTACKS[lsb] & oppPieces & captureMask
		quietmoves := KNIGHT_ATTACKS[lsb] & emptySqs & pushMask
		for captures > 0 {
			clsb := deBruijnLSB(captures)
			csq := BB_TO_BOARDSQUARE[clsb]
			moves = append(
				moves,
				Move{sq, csq, true, false, false, false, piece, false},
			)
			captures ^= BB(1 << clsb)
		}
		for quietmoves > 0 {
			qlsb := deBruijnLSB(quietmoves)
			qsq := BB_TO_BOARDSQUARE[qlsb]
			moves = append(
				moves,
				Move{sq, qsq, false, false, false, false, piece, false},
			)
			quietmoves ^= BB(1 << qlsb)
		}
		knights ^= BB(1 << lsb)
	}
	return moves
}

func (b Board) generateBitboardPawnMoves(bb bitboard, captureMask, pushMask, pinnedPieces, movesForPinned BB) []Move {
	moves := make([]Move, 0)
	if b.Side == WHITE {
		moves = append(
			moves,
			pawnMovesFromBB(bb.pushOneWhitePawns(pushMask, pinnedPieces, movesForPinned), -10, WHITE_PAWN, false, false)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.pushTwoWhitePawns(pushMask, pinnedPieces, movesForPinned), -20, WHITE_PAWN, false, true)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.whitePawnsCaptureWest(captureMask, pinnedPieces, movesForPinned), -11, WHITE_PAWN, true, false)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.whitePawnsCaptureEast(captureMask, pinnedPieces, movesForPinned), -9, WHITE_PAWN, true, false)...,
		)
	} else {
		moves = append(
			moves,
			pawnMovesFromBB(bb.pushOneBlackPawns(pushMask, pinnedPieces, movesForPinned), 10, BLACK_PAWN, false, false)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.pushTwoBlackPawns(pushMask, pinnedPieces, movesForPinned), 20, BLACK_PAWN, false, true)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.blackPawnsCaptureWest(captureMask, pinnedPieces, movesForPinned), 9, BLACK_PAWN, true, false)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.blackPawnsCaptureEast(captureMask, pinnedPieces, movesForPinned), 11, BLACK_PAWN, true, false)...,
		)
	}

	return moves
}

func pawnMovesFromBB(bb BB, delta, piece int, capture, doublePush bool) []Move {
	moves := make([]Move, 0)

	for bb > 0 {
		lsb := deBruijnLSB(bb)
		sq := BB_TO_BOARDSQUARE[lsb]
		// TODO: promotion checking
		moves = append(moves, Move{sq, sq + delta, capture, false, false, false, piece, doublePush})
		bb ^= BB(1 << lsb)
	}

	return moves
}

var deBruijnLSBIndex = [64]int{
	0, 1, 48, 2, 57, 49, 28, 3,
	61, 58, 50, 42, 38, 29, 17, 4,
	62, 55, 59, 36, 53, 51, 43, 22,
	45, 39, 33, 30, 24, 18, 12, 5,
	63, 47, 56, 27, 60, 41, 37, 16,
	54, 35, 52, 21, 44, 32, 23, 11,
	46, 26, 40, 15, 34, 20, 31, 10,
	25, 14, 19, 9, 13, 8, 7, 6,
}

var deBruijnMSBIndex = [64]int{
	0, 47, 1, 56, 48, 27, 2, 60,
	57, 49, 41, 37, 28, 16, 3, 61,
	54, 58, 35, 52, 50, 42, 21, 44,
	38, 32, 29, 23, 17, 11, 4, 62,
	46, 55, 26, 59, 40, 36, 15, 53,
	34, 51, 20, 43, 31, 22, 10, 45,
	25, 39, 14, 33, 19, 30, 9, 24,
	13, 18, 8, 12, 7, 6, 5, 63,
}

const deBruijnSeq BB = 0x03f79d71b4cb0a89

func deBruijnLSB(bb BB) int {
	return deBruijnLSBIndex[((bb&-bb)*deBruijnSeq)>>58]
}

func deBruijnMSB(bb BB) int {
	bb |= bb >> 1
	bb |= bb >> 2
	bb |= bb >> 4
	bb |= bb >> 8
	bb |= bb >> 16
	bb |= bb >> 32
	return deBruijnMSBIndex[(bb*deBruijnSeq)>>58]
}
