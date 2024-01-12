package day03

import (
	"regexp"
	"slices"
	"strconv"
	"unicode"

	"advent-of-code/util"
)

var symbols = []rune{'@', '#', '*', '-', '%', '+', '=', '$', '/', '&'}
var re = regexp.MustCompile(`\d+`)

func findAllNums(lines []string) map[int]*util.RangeMap {
	numPos := make(map[int]*util.RangeMap)

	for row, line := range lines {
		indices := re.FindAllStringIndex(line, -1)
		for _, loc := range indices {
			if _, ok := numPos[row]; !ok {
				numPos[row] = &util.RangeMap{}
			}

			num, _ := strconv.Atoi(line[loc[0]:loc[1]])
			numPos[row].Set(loc[0], loc[1], num)
		}
	}

	return numPos
}

func (s Solver) P1(lines []string) any {
	numPos := findAllNums(lines)

	res := 0
	included := make(map[int]bool)
	n, m := len(lines), len(lines[0])

	for row, line := range lines {
		for col, c := range line {
			if !slices.Contains(symbols, c) {
				continue
			}

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
						res += num
						included[(nr*m)+left] = true
					}
				}
			}
		}
	}

	return res
}
