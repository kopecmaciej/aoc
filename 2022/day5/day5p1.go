package day5

import (
	"aoc-2022-1/utils"
	"math"
	"strconv"
	"strings"
)

const (
	numOfColumns = 9
)

func Day5Part1() []string {
	input, _ := utils.GetInput("day5")
	return moveCargoP1(input)
}

func moveCargoP1(input []string) []string {
	crates := make(map[int][]string, 9)
	topCrates := make([]string, numOfColumns)

	for _, line := range input {
		if strings.Contains(line, "[") {
			pushToCorrectColumn(line, crates)
		} else if strings.Contains(line, "move") {
			parseCommandOne(line, crates)
		}
	}

	for i, crate := range crates {
		topCrates[i-1] = crate[0]
	}

	return topCrates
}

func pushToCorrectColumn(line string, crates map[int][]string) {
	currColumnt := 0

	for i, char := range line {
		if char >= 65 && char <= 90 {
			currColumnt = int(math.Round(float64(i/4) + 1))
			crates[currColumnt] = append(crates[currColumnt], string(char))
		}
	}
}

func parseCommandOne(command string, crates map[int][]string) {
	numToMovem, _ := strconv.Atoi(strings.Split(command, " ")[1])
	fromColumn, _ := strconv.Atoi(strings.Split(command, " ")[3])
	toColumn, _ := strconv.Atoi(strings.Split(command, " ")[5])
	//redundant, but it's ok
	var currentCargo []string

	currentCargo = crates[fromColumn][0:numToMovem]
	crates[fromColumn] = crates[fromColumn][numToMovem:]
	for _, crate := range currentCargo {
		crates[toColumn] = append([]string{crate}, crates[toColumn]...)
	}
}
