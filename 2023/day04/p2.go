package day04

func (s Solver) P2(input []string) any {
	n := len(input)
	counts := make([]int, n)
	for i := 0; i < n; i++ {
		counts[i] = 1
	}

	for i, line := range input {
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
