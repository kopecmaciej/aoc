package day9

import "testing"

func Test_Part1(t *testing.T) {
	result := Part1()
	t.Log("Part 1: ", result)
}

func TestSimulateRopePositions(t *testing.T) {
	input := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}
	result := SimulateRopePositions(input)
	if result != 13 {
		t.Error("Expected 13, got ", result)
	}
}

func Test_Part2(t *testing.T) {
	result := Part2()
	t.Log("Part 2: ", result)
}

func TestSimulateLargRopePostition(t *testing.T) {
	input := []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}
	result := SimulateLargRopePostition(input)
	if result != 36 {
		t.Error("Expected 36, got ", result)
	}
}
