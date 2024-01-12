package day02

import (
	"strconv"
	"strings"
)

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func countColors(set string) map[string]int {
	counts := make(map[string]int)
	colors := strings.Split(set, ",")

	for _, color := range colors {
		inner := strings.Fields(color)
		count, _ := strconv.Atoi(inner[0])
		colorName := inner[1]

		counts[colorName] += count
	}

	return counts
}

func calc(line string) int {
	meta, game, _ := strings.Cut(line, ":")
	sets := strings.Split(game, ";")

	for _, set := range sets {
		for color, count := range countColors(set) {
			if count > limits[color] {
				return 0
			}
		}
	}

	gameID, _ := strconv.Atoi(strings.Split(meta, " ")[1])
	return gameID
}

func doCalc(lines []string, f func(line string) int) any {
	res := 0

	for _, line := range lines {
		res += f(line)
	}

	return res
}

func (s Solver) P1(in []string) any {
	return doCalc(in, calc)
}
