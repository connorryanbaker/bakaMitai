package eval

import (
  "github.com/connorryanbaker/engine/board"
)

// definitions from chess programming simplified evaluation function
var WHITE_PAWN_SQ_VALUES = [64]float64{
 0,  0,  0,  0,  0,  0,  0,  0,
50, 50, 50, 50, 50, 50, 50, 50,
10, 10, 20, 30, 30, 20, 10, 10,
 5,  5, 10, 25, 25, 10,  5,  5,
 0,  0,  0, 20, 20,  0,  0,  0,
 5, -5,-10,  0,  0,-10, -5,  5,
 5, 10, 10,-20,-20, 10, 10,  5,
 0,  0,  0,  0,  0,  0,  0,  0,
}

var BLACK_PAWN_SQ_VALUES = [64]float64{
 0,  0,  0,  0,  0,  0,  0,  0,
 5, 10, 10,-20,-20, 10, 10,  5,
 5, -5,-10,  0,  0,-10, -5,  5,
 0,  0,  0, 20, 20,  0,  0,  0,
 5,  5, 10, 25, 25, 10,  5,  5,
10, 10, 20, 30, 30, 20, 10, 10,
50, 50, 50, 50, 50, 50, 50, 50,
 0,  0,  0,  0,  0,  0,  0,  0,
}

var KNIGHT_SQ_VALUES = [64]float64{
-50,-40,-30,-30,-30,-30,-40,-50,
-40,-20,  0,  0,  0,  0,-20,-40,
-30,  0, 10, 15, 15, 10,  0,-30,
-30,  5, 15, 20, 20, 15,  5,-30,
-30,  0, 15, 20, 20, 15,  0,-30,
-30,  5, 10, 15, 15, 10,  5,-30,
-40,-20,  0,  5,  5,  0,-20,-40,
-50,-40,-30,-30,-30,-30,-40,-50,
}

var BISHOP_SQ_VALUES = [64]float64{
-20,-10,-10,-10,-10,-10,-10,-20,
-10,  0,  0,  0,  0,  0,  0,-10,
-10,  0,  5, 10, 10,  5,  0,-10,
-10,  5,  5, 10, 10,  5,  5,-10,
-10,  0, 10, 10, 10, 10,  0,-10,
-10, 10, 10, 10, 10, 10, 10,-10,
-10,  5,  0,  0,  0,  0,  5,-10,
-20,-10,-10,-10,-10,-10,-10,-20,
}

var WHITE_ROOK_SQ_VALUES = [64]float64{
 0,  0,  0,  0,  0,  0,  0,  0,
 5, 10, 10, 10, 10, 10, 10,  5,
-5,  0,  0,  0,  0,  0,  0, -5,
-5,  0,  0,  0,  0,  0,  0, -5,
-5,  0,  0,  0,  0,  0,  0, -5,
-5,  0,  0,  0,  0,  0,  0, -5,
-5,  0,  0,  0,  0,  0,  0, -5,
 0,  0,  0,  5,  5,  3,  0,  0,
}

var BLACK_ROOK_SQ_VALUES = [64]float64{
 0,  0,  0,  5,  5,  3,  0,  0,
-5,  0,  0,  0,  0,  0,  0, -5,
-5,  0,  0,  0,  0,  0,  0, -5,
-5,  0,  0,  0,  0,  0,  0, -5,
-5,  0,  0,  0,  0,  0,  0, -5,
-5,  0,  0,  0,  0,  0,  0, -5,
 5, 10, 10, 10, 10, 10, 10,  5,
 0,  0,  0,  0,  0,  0,  0,  0,
}

var WHITE_KING_MIDDLE_GAME_SQ_VALUES = [64]float64{
-30,-40,-40,-50,-50,-40,-40,-30,
-30,-40,-40,-50,-50,-40,-40,-30,
-30,-40,-40,-50,-50,-40,-40,-30,
-30,-40,-40,-50,-50,-40,-40,-30,
-20,-30,-30,-40,-40,-30,-30,-20,
-10,-20,-20,-20,-20,-20,-20,-10,
 20, 20,  0,  0,  0,  0, 20, 20,
 20, 30, 10,  0,  0, 10, 30, 20,
}
var BLACK_KING_MIDDLE_GAME_SQ_VALUES = [64]float64{
 20, 30, 10,  0,  0, 10, 30, 20,
 20, 20,  0,  0,  0,  0, 20, 20,
-10,-20,-20,-20,-20,-20,-20,-10,
-20,-30,-30,-40,-40,-30,-30,-20,
-30,-40,-40,-50,-50,-40,-40,-30,
-30,-40,-40,-50,-50,-40,-40,-30,
-30,-40,-40,-50,-50,-40,-40,-30,
-30,-40,-40,-50,-50,-40,-40,-30,
}

var KING_END_GAME_SQ_VALUES = [64]float64{
-50,-40,-30,-20,-20,-30,-40,-50,
-30,-20,-10,  0,  0,-10,-20,-30,
-30,-10, 20, 30, 30, 20,-10,-30,
-30,-10, 30, 40, 40, 30,-10,-30,
-30,-10, 30, 40, 40, 30,-10,-30,
-30,-10, 20, 30, 30, 20,-10,-30,
-30,-30,  0,  0,  0,  0,-30,-30,
-50,-30,-30,-30,-30,-30,-30,-50,
}

