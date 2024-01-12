package main

import (
	_ "embed"
	"flag"
	"fmt"

	"advent-of-code/day01"
	"advent-of-code/day02"
	"advent-of-code/day03"
	"advent-of-code/day04"
	"advent-of-code/util"
)

//go:embed day01/data/input.txt
var in01 string

//go:embed day02/data/input.txt
var in02 string

//go:embed day03/data/input.txt
var in03 string

//go:embed day04/data/input.txt
var in04 string

type Solver interface {
	P1(in []string) any
	P2(in []string) any
}

var solvers = map[string]Solver{
	"01": day01.Solver{},
	"02": day02.Solver{},
	"03": day03.Solver{},
	"04": day04.Solver{},
}

var inputs = map[string]string{
	"01": in01,
	"02": in02,
	"03": in03,
	"04": in04,
}

func main() {
	task := ""
	part := 1

	flag.StringVar(&task, "task", "", "")
	flag.IntVar(&part, "part", 1, "")
	flag.Parse()

	solver, ok := solvers[task]
	if !ok {
		panic(fmt.Errorf("no such task: %v", task))
	}

	var res any
	switch part {
	case 1:
		res = solver.P1(util.PrepInput(inputs[task]))
	case 2:
		res = solver.P2(util.PrepInput(inputs[task]))
	default:
		panic(fmt.Errorf("no such part: %v", part))
	}

	fmt.Println(res)
}
