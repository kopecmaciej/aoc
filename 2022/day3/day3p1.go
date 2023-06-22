package day3

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func Day3PartOne() int {
	sum := 0
	duplicates := getAllDuplicates()
	for _, duplicate := range duplicates {
		sum += GetValueFromLetter(rune(duplicate[0]))
	}
	return sum
}

func GetInput() []string {
	currentDir, err := os.Getwd()
	file, err := os.Open(currentDir + "/day3/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func getAllDuplicates() []string {
	input := GetInput()
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
