package day02

import (
	"strings"

	"advent-of-code/util"
)

func calc2(line string) int {
	_, game, _ := strings.Cut(line, ":")
	sets := strings.Split(game, ";")

	minCounts := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range sets {
		for color, count := range countColors(set) {
			minCounts[color] = util.Max(minCounts[color], count)
		}
	}

	res := 1
	for _, v := range minCounts {
		res *= v
	}
	return res
}

func (s Solver) P2(in []string) any {
	return doCalc(in, calc2)
}
