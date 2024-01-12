package day02

import (
	_ "embed"
	"testing"

	"advent-of-code/util"

	"github.com/stretchr/testify/assert"
)

//go:embed data/ex1.txt
var in string

var solver = Solver{}

const (
	res1 = 8
	res2 = 2286
)

func TestSolver_P1(t *testing.T) {
	assert.Equal(t, res1, solver.P1(util.PrepInput(in)).(int))
}

func TestSolver_P2(t *testing.T) {
	assert.Equal(t, res2, solver.P2(util.PrepInput(in)).(int))
}
