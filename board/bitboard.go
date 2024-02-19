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

type bitboard struct {
	whitepawns   BB
	whiteknights BB
	whiteking    BB
	whiterooks   BB
	whitebishops BB
	whitequeen   BB
	blackpawns   BB
	blackknights BB
	blackking    BB
	blackrooks   BB
	blackbishops BB
	blackqueen   BB
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
		bb.whitequeen |
		bb.whiteking
}

func (bb bitboard) blackPieces() BB {
	return bb.blackpawns |
		bb.blackknights |
		bb.blackbishops |
		bb.blackrooks |
		bb.blackqueen |
		bb.blackking
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

func (bb bitboard) whitePawnsCaptureWest() BB {
	targets := shiftBB(bb.whitepawns, NORTHWEST) & (^HFILE & bb.blackPieces())
	return shiftBB(targets, SOUTHEAST) & bb.whitepawns
}

func (bb bitboard) whitePawnsCaptureEast() BB {
	targets := shiftBB(bb.whitepawns, NORTHEAST) & (^AFILE & bb.blackPieces())
	return shiftBB(targets, SOUTHWEST) & bb.whitepawns
}

func (bb bitboard) blackPawnsCaptureWest() BB {
	targets := shiftBB(bb.blackpawns, SOUTHWEST) & (^HFILE & bb.whitePieces())
	return shiftBB(targets, NORTHEAST) & bb.blackpawns
}

func (bb bitboard) blackPawnsCaptureEast() BB {
	targets := shiftBB(bb.blackpawns, SOUTHEAST) & (^AFILE & bb.whitePieces())
	return shiftBB(targets, NORTHWEST) & bb.blackpawns
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

// NORTH clockwise through NORTHWEST a1-h8
// mapping might get annoying to remember
var RAYS = [8]int{
	NORTH,
	NORTHEAST,
	EAST,
	SOUTHEAST,
	SOUTH,
	SOUTHWEST,
	WEST,
	NORTHWEST,
}
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
func (bb bitboard) fillEast(p BB) BB {
	e := bb.emptySquares() & ^AFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, EAST))
	}
	return p
}

func (bb bitboard) fillNorthEast(p BB) BB {
	e := bb.emptySquares() & ^AFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, NORTHEAST))
	}
	return p
}

func (bb bitboard) fillSouthEast(p BB) BB {
	e := bb.emptySquares() & ^AFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, SOUTHEAST))
	}
	return p
}

func (bb bitboard) fillWest(p BB) BB {
	e := bb.emptySquares() & ^HFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, WEST))
	}
	return p
}

func (bb bitboard) fillNorthWest(p BB) BB {
	e := bb.emptySquares() & ^HFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, NORTHWEST))
	}
	return p
}

func (bb bitboard) fillSouthWest(p BB) BB {
	e := bb.emptySquares() & ^HFILE
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, SOUTHWEST))
	}
	return p
}

func (bb bitboard) fillNorth(p BB) BB {
	e := bb.emptySquares()
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, NORTH))
	}
	return p
}

