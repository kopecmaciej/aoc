package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7Part1(t *testing.T) {
	output := Day7Part1()
	t.Logf("Day 7 Part 1 result: %d", output)
}

func TestFindTotalDirSize(t *testing.T) {
	testInput := []string{
		"$ cd /",
		"$ ls",
		"1278 test.txt",
		"255 test2.txt",
		"dir 1",
		"dir 2",
		"$ cd 1",
		"$ ls",
		"dir 3",
		"dir 4",
		"1278 test.txt",
		"255 test2.txt",
		"$ cd ..",
		"$ cd 2",
		"$ ls",
		"12342 test.txt",
		"255 test2.txt",
		"$ cd ..",
		"$ cd ..",
		"46890 test2.txt",
	}

	output := findTotalDirsSize(testInput)
	expected := 92346
	assert.Equal(t, expected, output)
}
