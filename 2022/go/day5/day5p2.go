package day5

import (
	"aoc-2022-1/utils"
	"strconv"
	"strings"
)

func Day5Part2() []string {
	input, _ := utils.GetInput("day5")

	return moveCargoP2(input)
}

func moveCargoP2(input []string) []string {
	crates := make(map[int][]string, 9)
	topCrates := make([]string, numOfColumns)

	for _, line := range input {
		if strings.Contains(line, "[") {
			pushToCorrectColumn(line, crates)
		} else if strings.Contains(line, "move") {
			parseCommand2(line, crates)
		}
	}

	for i, crate := range crates {
		topCrates[i-1] = crate[0]
	}

	return topCrates
}

func parseCommand2(command string, crates map[int][]string) {
	numToMove, _ := strconv.Atoi(strings.Split(command, " ")[1])
	fromColumn, _ := strconv.Atoi(strings.Split(command, " ")[3])
	toColumn, _ := strconv.Atoi(strings.Split(command, " ")[5])
	currentCargo := crates[fromColumn][0:numToMove]
	crates[fromColumn] = crates[fromColumn][numToMove:]
	crates[toColumn] = prependToString(currentCargo, crates[toColumn])
}

func prependToString(merge ...[]string) []string {
	var result []string
	for _, s := range merge {
		result = append(result, s...)
	}
	return result
}
