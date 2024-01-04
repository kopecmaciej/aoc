package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8Part1(t *testing.T) {
	result := Day8Part1()
	t.Logf("Result: %d", result)
}

func TestCountVisableTrees(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	result := countVisableTrees(input)
	assert.Equal(t, 21, result)
}

func TestDay8Part2(t *testing.T) {
	result := Day8Part2()
	t.Logf("Result: %d", result)
}

func TestCountMaxScenicScore(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	result := calculateMaxScenicScore(input)
	assert.Equal(t, 16, result)
}
