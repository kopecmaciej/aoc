package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func GetTopThreeMaxCalories() int {
	sumAllCalories := []int{}
	calories := getAllCalories()
	for _, elfCalories := range calories {
		sumSingleElfsCalories := 0
		for _, calorie := range elfCalories {
			sumSingleElfsCalories += calorie
		}
		sumAllCalories = append(sumAllCalories, sumSingleElfsCalories)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sumAllCalories)))
	fmt.Println("sumAllCalories: ", sumAllCalories)
	sumTopThreeMaxCalories := 0
	for i := 0; i < 3; i++ {
		sumTopThreeMaxCalories += sumAllCalories[i]
	}

	return sumTopThreeMaxCalories
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
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
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