var PIECE_TO_LOOKUP = map[int][64]float64{
  board.WHITE_PAWN: WHITE_PAWN_SQ_VALUES,
  board.BLACK_PAWN: WHITE_PAWN_SQ_VALUES,
  board.WHITE_KNIGHT: KNIGHT_SQ_VALUES,
  board.BLACK_KNIGHT: KNIGHT_SQ_VALUES,
  board.WHITE_BISHOP: BISHOP_SQ_VALUES,
  board.BLACK_BISHOP: BISHOP_SQ_VALUES,
  board.WHITE_ROOK: WHITE_ROOK_SQ_VALUES,
  board.BLACK_ROOK: WHITE_ROOK_SQ_VALUES,
  // board.WHITE_QUEEN: QUEEN_SQ_VALUES,
  // board.BLACK_QUEEN: QUEEN_SQ_VALUES, // TODO: flip / edit
  board.WHITE_KING: WHITE_KING_MIDDLE_GAME_SQ_VALUES,
  board.BLACK_KING: BLACK_KING_MIDDLE_GAME_SQ_VALUES,
}

// var PIECE_WEIGHTS = map[int]float64{
//   board.WHITE_PAWN: 1,
//   board.BLACK_PAWN: 1,
//   board.WHITE_KNIGHT: 1000,
//   board.BLACK_KNIGHT: 1000,
//   board.WHITE_BISHOP: 1000,
//   board.BLACK_BISHOP: 1000,
//   board.WHITE_ROOK: 1000,
//   board.BLACK_ROOK: 1000,
//   board.WHITE_QUEEN: 1,
//   board.BLACK_QUEEN: 1,
//   board.WHITE_KING: 200,
//   board.BLACK_KING: 200,
// }

func evalPieceSquares(b board.Board) float64 {
  var e float64
  e += whitePawnSqValues(b)
  e += blackPawnSqValues(b) * -1
  e += whiteKnightSqValues(b)
  e += blackKnightSqValues(b) * -1
  e += whiteBishopSqValues(b)
  e += blackBishopSqValues(b) * -1
  e += whiteKingSqValues(b)
  e += blackKingSqValues(b) * -1

  // for p, sqs := range b.Pieces() {
  //   lookup := PIECE_TO_LOOKUP[p]
  //   // TODO: king endgame
  //   for sq := range sqs {
  //     v := lookup[board.SQ_NAME_TO_SQ_64[sq]]
  //     if p < board.WHITE_KING {
  //       v *= -1
  //     }
  //     e += v
  //   }
  // }
  return e
}

func blackKingSqValues(b board.Board) float64 {
  var e float64
  sqs := b.Pieces()[board.BLACK_KING]

  for _, s := range sqs {
    v := BLACK_KING_MIDDLE_GAME_SQ_VALUES[board.SQ_NAME_TO_SQ_64[s]]
    e += v
  }

  return e
}


func whiteKingSqValues(b board.Board) float64 {
  var e float64
  sqs := b.Pieces()[board.WHITE_KING]

  for _, s := range sqs {
    v := WHITE_KING_MIDDLE_GAME_SQ_VALUES[board.SQ_NAME_TO_SQ_64[s]]
    e += v
  }

  return e
}


func blackRookSqValues(b board.Board) float64 {
  var e float64
  sqs := b.Pieces()[board.BLACK_ROOK]
  if sqs == nil {
    return e
  }

  for _, s := range sqs {
    e += BLACK_ROOK_SQ_VALUES[board.SQ_NAME_TO_SQ_64[s]]
  }

  return e
}

func whiteRookSqValues(b board.Board) float64 {
  var e float64
  sqs := b.Pieces()[board.WHITE_ROOK]
  if sqs == nil {
    return e
  }

  for _, s := range sqs {
    e += WHITE_ROOK_SQ_VALUES[board.SQ_NAME_TO_SQ_64[s]]
  }

  return e
}

func blackBishopSqValues(b board.Board) float64 {
  var e float64
  sqs := b.Pieces()[board.BLACK_BISHOP]
  if sqs == nil {
    return e
  }

  for _, s := range sqs {
    e += BISHOP_SQ_VALUES[board.SQ_NAME_TO_SQ_64[s]]
  }

  return e
}

func whiteBishopSqValues(b board.Board) float64 {
  var e float64
  sqs := b.Pieces()[board.WHITE_BISHOP]
  if sqs == nil {
    return e
  }

  for _, s := range sqs {
    e += BISHOP_SQ_VALUES[board.SQ_NAME_TO_SQ_64[s]]
  }

  return e
}

func whiteKnightSqValues(b board.Board) float64 {
  var e float64
  sqs := b.Pieces()[board.WHITE_KNIGHT]
  if sqs == nil {
    return e
  }

  for _, s := range sqs {
    e += KNIGHT_SQ_VALUES[board.SQ_NAME_TO_SQ_64[s]]
  }

  return e
}

func blackKnightSqValues(b board.Board) float64 {
  var e float64
  sqs := b.Pieces()[board.BLACK_KNIGHT]
  if sqs == nil {
    return e
  }

  for _, s := range sqs {
    e += KNIGHT_SQ_VALUES[board.SQ_NAME_TO_SQ_64[s]]
  }

  return e
}

func whitePawnSqValues(b board.Board) float64 {
  var e float64
  sqs := b.Pieces()[board.WHITE_PAWN]
  if sqs == nil {
    return e
  }

  for _, s := range sqs {
    v := WHITE_PAWN_SQ_VALUES[board.SQ_NAME_TO_SQ_64[s]]
    e += v
  }

  return e
}

func blackPawnSqValues(b board.Board) float64 {
  var e float64
  sqs := b.Pieces()[board.BLACK_PAWN]
  if sqs == nil {
    return e
  }

  for _, s := range sqs {
    v := BLACK_PAWN_SQ_VALUES[board.SQ_NAME_TO_SQ_64[s]]
    e += v
  }

  return e
}