func (bb bitboard) fillSouth(p BB) BB {
	e := bb.emptySquares()
	for i := 0; i < 7; i++ {
		p = p | (e & shiftBB(p, SOUTH))
	}
	return p
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
		whitequeen:   pieceSquareToBB(b, WHITE_QUEEN),
		whiteking:    pieceSquareToBB(b, WHITE_KING),
		blackpawns:   pieceSquareToBB(b, BLACK_PAWN),
		blackknights: pieceSquareToBB(b, BLACK_KNIGHT),
		blackbishops: pieceSquareToBB(b, BLACK_BISHOP),
		blackrooks:   pieceSquareToBB(b, BLACK_ROOK),
		blackqueen:   pieceSquareToBB(b, BLACK_QUEEN),
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

func (b Board) GenerateBitboardMoves() []Move {
	bb := b.newBitboard()
	moves := make([]Move, 0)
	// TODO:
	// sliding piece moves
	// ep capture & promotions
	// castling
	// absolute pins / legal moves / taboo sqs
	moves = append(moves, b.generateBitboardPawnMoves(*bb)...)
	moves = append(moves, b.generateBitboardKnightMoves(*bb)...)
	moves = append(moves, b.generateBitboardKingMoves(*bb)...)

	return moves
}

func (b Board) generateBitboardKingMoves(bb bitboard) []Move {
	// TODO: castling
	if b.Side == WHITE {
		return generateKingMovesForSide(
			bb.whiteking, bb.blackPieces(), bb.emptySquares(), WHITE_KING,
		)
	}
	return generateKingMovesForSide(
		bb.blackking, bb.whitePieces(), bb.emptySquares(), BLACK_KING,
	)
}

// TODO: update to pass attacks constant to reuse for king / knight
func generateKingMovesForSide(king, oppPieces, emptySqs BB, piece int) []Move {
	moves := make([]Move, 0)
	for king > 0 {
		lsb := deBruijnBitscan(king)
		sq := BB_TO_BOARDSQUARE[lsb]
		captures := KING_ATTACKS[lsb] & oppPieces
		quietmoves := KING_ATTACKS[lsb] & emptySqs
		for captures > 0 {
			clsb := deBruijnBitscan(captures)
			csq := BB_TO_BOARDSQUARE[clsb]
			moves = append(
				moves,
				Move{sq, csq, true, false, false, false, piece, false},
			)
			captures ^= BB(1 << clsb)
		}
		for quietmoves > 0 {
			qlsb := deBruijnBitscan(quietmoves)
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

func (b Board) generateBitboardKnightMoves(bb bitboard) []Move {
	if b.Side == WHITE {
		return generateKnightMovesForSide(
			bb.whiteknights, bb.blackPieces(), bb.emptySquares(), WHITE_KNIGHT,
		)
	}
	return generateKnightMovesForSide(
		bb.blackknights, bb.whitePieces(), bb.emptySquares(), BLACK_KNIGHT,
	)
}

func generateKnightMovesForSide(knights, oppPieces, emptySqs BB, piece int) []Move {
	moves := make([]Move, 0)
	for knights > 0 {
		lsb := deBruijnBitscan(knights)
		sq := BB_TO_BOARDSQUARE[lsb]
		captures := KNIGHT_ATTACKS[lsb] & oppPieces
		quietmoves := KNIGHT_ATTACKS[lsb] & emptySqs
		for captures > 0 {
			clsb := deBruijnBitscan(captures)
			csq := BB_TO_BOARDSQUARE[clsb]
			moves = append(
				moves,
				Move{sq, csq, true, false, false, false, piece, false},
			)
			captures ^= BB(1 << clsb)
		}
		for quietmoves > 0 {
			qlsb := deBruijnBitscan(quietmoves)
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

func (b Board) generateBitboardPawnMoves(bb bitboard) []Move {
	moves := make([]Move, 0)
	if b.Side == WHITE {
		moves = append(
			moves,
			pawnMovesFromBB(bb.pushOneWhitePawns(), -10, WHITE_PAWN, false, false)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.pushTwoWhitePawns(), -20, WHITE_PAWN, false, true)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.whitePawnsCaptureWest(), -11, WHITE_PAWN, true, false)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.whitePawnsCaptureEast(), -9, WHITE_PAWN, true, false)...,
		)
	} else {
		moves = append(
			moves,
			pawnMovesFromBB(bb.pushOneBlackPawns(), 10, BLACK_PAWN, false, false)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.pushTwoBlackPawns(), 20, BLACK_PAWN, false, true)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.blackPawnsCaptureWest(), 9, BLACK_PAWN, true, false)...,
		)
		moves = append(
			moves,
			pawnMovesFromBB(bb.blackPawnsCaptureEast(), 11, BLACK_PAWN, true, false)...,
		)
	}

	return moves
}

func pawnMovesFromBB(bb BB, delta, piece int, capture, doublePush bool) []Move {
	moves := make([]Move, 0)

	for bb > 0 {
		lsb := deBruijnBitscan(bb)
		sq := BB_TO_BOARDSQUARE[lsb]
		// TODO: promotion checking
		moves = append(moves, Move{sq, sq + delta, capture, false, false, false, piece, doublePush})
		bb ^= BB(1 << lsb)
	}

	return moves
}

var deBruijnIndex64 = [64]int{
	0, 1, 48, 2, 57, 49, 28, 3,
	61, 58, 50, 42, 38, 29, 17, 4,
	62, 55, 59, 36, 53, 51, 43, 22,
	45, 39, 33, 30, 24, 18, 12, 5,
	63, 47, 56, 27, 60, 41, 37, 16,
	54, 35, 52, 21, 44, 32, 23, 11,
	46, 26, 40, 15, 34, 20, 31, 10,
	25, 14, 19, 9, 13, 8, 7, 6,
}

const deBruijnSeq BB = 0x03f79d71b4cb0a89

func deBruijnBitscan(bb BB) int {
	return deBruijnIndex64[((bb&-bb)*deBruijnSeq)>>58]
}
