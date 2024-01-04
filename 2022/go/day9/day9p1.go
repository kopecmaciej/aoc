package day9

import (
	"aoc-2022-1/utils"
	"math"
	"strconv"
	"strings"
)

func Part1() int {
	input, err := utils.GetInput("day9")
	if err != nil {
		panic(err)
	}
	return SimulateRopePositions(input)
}

func SimulateRopePositions(input []string) (positions int) {
	//head and tail
	hx, hy, tx, ty := 0, 0, 0, 0

	// positionVisited with 0,0
	positionVisited := make(map[[2]int]bool)
	positionVisited[[2]int{0, 0}] = true
  positions++

	for _, line := range input {
		split := strings.Split(line, " ")
		command := split[0]
		distance, _ := strconv.Atoi(split[1])
		switch command {
		case "U":
			hy += distance
			adjustTailVertical(hx, hy, &tx, &ty, positionVisited, &positions)
		case "D":
			hy -= distance
			adjustTailVertical(hx, hy, &tx, &ty, positionVisited, &positions)
		case "L":
			hx -= distance
			adjustTailHorizontal(hx, hy, &tx, &ty, positionVisited, &positions)
		case "R":
			hx += distance
			adjustTailHorizontal(hx, hy, &tx, &ty, positionVisited, &positions)
		default:
		}
	}
  return positions
}

func adjustTailHorizontal(hx, hy int, tx, ty *int, tailVisited map[[2]int]bool, pos *int) {
	if isContact(hx, hy, *tx, *ty) {
		return
	}
	xDiff := math.Abs(float64(hx - *tx))
	for i := 0; i < int(xDiff)-1; i++ {
		if hx > *tx {
			*tx += 1
			*ty = hy
		} else {
			*tx -= 1
			*ty = hy
		}
		if !tailVisited[[2]int{*tx, *ty}] {
			tailVisited[[2]int{*tx, *ty}] = true
      *pos++
		}
	}
}

func adjustTailVertical(hx, hy int, tx, ty *int, tailVisited map[[2]int]bool, pos *int) {
	if isContact(hx, hy, *tx, *ty) {
		return
	}
	yDiff := math.Abs(float64(hy - *ty))
	for i := 0; i < int(yDiff)-1; i++ {
		if hy > *ty {
			*ty += 1
			*tx = hx
		} else {
			*ty -= 1
			*tx = hx
		}
    if !tailVisited[[2]int{*tx, *ty}] {
      tailVisited[[2]int{*tx, *ty}] = true
      *pos++
    }
	}
}

func isContact(hx, hy, tx, ty int) bool {
	return math.Abs(float64(hx-tx)) <= 1 && math.Abs(float64(hy-ty)) <= 1
}
