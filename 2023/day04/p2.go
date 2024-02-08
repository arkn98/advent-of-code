package day04

import "advent-of-code/util"

func (s Solver) P2(in string) any {
	lines := util.PrepInput(in)

	n := len(lines)
	counts := make([]int, n)
	for i := 0; i < n; i++ {
		counts[i] = 1
	}

	for i, line := range lines {
		matchCount := processLine(line)

		for j := 1; j <= matchCount; j++ {
			counts[i+j] += counts[i]
		}
	}

	res := 0
	for _, v := range counts {
		res += v
	}

	return res
}
