package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	res := 0

	if os.Args[1] == "1" {
		res = calc(true, file)
	} else {
		res = calc(false, file)
	}

	fmt.Println(res)
}

func calc(p1 bool, file *os.File) int {
	res := 0
	nums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n := len(line)

		l, r, lVal, rVal := n, -1, 0, 0

		if !p1 {
			for k, v := range nums {
				pos := strings.Index(line, k)
				if pos != -1 && pos < l {
					l = pos
					lVal = v
				}
			}
		}

		for i := 0; i < l; i++ {
			if unicode.IsDigit(rune(line[i])) {
				l = i
				lVal = int(line[i] - '0')
				break
			}
		}

		if !p1 {
			for k, v := range nums {
				pos := strings.LastIndex(line, k)
				if pos != -1 && pos > r {
					r = pos
					rVal = v
				}
			}
		}

		for i := n - 1; i > r; i-- {
			if unicode.IsDigit(rune(line[i])) {
				r = i
				rVal = int(line[i] - '0')
				break
			}
		}

		res += (lVal * 10) + rVal
	}

	return res
}
