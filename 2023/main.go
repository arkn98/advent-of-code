package main

import (
	"flag"
	"fmt"
	"os"

	"advent-of-code/day01"
	"advent-of-code/day02"
	"advent-of-code/day03"
	"advent-of-code/day04"
)

type Solver interface {
	P1(in string) any
	P2(in string) any
}

var solvers = map[string]Solver{
	"01": day01.Solver{},
	"02": day02.Solver{},
	"03": day03.Solver{},
	"04": day04.Solver{},
}

func main() {
	task := ""
	part := 1
	inputPath := ""

	flag.StringVar(&task, "task", "", "")
	flag.IntVar(&part, "part", 1, "")
	flag.StringVar(&inputPath, "input", "", "")
	flag.Parse()

	solver, ok := solvers[task]
	if !ok {
		panic(fmt.Errorf("no such task: %v", task))
	}

	b, err := os.ReadFile(inputPath)
	if err != nil {
		panic(fmt.Errorf("cannot read input file: %v", err))
	}

	input := string(b)

	var res any
	switch part {
	case 1:
		res = solver.P1(input)
	case 2:
		res = solver.P2(input)
	default:
		panic(fmt.Errorf("no such part: %v", part))
	}

	fmt.Println(res)
}
