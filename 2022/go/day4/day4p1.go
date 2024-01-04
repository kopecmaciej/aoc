package day4

import (
	"aoc-2022-1/utils"
	"strconv"
	"strings"
)

const (
	day = "day4"
)

func Part1() int {
	input, _ := utils.GetInput(day)

	return checkNumOfOverlaps(input)
}

func checkNumOfOverlaps(input []string) (overlaps int) {
	for _, line := range input {
		a, b, c, d := getBoundaryValues(line)

		switch {
		case a >= c && b <= d:
			overlaps++
		case a <= c && b >= d:
			overlaps++
		default:
			continue
		}
	}

	return overlaps
}

func getBoundaryValues(line string) (int, int, int, int) {
	line = strings.Replace(line, " ", "", -1)
	split := strings.Split(line, ",")
	first := split[0]
	second := split[1]

	a, _ := strconv.Atoi(strings.SplitN(first, "-", 2)[0])
	b, _ := strconv.Atoi(strings.SplitN(first, "-", 2)[1])
	c, _ := strconv.Atoi(strings.SplitN(second, "-", 2)[0])
	d, _ := strconv.Atoi(strings.SplitN(second, "-", 2)[1])

	return a, b, c, d
}
