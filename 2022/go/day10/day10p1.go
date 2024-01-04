package main

import (
	"aoc-2022-1/utils"
	"log"
	"strconv"

	"golang.org/x/exp/slices"
)

var (
	cycleToRegister = []int{20, 60, 100, 140, 180, 220}
)

func Part1() int {
	input, err := utils.GetInput("day10")
	if err != nil {
		log.Fatal(err)
	}

	return countSignalStrenght(input)
}

func countSignalStrenght(input []string) int {
	signalValue := 1
	globalCycle := 0

	sum := 0

	for _, line := range input {
		command := line[:4]
		var value int
		if len(line) > 5 {
			value, _ = strconv.Atoi(line[5:])
		}

		switch command {
		case "noop":
			sum += computeCycle(value, 1, &globalCycle, &signalValue)
		default:
			sum += computeCycle(value, 1, &globalCycle, &signalValue)
			sum += computeCycle(value, 2, &globalCycle, &signalValue)
		}

	}

	return sum
}

func computeCycle(value int, commandCycle int, globalCycle, signalStrength *int) int {
	*globalCycle++
	if commandCycle > 1 {
		*signalStrength += value
	}

	if slices.Contains(cycleToRegister, *globalCycle) {
		return *globalCycle * *signalStrength
	}
	return 0
}
