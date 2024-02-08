package day01

import (
	"unicode"

	"advent-of-code/util"
)

func calc(line string) int {
	digits := make([]rune, 0)
	for _, c := range line {
		if unicode.IsDigit(c) {
			digits = append(digits, c)
		}
	}

	left := int(digits[0] - '0')
	right := int(digits[len(digits)-1] - '0')

	return (left * 10) + right
}

func doCalc(in string, f func(string) int) any {
	res := 0

	lines := util.PrepInput(in)
	for _, line := range lines {
		res += f(line)
	}

	return res
}

func (s Solver) P1(in string) any {
	return doCalc(in, calc)
}
