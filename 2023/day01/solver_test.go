package day01

import (
	_ "embed"
	"testing"

	"advent-of-code/util"

	"github.com/stretchr/testify/assert"
)

//go:embed data/ex1.txt
var in1 string

//go:embed data/ex2.txt
var in2 string

var solver = Solver{}

const (
	res1 = 142
	res2 = 281
)

func TestSolver_P1(t *testing.T) {
	assert.Equal(t, res1, solver.P1(util.PrepInput(in1)).(int))
}

func TestSolver_P2(t *testing.T) {
	assert.Equal(t, res2, solver.P2(util.PrepInput(in2)).(int))
}
