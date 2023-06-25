package day3

import (
	"aoc-2022-1/utils"
	"strings"
	"unicode"
)

func Day3PartOne() int {
	input, err := utils.GetInput()
	if err != nil {
		panic(err)
	}
	sum := 0
	duplicates := getAllDuplicates(input)
	for _, duplicate := range duplicates {
		sum += GetValueFromLetter(rune(duplicate[0]))
	}
	return sum
}

func getAllDuplicates(input []string) []string {
	duplicates := make([]string, 0)

	for _, line := range input {
		firstRucksack := line[:len(line)/2]
		secondRucksack := line[len(line)/2:]

		firstRucksackSet := make(map[rune]bool)
		for _, char := range firstRucksack {
			firstRucksackSet[char] = true
		}

		for firstRucksackItem := range firstRucksackSet {
			contains := strings.Contains(secondRucksack, string(firstRucksackItem))
			if contains {
				duplicates = append(duplicates, string(firstRucksackItem))
			}
		}
	}
	return duplicates
}

func GetValueFromLetter(letter rune) int {
	if unicode.IsUpper(letter) {
		return int(letter) - 38
	}
	return int(letter) - 96
}
