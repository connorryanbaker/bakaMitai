package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/connorryanbaker/engine/board"
	"github.com/connorryanbaker/engine/search"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var depth = flag.Int("depth", 3, "engine halfply search depth")

func main() {
	flag.Parse()

	b := board.NewBoard()
	if *cpuprofile != "" {
		profileSearch(b)
		return
	}
	play(b)
}

func sum(n []int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}

func play(b board.Board) {
	nodelist := make([]int, 0)
	for true {
		var nodes int
		b.Print()
		moves := search.SearchNodeCount(&b, *depth, &nodes)
		fmt.Println(nodes)
		fmt.Println(moves)
		b.MakeMove(moves[0])
		nodelist = append(nodelist, nodes)
		if b.Checkmate() {
			b.Print()
			fmt.Println("checkmate!")
			printHistory(b)
			fmt.Println(sum(nodelist))
			return
		} else if b.Stalemate() {
			b.Print()
			fmt.Println("stalemate!")
			printHistory(b)
			fmt.Println(sum(nodelist))
			fmt.Println(nodelist)
			return
		} else if b.FiftyMoveDraw() {
			b.Print()
			fmt.Println("50 move draw!")
			printHistory(b)
			fmt.Println(sum(nodelist))
			fmt.Println(nodelist)
			return
		} else if b.ThreefoldRepetition() {
			b.Print()
			fmt.Println("3fold!")
			printHistory(b)
			fmt.Println(sum(nodelist))
			fmt.Println(nodelist)
			return
		} else if b.InsufficientMaterial() {
			b.Print()
			fmt.Println("insufficient material!")
			printHistory(b)
			fmt.Println(sum(nodelist))
			fmt.Println(nodelist)
			return
		}
	}
}

func printHistory(b board.Board) {
	h := b.History
	for i := 0; i < len(h); i++ {
		fmt.Printf("%d: FROM: %s TO: %s\n", i+1, board.SQ_NUM_TO_NAME[h[i].Move.From], board.SQ_NUM_TO_NAME[h[i].Move.To])
	}
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
	search.Search(&b, *depth)
}
