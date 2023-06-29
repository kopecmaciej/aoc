package day6

import "testing"

func TestDay6Part1(t *testing.T) {
  output := Day6Part1()
  t.Logf("Day 6 Part 1 result: %d", output)
}

func TestCheckSliceUniqnes(t *testing.T) {
  input := []rune{'a', 'b', 'c', 'd'}
  output := checkIfSliceUnique(input)
  if output != true {
    t.Errorf("Expected true, got %t", output)
  }
}

func TestCheckSliceUniqnes2(t *testing.T) {
  input := []rune{'a', 'b', 'c', 'a'}
  output := checkIfSliceUnique(input)
  if output != false {
    t.Errorf("Expected false, got %t", output)
  }
}
