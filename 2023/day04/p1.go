package day04

import (
	"math"
	"strconv"
	"strings"
)

func processLine(line string) int {
	matchCount := 0

	_, numbers, _ := strings.Cut(line, ":")
	winningNums, ourNums, _ := strings.Cut(numbers, "|")
	winning := make(map[int]bool)

	for _, v := range strings.Fields(winningNums) {
		num, _ := strconv.Atoi(v)
		winning[num] = true
	}

	for _, v := range strings.Fields(ourNums) {
		num, _ := strconv.Atoi(v)
		if _, ok := winning[num]; ok {
			matchCount++
		}
	}

	return matchCount
}

func (s Solver) P1(input []string) any {
	res := 0

	for _, line := range input {
		matchCount := processLine(line)
		res += int(math.Pow(2, float64(matchCount-1)))
	}

	return res
}
