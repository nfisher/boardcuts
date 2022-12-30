package main

import (
	"fmt"
	"sort"
)

func main() {
	var cuts = []float64{
		15.5,
		15.5,
		15.5,
		15.5,
		15.5,

		30.75,
		30.75,
		30.75,
		30.75,

		17.0,
		17.0,
		17.0,
		17.0,
		17.0,
		17.0,
		17.0,
		17.0,

		24.0,
		24.0,
		24.0,
		24.0,
	}

	boardLength := 96.0
	kerf := 0.125
	boards := packAvailable(cuts, boardLength, kerf)
	packPrint(boards, boardLength, kerf)
}

func packPrint(boards []*Board, boardLength, kerf float64) {
	for i, b := range boards {
		fmt.Printf("Board %v, waste %v\"\n", i+1, boardLength-b.CutLength)
		for _, c := range b.Cuts {
			if c != kerf {
				fmt.Println(" ", c)
			}
		}
	}
}

func packAvailable(cuts []float64, boardLength, kerf float64) []*Board {
	sort.Sort(sort.Reverse(sort.Float64Slice(cuts)))

	var boards []*Board
	for _, c := range cuts {
		isPacked := false
		for _, b := range boards {
			if boardLength-b.CutLength >= c+kerf {
				isPacked = true
				b.Cuts = append(b.Cuts, c, kerf)
				b.CutLength += c + kerf
				break
			}
		}
		if !isPacked {
			var b = Board{}
			b.Cuts = append(b.Cuts, c, kerf)
			b.CutLength += c + kerf
			boards = append(boards, &b)
		}
	}

	return boards
}

type Board struct {
	Cuts      []float64
	CutLength float64
}
