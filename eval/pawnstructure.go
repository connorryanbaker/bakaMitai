package eval

import "github.com/connorryanbaker/bakaMitai/board"

func evalPawnStructure(b board.Board) float64 {
	pcm := pawnCountMap(b)
	return float64(doubledPawnPenalties(pcm)) + float64(isolatedPawnPenalties(pcm))
}

func doubledPawnPenalties(pcm map[int]map[int]int) int {
	wdp := 0
	bdp := 0
	for i := 0; i < 8; i++ {
		if pcm[board.WHITE_PAWN][i] > 1 {
			wdp += (pcm[board.WHITE_PAWN][i] - 1)
		}
		if pcm[board.BLACK_PAWN][i] > 1 {
			bdp += (pcm[board.BLACK_PAWN][i] - 1)
		}
	}
	return bdp - wdp
}

func isolatedPawnPenalties(pcm map[int]map[int]int) int {
	wip := 0
	bip := 0
	for i := 0; i < 8; i++ {
		if i == 0 {
			if pcm[board.WHITE_PAWN][1] == 0 {
				wip += 1
			}
			if pcm[board.BLACK_PAWN][1] == 0 {
				bip += 1
			}
		} else if i == 7 {
			if pcm[board.WHITE_PAWN][6] == 0 {
				wip += 1
			}
			if pcm[board.BLACK_PAWN][6] == 0 {
				bip += 1
			}
		} else {
			if pcm[board.WHITE_PAWN][i-1] == 0 && pcm[board.WHITE_PAWN][i+1] == 0 {
				wip += 1
			}
			if pcm[board.BLACK_PAWN][i-1] == 0 && pcm[board.BLACK_PAWN][i+1] == 0 {
				bip += 1
			}
		}
	}
	return bip - wip
}

func pawnCountMap(b board.Board) map[int]map[int]int {
	var m = map[int]map[int]int{
		board.WHITE_PAWN: {
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
			6: 0,
			7: 0,
		},
		board.BLACK_PAWN: {
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
			6: 0,
			7: 0,
		},
	}
	for i := 0; i < 8; i++ {
		r := board.MAILBOX_64[i]
		for ; r <= board.IH1; r += 10 {
			if b.PieceAt(r) == board.WHITE_PAWN {
				m[board.WHITE_PAWN][i] += 1
			}
			if b.PieceAt(r) == board.BLACK_PAWN {
				m[board.BLACK_PAWN][i] += 1
			}
		}
	}
	return m
}
