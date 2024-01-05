package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"unicode"

	"advent-of-code/util"
)

var symbols []rune

func main() {
	file, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// load file into a []string
	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	res := 0
	symbols = []rune{'@', '#', '*', '-', '%', '+', '=', '$', '/', '&'}

	if os.Args[1] == "1" {
		res = calc(true, input)
	} else {
		res = calc(false, input)
	}

	fmt.Println(res)
}

func calc(p1 bool, input []string) int {
	n, m := len(input), len(input[0])
	res := 0

	positions := make(map[int]*util.RangeMap)

	for r, line := range input {
		num, l := 0, -1

		for c, v := range line {
			if unicode.IsDigit(v) {
				if l == -1 {
					l = c
				}
				num = (num * 10) + int(v-'0')
			} else if num > 0 {
				if _, ok := positions[r]; !ok {
					positions[r] = &util.RangeMap{}
				}

				positions[r].Set(l, c-1, num)
				num, l = 0, -1
			}
		}

		if num > 0 {
			if _, ok := positions[r]; !ok {
				positions[r] = &util.RangeMap{}
			}

			positions[r].Set(l, m-1, num)
		}
	}

	if p1 {
		included := make(map[int]bool)

		for row, line := range input {
			for col, c := range line {
				if !slices.Contains(symbols, c) {
					continue
				}

				for tr := row - 1; tr <= row+1; tr++ {
					for tc := col - 1; tc <= col+1; tc++ {
						if tr < 0 || tr >= n || tc < 0 || tc >= m ||
							!unicode.IsDigit(rune(input[tr][tc])) {
							continue
						}

						num, left, _, _ := positions[tr].Get(tc)

						if _, ok := included[(tr*m)+left]; !ok {
							res += num
							included[(tr*m)+left] = true
						}
					}
				}
			}
		}
	} else {
		for row, line := range input {
			for col, c := range line {
				if !(c == '*') {
					continue
				}

				included := make(map[int]bool)
				found := 0
				ratio := 1

				for tr := row - 1; tr <= row+1; tr++ {
					for tc := col - 1; tc <= col+1; tc++ {
						if tr < 0 || tr >= n || tc < 0 || tc >= m ||
							!unicode.IsDigit(rune(input[tr][tc])) {
							continue
						}

						num, left, _, _ := positions[tr].Get(tc)

						if _, ok := included[(tr*m)+left]; !ok {
							found++
							ratio *= num
							included[(tr*m)+left] = true
						}
					}
				}

				if found == 2 {
					res += ratio
				}
			}
		}
	}

	return res
}
