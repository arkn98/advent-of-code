package day03

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test.txt
var in string

var solver = Solver{}

const (
	res1 = 4361
	res2 = 467835
)

func TestSolver_P1(t *testing.T) {
	assert.Equal(t, res1, solver.P1(in).(int))
}

func TestSolver_P2(t *testing.T) {
	assert.Equal(t, res2, solver.P2(in).(int))
}
