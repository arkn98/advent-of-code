package day01

import (
	"strings"
	"unicode"
)

var nums = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func calc2(line string) int {
	digits := make([]rune, 0)
	for i, c := range line {
		if unicode.IsDigit(c) {
			digits = append(digits, c)
		} else {
			for n, v := range nums {
				if strings.HasPrefix(line[i:], n) {
					digits = append(digits, v)
				}
			}
		}
	}

	left := int(digits[0] - '0')
	right := int(digits[len(digits)-1] - '0')

	return (left * 10) + right
}

func (s Solver) P2(in string) any {
	return doCalc(in, calc2)
}
