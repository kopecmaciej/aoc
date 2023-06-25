package day4

import (
	"testing"
)

func TestPart1Success(t *testing.T) {
	input := []string{
		"2-3,1-6",
		"2-9,6-8",
		"3-5, 6-9",
		"4-20,12-18",
	}

	expected := 3
	actual := part1Logic(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPart1Failure(t *testing.T) {
	input := []string{
		"1-3,5-7",
		"2-4,6-8",
		"3-5,7-9",
	}

	expected := 2
	actual := part1Logic(input)

	if actual == expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
