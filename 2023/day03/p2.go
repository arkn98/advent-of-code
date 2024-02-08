package day03

import (
	"unicode"

	"advent-of-code/util"
)

func (s Solver) P2(in string) any {
	lines := util.PrepInput(in)
	numPos := findAllNums(lines)

	res := 0
	n, m := len(lines), len(lines[0])

	for row, line := range lines {
		for col, c := range line {
			if !(c == '*') {
				continue
			}

			included := make(map[int]bool)
			found := 0
			ratio := 1

			for _, dr := range []int{-1, 0, 1} {
				for _, dc := range []int{-1, 0, 1} {
					nr := row + dr
					nc := col + dc

					if nr < 0 || nr >= n || nc < 0 || nc >= m ||
						!unicode.IsDigit(rune(lines[nr][nc])) {
						continue
					}

					num, left, _, _ := numPos[nr].Get(nc)

					if _, ok := included[(nr*m)+left]; !ok {
						found++
						ratio *= num
						included[(nr*m)+left] = true
					}
				}
			}

			if found == 2 {
				res += ratio
			}
		}
	}

	return res
}
