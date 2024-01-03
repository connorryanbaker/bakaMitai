package board

import "testing"

type expectation struct {
  sq    int
  piece int
}

func TestFromFENString(t *testing.T) {
  var tests = []struct{
    fen string
    expectations []expectation
    side int
    castle [4]bool
    ep *int
    hply int
    ply int
  }{
    {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      []expectation{
        {A8, BLACK_ROOK},
        {B8, BLACK_KNIGHT},
        {C8, BLACK_BISHOP},
        {D8, BLACK_QUEEN},
        {E8, BLACK_KING},
        {F8, BLACK_BISHOP},
        {G8, BLACK_KNIGHT},
        {H8, BLACK_ROOK},
        {A7, BLACK_PAWN},
        {B7, BLACK_PAWN},
        {C7, BLACK_PAWN},
        {D7, BLACK_PAWN},
        {E7, BLACK_PAWN},
        {F7, BLACK_PAWN},
        {G7, BLACK_PAWN},
        {H7, BLACK_PAWN},
        {A6, EMPTY_SQUARE},
        {B6, EMPTY_SQUARE},
        {C6, EMPTY_SQUARE},
        {D6, EMPTY_SQUARE},
        {E6, EMPTY_SQUARE},
        {F6, EMPTY_SQUARE},
        {G6, EMPTY_SQUARE},
        {H6, EMPTY_SQUARE},
        {A5, EMPTY_SQUARE},
        {B5, EMPTY_SQUARE},
        {C5, EMPTY_SQUARE},
        {D5, EMPTY_SQUARE},
        {E5, EMPTY_SQUARE},
        {F5, EMPTY_SQUARE},
        {G5, EMPTY_SQUARE},
        {H5, EMPTY_SQUARE},
        {A4, EMPTY_SQUARE},
        {B4, EMPTY_SQUARE},
        {C4, EMPTY_SQUARE},
        {D4, EMPTY_SQUARE},
        {E4, EMPTY_SQUARE},
        {F4, EMPTY_SQUARE},
        {G4, EMPTY_SQUARE},
        {H4, EMPTY_SQUARE},
        {A3, EMPTY_SQUARE},
        {B3, EMPTY_SQUARE},
        {C3, EMPTY_SQUARE},
        {D3, EMPTY_SQUARE},
        {E3, EMPTY_SQUARE},
        {F3, EMPTY_SQUARE},
        {G3, EMPTY_SQUARE},
        {H3, EMPTY_SQUARE},
        {A2, WHITE_PAWN},
        {B2, WHITE_PAWN},
        {C2, WHITE_PAWN},
        {D2, WHITE_PAWN},
        {E2, WHITE_PAWN},
        {F2, WHITE_PAWN},
        {G2, WHITE_PAWN},
        {H2, WHITE_PAWN},
        {A1, WHITE_ROOK},
        {B1, WHITE_KNIGHT},
        {C1, WHITE_BISHOP},
        {D1, WHITE_QUEEN},
        {E1, WHITE_KING},
        {F1, WHITE_BISHOP},
        {G1, WHITE_KNIGHT},
        {H1, WHITE_ROOK},
      },
      0,
      [4]bool{true,true,true,true},
      nil,
      0,
      0,
    },
  }

  // TODO: more tests obviously

  for _, tt := range tests {
    b := FromFENString(tt.fen)
    for _, expectation := range tt.expectations {
      received := b.PieceAt(expectation.sq)
      if received != expectation.piece {
        t.Errorf("sq: %d, received %d, expected %d", expectation.sq, received, expectation.piece)
      }
    }
    if b.side != tt.side {
        t.Errorf("side - received %d, expected %d", b.side, tt.side)
    }
    if b.castle != tt.castle {
      t.Errorf("castle - received %v, expected %v", b.castle, tt.castle)
    }
    if b.ep != tt.ep {
      t.Errorf("ep - received %p, expected %p", b.ep, tt.ep)
    }
    if b.hply != tt.hply {
      t.Errorf("hply - received %d, expected %d", b.hply, tt.hply)
    }
    if b.ply != tt.ply {
      t.Errorf("ply - received %d, expected %d", b.ply, tt.ply)
    }
  }
}
