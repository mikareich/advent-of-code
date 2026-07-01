package day6

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ADDITION       = "+"
	MULTIPLICATION = "*"
	SPACE          = " "
)

func Solve(input string) {
	// 	totalSum = 0
	// 	for col_1...col_m of input
	// 		intermediate_sum = 0
	// 		for row_n ... row_1 of col_i
	// 			row_(i=n): determine op
	// 			row_(i!=n): intermediate_sum [op] row_i
	// 		totalSum += intermediate_sum

	rows := strings.Split(input, "\n")
	topRow := rows[0]
	inputWidth := len(topRow)
	inputHeight := len(rows)

	totalSum := 0

	var readAt = func(x int, y int) (string, bool) {
		if x < 0 || x >= inputWidth || y < 0 || y >= inputHeight {
			return "", false
		}

		return string(rows[y][x]), true
	}

	colOffset := 0
	for n := 0; ; n += 1 {
		// determine col width
		colWidth := 0
		foundFirstChar := false
		foundLastChar := false
		for ; !foundLastChar; colWidth += 1 {
			offset := colOffset + colWidth
			char, _ := readAt(offset, inputHeight-1)
			nextChar, hasNextChar := readAt(offset+1, inputHeight-1)

			foundFirstChar = foundFirstChar || (char != SPACE && nextChar == SPACE)
			foundLastChar = !hasNextChar || foundFirstChar && nextChar != SPACE
		}

		// read OP
		op, _ := readAt(colOffset, inputHeight-1)
		var predicate func(a int, b int) int

		var intermediateSum int
		switch op {
		case ADDITION:
			intermediateSum = 0
			predicate = func(a int, b int) int { return a + b }
		case MULTIPLICATION:
			intermediateSum = 1
			predicate = func(a int, b int) int { return a * b }
		default:
			panic("Invalid op found!")
		}

		// compute intermediate sum
		for offset := colOffset; offset < colOffset+colWidth; offset += 1 {
			var builder strings.Builder

			for m := range inputHeight - 1 {
				char, _ := readAt(offset, m)
				if char != SPACE {
					builder.WriteString(char)
				}
			}

			value, err := strconv.Atoi(builder.String())
			if err == nil {
				intermediateSum = predicate(intermediateSum, value)
			}

			fmt.Printf("%v \n", value)

		}

		// compute total sum
		totalSum += intermediateSum

		colOffset += colWidth
		if colOffset == inputWidth {
			break
		}
	}

	fmt.Printf("%d \n", totalSum)
}
