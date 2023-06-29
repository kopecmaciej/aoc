package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5Part1(t *testing.T) {
	input := []string{
		"[H]     [P]     [D]     [T]     [Q]",
		"[C] [H] [R]     [P] [D] [C] [T] [M]",
		"[M] [C] [H] [R] [P] [P] [D] [C] [T]",
		"[T] [M] [C] [H] [R] [P] [P] [D] [C]",
		" 1   2   3   4   5   6  7   8   9",
		"move 1 from 4 to 2",
		"move 2 from 7 to 9",
	}

	expectedTop := []string{"H", "R", "P", "H", "D", "D", "D", "T", "C"}

	results := moveCargo(input)

	assert.Equal(t, expectedTop, results)
}

func TestInputDay5Part1(t *testing.T) {
	input := Day5Part1()

	t.Logf("Result: %v", input)
}
