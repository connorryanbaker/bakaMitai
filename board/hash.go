package board

import "math/rand"

type hash struct {
	pieceSquares [64][12]uint64 // random number per square per piece + ep squares
	epFile       [8]uint64
	castle       [4]uint64
	blackToMove  uint64
}

func (b Board) Hash() uint64 {
	var h uint64
	for piece, squares := range b.PieceSquares {
		for _, sq := range squares {
			h ^= b.hashSeed.pieceSquares[SQ_NAME_TO_SQ_64[sq]][piece-1]
		}
	}
	for i, v := range b.Castle {
		if v {
			h ^= b.hashSeed.castle[i]
		}
	}
	if b.Ep != nil {
		h ^= b.hashSeed.epFile[epSquareFile(b.Ep)]
	}
	if b.Side == BLACK {
		h ^= b.hashSeed.blackToMove
	}
	return h
}

func newHashSeed() hash {
	h := hash{}
	for i := 0; i < 63; i++ {
		for j := WHITE_PAWN; j <= BLACK_KING; j++ {
			h.pieceSquares[i][j-1] = rand.Uint64()
		}
	}
	for i := 0; i < 8; i++ {
		h.epFile[i] = rand.Uint64()
	}
	for i := 0; i < 4; i++ {
		h.castle[i] = rand.Uint64()
	}
	h.blackToMove = rand.Uint64()
	return h
}
