package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/connorryanbaker/bakaMitai/board"
	"github.com/connorryanbaker/bakaMitai/engine"
	"github.com/connorryanbaker/bakaMitai/search"
	"github.com/connorryanbaker/bakaMitai/uci"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var depth = flag.Int("depth", 3, "perft movegeneration depth")
var selfPlay = flag.Bool("selfPlay", false, "play against self, print board to console")

func main() {
	flag.Parse()

	if *selfPlay {
		b := board.NewBoard()
		play(b)
		return
	}
	uci.Run()
}

func play(b board.Board) {
	e := engine.New()
	for {
		b.Print()
		m := e.GenMove(&b, 5000)
		m.Print()
		b.MakeMove(m)
		if b.Checkmate() {
			b.Print()
			fmt.Println("checkmate!")
			fmt.Println(board.ToPGN(b.History))
			return
		} else if b.Stalemate() {
			b.Print()
			fmt.Println("stalemate!")
			fmt.Println(board.ToPGN(b.History))
			return
		} else if b.FiftyMoveDraw() {
			b.Print()
			fmt.Println("50 move draw!")
			fmt.Println(board.ToPGN(b.History))
			return
		} else if b.ThreefoldRepetition() {
			b.Print()
			fmt.Println("3fold!")
			fmt.Println(board.ToPGN(b.History))
			return
		} else if b.InsufficientMaterial() {
			b.Print()
			fmt.Println("insufficient material!")
			fmt.Println(board.ToPGN(b.History))
			return
		}
	}
}

func profileBBPerft(b board.Board) {
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal("couldn't create CPU profile file: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("couldn't start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
	board.BBperft(&b, *depth)
}

func profileSearch(b board.Board) {
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal("couldn't create CPU profile file: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("couldn't start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
	pv := search.NewLine(*depth)
	search.Search(&b, *depth, &pv, time.Now())
}
