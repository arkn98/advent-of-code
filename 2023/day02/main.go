package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"advent-of-code/util"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		gameInfo, game, _ := strings.Cut(line, ":")
		gameId, _ := strconv.Atoi(strings.Split(gameInfo, " ")[1])
		sets := strings.Split(game, ";")

		counts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, set := range sets {
			colors := strings.Split(set, ",")

			for _, color := range colors {
				inner := strings.Fields(color)
				count, _ := strconv.Atoi(inner[0])
				colorName := inner[1]

				counts[colorName] = util.Max(counts[colorName], count)
			}
		}

		if p1 {
			if counts["red"] <= 12 && counts["green"] <= 13 && counts["blue"] <= 14 {
				res += gameId
			}
		} else {
			res += counts["red"] * counts["green"] * counts["blue"]
		}
	}

	return res
}
