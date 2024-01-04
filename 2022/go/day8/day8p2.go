package day8

import (
	"aoc-2022-1/utils"
)

func Day8Part2() int {
	input, err := utils.GetInput("day8")
	if err != nil {
		panic(err)
	}

	return calculateMaxScenicScore(input)
}

func calculateMaxScenicScore(input []string) int {
	treeMatrix := make([][]int, len(input))
	columns := make([][]byte, len(input))
	rows := make([][]byte, len(input))

	//populate with 0s
	for i := 0; i < len(input); i++ {
		row := make([]int, len(input))
		for j := 0; j < len(input); j++ {
			row[j] = 1
		}
		treeMatrix[i] = row
	}

	for i, row := range input {
		for j, column := range row {
			rows[i] = append(rows[i], byte(column))
			columns[j] = append(columns[j], byte(column))
		}
	}

	for i := 0; i < len(input); i++ {
		countMaxScenicScore(rows[i], columns[i], i, treeMatrix)
	}
	maxScenicScore := 0
	for _, row := range treeMatrix {
		for _, column := range row {
			if column > maxScenicScore {
				maxScenicScore = column
			}
		}
	}

	return maxScenicScore
}

func countMaxScenicScore(row, column []byte, inc int, treeMatrix [][]int) {
	//vertical
	for i, val := range row {
		switch i {
		case 0:
			score := countScore(row[i+1:], val, false)
			treeMatrix[inc][i] *= score
		case len(row) - 1:
			score := countScore(row[:i], val, true)
			treeMatrix[inc][i] *= score
		default:
			score := countScore(row[:i], val, true)
			score *= countScore(row[i+1:], val, false)
			treeMatrix[inc][i] *= score
		}
	}

	//horizontal
	for i, val := range column {
		switch i {
		case 0:
			score := countScore(column[i+1:], val, false)
			treeMatrix[i][inc] *= score
		case len(column) - 1:
			score := countScore(column[:i], val, true)
			treeMatrix[i][inc] *= score
		default:
			score := countScore(column[:i], val, true)
			score *= countScore(column[i+1:], val, false)
			treeMatrix[i][inc] *= score
		}
	}
}

func countScore(arr []byte, curr byte, reverse bool) int {
	if reverse {
    arr = reverseSlice(append([]byte{}, arr...))
	}
	score := 0
	for _, val := range arr {
		if curr > val {
			score++
		} else {
			score++
			break
		}
	}
	if score == 0 {
		return 1
	}
	return score
}

func reverseSlice(arr []byte) []byte {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}
