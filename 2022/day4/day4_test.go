package day4

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"2-3,1-6",
		"2-9,6-8",
		"3-5, 6-9",
		"4-20,12-18",
	}

	expected := 3
	actual := checkNumOfOverlaps(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"2-4,3-6",
		"2-9,6-8",
		"3-5,6-9",
		"4-20,12-18",
	}

	expected := 3
	actual := overlapAll(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
