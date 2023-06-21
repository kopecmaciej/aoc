package main

import (
	"aoc-2022-1/day1"
	"aoc-2022-1/day2"
	"aoc-2022-1/day3"
	"fmt"
)

func main() {
	exec_day3()
}

func exec_day1() {
	maxCalories := day1.GetMaxcallories()
	fmt.Println("Max calories: ", maxCalories)
	topThreeMaxCalories := day1.GetTopThreeMaxCalories()
	fmt.Println("Top three max calories: ", topThreeMaxCalories)
}

func exec_day2() {
	fmt.Println("Day 2 part 1: ", day2.Day2PartOne())
  fmt.Println("Day 2 part 2: ", day2.Day2PartTwo())
}

func exec_day3() {
  fmt.Println("Day 3 part 1: ", day3.Day3PartOne())
}

