// https://adventofcode.com/2025/day/7

package day7

import (
	"fmt"
	"strings"
)

const SPLITTER_SYMBOL = "^"
const START_SYMBOL = "S"

type EdgeMap struct {
	Width  int
	Height int
	Map    [][]int
}

func NewEdgeMap(width int, height int) *EdgeMap {
	edgeMap := EdgeMap{width, height, make([][]int, height)}

	for y := range height {
		edgeMap.Map[y] = make([]int, width)
	}

	return &edgeMap
}

func (edgeMap EdgeMap) Get(x int, y int) int {
	if x >= 0 && x < edgeMap.Width && y >= 0 && y < edgeMap.Height {
		return edgeMap.Map[y][x]
	}

	return 0
}

func (edgeMap EdgeMap) Set(x int, y int, paths int) {
	if x >= 0 && x < edgeMap.Width && y >= 0 && y < edgeMap.Height {
		edgeMap.Map[y][x] = paths
	}
}

func Solve(input string) {
	rows := strings.Split(input, "\n")
	width := len(rows[0])
	height := len(rows)

	edgeMap := *NewEdgeMap(width, height)

	for y := range height {
		for x := range width {
			propagatedPaths := edgeMap.Get(x, y)

			switch string(rows[y][x]) {
			case START_SYMBOL:
				edgeMap.Set(x, y+1, 1)
			case SPLITTER_SYMBOL:
				edgeMap.Set(x, y+1, 0)

				if edgeMap.Get(x, y) > 0 {
					edgeMap.Set(x-1, y+1, edgeMap.Get(x-1, y+1)+propagatedPaths)
					edgeMap.Set(x+1, y+1, edgeMap.Get(x+1, y+1)+propagatedPaths)
				}
			default:
				edgeMap.Set(x, y+1, edgeMap.Get(x, y+1)+propagatedPaths)
			}
		}
	}

	totalPaths := 0
	for x := range width {
		totalPaths += edgeMap.Get(x, height-1)
	}

	fmt.Printf("%d\n", totalPaths)
}
