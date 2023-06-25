package day4

import (
	"aoc-2022-1/utils"
	"strconv"
	"strings"
)

func Part1() int {
	input, _ := utils.GetInput()

	return part1Logic(input)
}

func part1Logic(input []string) int {
	numOfOverlaps := 0

	for _, line := range input {
		line = strings.Replace(line, " ", "", -1)
		split := strings.Split(line, ",")
		first := split[0]
		second := split[1]

		a, _ := strconv.Atoi(strings.SplitN(first, "-", 2)[0])
		b, _ := strconv.Atoi(strings.SplitN(first, "-", 2)[1])
		c, _ := strconv.Atoi(strings.SplitN(second, "-", 2)[0])
		d, _ := strconv.Atoi(strings.SplitN(second, "-", 2)[1])

		switch {
		case a >= c && b <= d:
			numOfOverlaps++
		case a <= c && b >= d:
			numOfOverlaps++
		default:
			continue
		}
	}

	return numOfOverlaps
}
