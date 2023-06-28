package day4

import (
	"aoc-2022-1/utils"
)

func Part2() int {
	input, _ := utils.GetInput(day)

	return overlapAll(input)
}

func overlapAll(input []string) (overlaps int) {
	for _, line := range input {
		a, b, c, d := getBoundaryValues(line)

		switch {
		case a <= c && b >= c:
			overlaps++
		case a >= c && a <= d:
			overlaps++
		default:
			continue
		}

	}

	return overlaps
}
