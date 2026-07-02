package main

import (
	"os"
	"strconv"

	"github.com/mikareich/advent-of-code/g/day6"
	"github.com/mikareich/advent-of-code/g/day7"
)

const DAY_ARG = "--day"
const INPUT_ARG = "--input"

func parseArgs() (int, string) {
	args := os.Args[1:]

	var day int
	var inputPath string
	var err error

	for i := range len(args) {
		if i == 0 {
			continue
		}

		arg := args[i]

		if args[i-1] == DAY_ARG {
			day, err = strconv.Atoi(arg)
		}

		if args[i-1] == INPUT_ARG {
			inputPath = arg
		}

		if err != nil {
			break
		}
	}

	if err != nil || day == 0 || inputPath == "" {
		panic("Could not read arguments")
	}

	buffer, err := os.ReadFile(inputPath)
	if err != nil {
		panic("Could not read input!")
	}

	input := string(buffer)

	return day, input
}

func main() {
	day, input := parseArgs()

	switch day {
	case 6:
		day6.Solve(input)
	case 7:
		day7.Solve(input)
	default:
		panic("Invalid day specified")
	}
}
