package day9

import (
	"aoc-2022-1/utils"
	"math"
	"strconv"
	"strings"
)

func Part2() int {
	input, err := utils.GetInput("day9")
	if err != nil {
		panic(err)
	}
	return SimulateLargRopePostition(input)
}

func SimulateLargRopePostition(input []string) int {
	lastKnotPositionsVisited := make(map[[2]int]bool)
	lastKnotPositionsVisited[[2]int{0, 0}] = true

	linePositions := make([][2]int, 10)
	for i := 0; i <= 9; i++ {
		linePositions[i] = [2]int{0, 0}
	}

	for _, line := range input {
		split := strings.Split(line, " ")
		command := split[0]
		distance, _ := strconv.Atoi(split[1])
		for i := 0; i < distance; i++ {
			move := [2]int{0, 0}
			switch command {
			case "U":
				move = [2]int{0, 1}
			case "D":
				move = [2]int{0, -1}
			case "L":
				move = [2]int{-1, 0}
			case "R":
				move = [2]int{1, 0}
			default:
			}
			calculateKnotsPositions(move, linePositions)
			lastKnot := linePositions[9]
			if ok := lastKnotPositionsVisited[lastKnot]; !ok {
				lastKnotPositionsVisited[lastKnot] = true
			}
		}
	}
	return len(lastKnotPositionsVisited)
}

func calculateKnotsPositions(direction [2]int, linePositions [][2]int) {
	for k, t := range linePositions {
		tx, ty := t[0], t[1]
		// if k == 0, it's the first knot which is head
		if k == 0 {
			linePositions[k][0] = tx + direction[0]
			linePositions[k][1] = ty + direction[1]
			continue
		}
		pt := linePositions[k-1]
		hx, hy := pt[0], pt[1]
		if isContact(hx, hy, tx, ty) {
			return
		}
		// we're calculating the difference between the previous knot and the current knot
		// and we're moving the current knot by 1 in the direction of the difference
		// for both x and y axis
		dx, dy := 0, 0
		if hx != tx {
			dx = (hx - tx) / int(math.Abs(float64(hx-tx)))
		}
		if hy != ty {
			dy = (hy - ty) / int(math.Abs(float64(hy-ty)))
		}

		linePositions[k][0] += int(dx)
		linePositions[k][1] += int(dy)
	}
}
