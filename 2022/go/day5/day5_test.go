package day5

import (
	"strings"
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

	results := moveCargoP1(input)

	assert.Equal(t, expectedTop, results)
}

func TestDay5Part2(t *testing.T) {
	input := []string{
		"[H]     [P]     [D]     [T]     [Q]",
		"[C] [H] [R]     [P] [S] [C] [T] [M]",
		"[M] [C] [H] [R] [P] [P] [D] [C] [T]",
		"[T] [M] [C] [H] [R] [P] [P] [D] [C]",
		" 1   2   3   4   5   6  7   8   9",
		"move 1 from 4 to 2",
		"move 2 from 7 to 9",
		"move 3 from 5 to 8",
		"move 4 from 3 to 1",
		"move 4 from 1 to 3",
		"move 1 from 8 to 1",
	}

	expectedTop := []string{"D", "R", "P", "H", "R", "S", "D", "P", "T"}

	results := moveCargoP2(input)

	assert.Equal(t, expectedTop, results)
}

func TestInputDay5Part1(t *testing.T) {
	result := Day5Part1()

	t.Log(strings.Join(result, ""))
}

func TestInputDay5Part2(t *testing.T) {
	result := Day5Part2()

	t.Log(strings.Join(result, ""))
}
