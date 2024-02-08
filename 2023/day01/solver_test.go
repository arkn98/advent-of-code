package day01

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test1.txt
var in1 string

//go:embed test2.txt
var in2 string

var solver = Solver{}

const (
	res1 = 142
	res2 = 281
)

func TestSolver_P1(t *testing.T) {
	assert.Equal(t, res1, solver.P1(in1).(int))
}

func TestSolver_P2(t *testing.T) {
	assert.Equal(t, res2, solver.P2(in2).(int))
}
