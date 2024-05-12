package uci

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/connorryanbaker/bakaMitai/board"
	"github.com/connorryanbaker/bakaMitai/engine"
)

// commands
// uci
// isready
// position
// go
// stop
// quit

type uci struct {
	board  *board.Board
	engine engine.Engine
	output chan board.Move
}

func getInput(input chan<- string) {
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		if cmd != "" {
			input <- cmd
		}
	}
}

func Run() {
	var input = make(chan string)
	u := uci{}
	u.output = make(chan board.Move)
	defer close(u.output)

	go func() {
		defer close(input)
		getInput(input)
	}()

	for {
		select {
		case move, ok := <-u.output:
			if !ok {
				return
			}
			u.bestMove(move)
		case command, ok := <-input:
			if !ok || command == "quit" {
				return
			}
			u.handleInput(command)
		}
	}
}

func (u *uci) bestMove(m board.Move) {
	var promote string
	if m.Promote {
		promote = promoTranslate[u.board.Side][m.PromotionPiece]
	}
	from := board.SQ_NUM_TO_NAME[m.From]
	to := board.SQ_NUM_TO_NAME[m.To]

	fmt.Printf("bestmove %s%s%s\n", from, to, promote)
}

func (u *uci) handleInput(input string) {
	commands := strings.Fields(input)
	command := commands[0]
	switch command {
	case "uci":
		fmt.Println("id name BakaMitai")
		fmt.Println("id author connorryanbaker")
		fmt.Println("uciok")
	case "isready":
		fmt.Println("readyok")
	case "position":
		u.parsePosition(commands[1:])
	case "go":
		timeRemaining := u.parseTimeRemaining(commands)
		go func() {
			m := u.engine.GenMove(u.board, timeRemaining)
			u.output <- m
		}()
	case "stop":
		// stop and return best move
		u.bestMove(u.engine.PV.Moves[0])
	}
}

func (u *uci) parseTimeRemaining(commands []string) int {
	side := u.board.Side
	var idx int
	if side == board.WHITE {
		idx = indexOf("wtime", commands)
	} else {
		idx = indexOf("btime", commands)
	}
	ms, err := strconv.Atoi(commands[idx+1])
	if err != nil {
		return 1000
	}
	return ms
}

func indexOf(s string, strs []string) int {
	for i, e := range strs {
		if e == s {
			return i
		}
	}
	// shouldn't happen
	return -1
}

func (u *uci) parsePosition(fenAndMoves []string) {
	fen := fenAndMoves[0]
	moves := make([]string, 0)
	if len(fenAndMoves) > 2 {
		moves = fenAndMoves[2:]
	}

	u.board = parseFen(fen)
	u.playMoves(moves)
}

func (u *uci) playMoves(moves []string) {
	for _, m := range moves {
		legals := u.board.GenerateBitboardMoves()
		var toPlay board.Move
		if len(m) == 5 {
			toPlay = findPromotionMove(m, legals, u.board.Side)
		} else {
			toPlay = findMove(m, legals)
		}
		u.board.MakeBBMove(toPlay)
	}
}

var promoLookup = map[int]map[string]int{
	board.WHITE: {
		"q": board.WHITE_QUEEN,
		"r": board.WHITE_ROOK,
		"b": board.WHITE_BISHOP,
		"n": board.WHITE_KNIGHT,
	},
	board.BLACK: {
		"q": board.BLACK_QUEEN,
		"r": board.BLACK_ROOK,
		"b": board.BLACK_BISHOP,
		"n": board.BLACK_KNIGHT,
	},
}

var promoTranslate = map[int]map[int]string{
	board.WHITE: {
		board.WHITE_QUEEN:  "q",
		board.WHITE_ROOK:   "r",
		board.WHITE_BISHOP: "b",
		board.WHITE_KNIGHT: "n",
	},
	board.BLACK: {
		board.BLACK_QUEEN:  "q",
		board.BLACK_ROOK:   "r",
		board.BLACK_BISHOP: "b",
		board.BLACK_KNIGHT: "n",
	},
}

func findPromotionMove(move string, moves []board.Move, side int) board.Move {
	for _, m := range moves {
		if board.SQ_NUM_TO_NAME[m.From] == move[0:2] && board.SQ_NUM_TO_NAME[m.To] == move[2:5] {
			if m.PromotionPiece == promoLookup[side][move[5:]] {
				return m
			}
		}
	}
	// shouldn't happen
	return board.Move{}
}

func findMove(move string, moves []board.Move) board.Move {
	for _, m := range moves {
		if board.SQ_NUM_TO_NAME[m.From] == move[0:2] && board.SQ_NUM_TO_NAME[m.To] == move[2:] {
			return m
		}
	}
	// shouldn't happen
	return board.Move{}
}

func parseFen(fen string) *board.Board {
	b := board.NewBoard()
	if fen == "startpos" {
		return &b
	}

	b = board.FromFENString(fen)
	return &b
}
