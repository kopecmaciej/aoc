package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
  fmt.Println("Max calories: ", GetMaxcallories())
}

func GetMaxcallories() int {
	calories := getAllCalories()
	maxCalories := 0
	for _, elfCalories := range calories {
		sumElfsCalories := 0
		for _, calorie := range elfCalories {
			sumElfsCalories += calorie
		}
		if sumElfsCalories > maxCalories {
			maxCalories = sumElfsCalories
		}
	}
	return maxCalories
}

func getAllCalories() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	var allElfsCalories [][]int
	currIndex := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			currIndex++
		} else {
			if len(allElfsCalories) < currIndex+1 {
				allElfsCalories = append(allElfsCalories, []int{})
			}
			intCalories, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("Error converting to int")
				os.Exit(1)
			}

			allElfsCalories[currIndex] = append(allElfsCalories[currIndex], intCalories)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	return allElfsCalories
}
