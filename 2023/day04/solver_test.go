package day04

import (
	"advent-of-code/util"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed data/ex1.txt
var in string

var solver = Solver{}

const (
	res1 = 13
	res2 = 30
)

func TestSolver_P1(t *testing.T) {
	assert.Equal(t, res1, solver.P1(util.PrepInput(in)).(int))
}

func TestSolver_P2(t *testing.T) {
	assert.Equal(t, res2, solver.P2(util.PrepInput(in)).(int))
}
