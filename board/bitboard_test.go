package board

import "testing"

func TestEmptySquares(t *testing.T) {
  var expected BB = 0b1111111100000000111111111111111111111111111111110000000011111111
  bb := NewBitboard()
  if bb.emptySquares() != expected {
    t.Errorf("Unexpected empty squares; expected %b, received %b", expected, bb.emptySquares())
  }
}

func TestAllPieces(t *testing.T) {
  var expected BB = 0b0000000011111111000000000000000000000000000000001111111100000000
  bb := NewBitboard()
  if bb.allPieces() != expected {
    t.Errorf("Unexpected allPieces; expected %b, received %b", expected, bb.allPieces())
  }
}

func TestPushOneWhitePawns(t *testing.T) {
  var tests = []struct{
    bb bitboard
    e  BB
  }{
    {
      NewBitboard(),
      0x000000000000FF00,
    },
    {
      bitboard{
        INIT_WHITE_PAWN_BB,
        0x0000000000FF0000,
      },
      0x0000000000000000,
    },
  }

  for _, tt := range tests {
    res := tt.bb.pushOneWhitePawns()
    if res != tt.e {
      t.Errorf("Unexpected pushOneWhitePawns; expected %b, received %b", tt.e, res)
    }
  }
}

func TestPushOneBlackPawns(t *testing.T) {
  var tests = []struct{
    bb bitboard
    e  BB
  }{
    {
      NewBitboard(),
      0x00FF000000000000,
    },
    {
      bitboard{
        0x0000FF0000000000,
        INIT_BLACK_PAWN_BB,
      },
      0x0000000000000000,
    },
  }

  for _, tt := range tests {
    res := tt.bb.pushOneBlackPawns()
    if res != tt.e {
      t.Errorf("Unexpected pushOneBlackPawns; expected %b, received %b", tt.e, res)
    }
  }
}

func TestPushTwoWhitePawns(t *testing.T) {
  var tests = []struct{
    bb bitboard
    e  BB
  }{
    {
      NewBitboard(),
      0x000000000000FF00,
    },
    {
      bitboard{
        INIT_WHITE_PAWN_BB,
        0x00000000FF000000,
      },
      0x0000000000000000,
    },
    {
      bitboard{
        INIT_WHITE_PAWN_BB,
        0x0000000000FF0000,
      },
      0x0000000000000000,
    },
    {
      bitboard{
        INIT_WHITE_PAWN_BB,
        0b0000000000000000000000000000000010000000000000000000000000000000,
      },
      0b0000000000000000000000000000000000000000000000000111111100000000,
    },
    {
      bitboard{
        INIT_WHITE_PAWN_BB,
        0b0000000000000000000000000000000010101010000000000000000000000000,
      },
      0b0000000000000000000000000000000000000000000000000101010100000000,
    },
    {
      bitboard{
        INIT_WHITE_PAWN_BB,
        0b0000000000000000000000000000000010101010010101010000000000000000,
      },
      0x0000000000000000,
    },
    {
      bitboard{
        INIT_WHITE_PAWN_BB,
        0b0000000000000000000000000000000000000000010101010000000000000000,
      },
      0b0000000000000000000000000000000000000000000000001010101000000000,
    },
  }

  for _, tt := range tests {
    res := tt.bb.pushTwoWhitePawns()
    if res != tt.e {
      t.Errorf("Unexpected pushTwoWhitePawns; expected %b, received %b", tt.e, res)
    }
  }
}


func TestPushTwoBlackPawns(t *testing.T) {
  var tests = []struct{
    bb bitboard
    e  BB
  }{
    {
      NewBitboard(),
      0x00FF000000000000,
    },
    {
      bitboard{
        0x000000FF00000000,
        INIT_BLACK_PAWN_BB,
      },
      0x0000000000000000,
    },
    {
      bitboard{
        0x0000FF0000000000,
        INIT_BLACK_PAWN_BB,
      },
      0x0000000000000000,
    },
    {
      bitboard{
        0b0000000000000000000000001000000000000000000000000000000000000000,
        INIT_BLACK_PAWN_BB,
      },
      0b0000000001111111000000000000000000000000000000000000000000000000,
    },
    {
      bitboard{
        0b0000000000000000000000001010101000000000000000000000000000000000,
        INIT_BLACK_PAWN_BB,
      },
      0b0000000001010101000000000000000000000000000000000000000000000000,
    },
    {
      bitboard{
        0b0000000000000000010101011010101000000000000000000000000000000000,
        INIT_BLACK_PAWN_BB,
      },
      0x0000000000000000,
    },
    {
      bitboard{
        0b0000000000000000101010100000000000000000000000000000000000000000,
        INIT_BLACK_PAWN_BB,
      },
      0b0000000001010101000000000000000000000000000000000000000000000000,
    },
  }

  for _, tt := range tests {
    res := tt.bb.pushTwoBlackPawns()
    if res != tt.e {
      t.Errorf("Unexpected pushTwoBlackPawns; expected %b, received %b", tt.e, res)
    }
  }
}

func TestWhitePawnAttacks(t *testing.T) {
  var tests = []struct{
    bb bitboard
    e  BB
  }{
    {
      NewBitboard(),
      0x0000000000FF0000,
    },
    {
      bitboard{
        0b0000000000000000000000000000000000000000000000001000000100000000,
        INIT_BLACK_PAWN_BB,
      },
      0b0000000000000000000000000000000000000000010000100000000000000000,
    },
  }

  for _, tt := range tests {
    res := tt.bb.whitePawnAttacks()
    if tt.e != res {
      printBB(tt.bb.whitePawnAttacks())
      t.Errorf("Unexpected white pawn attacks; expected %b, received %b", tt.e, res)
    }
  }
}

func TestBlackPawnAttacks(t *testing.T) {
  var tests = []struct{
    bb bitboard
    e  BB
  }{
    {
      NewBitboard(),
      0x0000FF0000000000,
    },
    {
      bitboard{
        INIT_WHITE_PAWN_BB,
        0b0000000010000001000000000000000000000000000000000000000000000000,
      },
      0b0000000000000000010000100000000000000000000000000000000000000000,
    },
  }

  for _, tt := range tests {
    res := tt.bb.blackPawnAttacks()
    if tt.e != res {
      printBB(tt.bb.blackPawnAttacks())
      printBB(tt.e)
      t.Errorf("Unexpected black pawn attacks; expected %b, received %b", tt.e, res)
    }
  }
}


