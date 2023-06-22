package day1

import (
	"fmt"
	"sort"
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

