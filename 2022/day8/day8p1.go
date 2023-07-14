package day8

import (
	"aoc-2022-1/utils"
)

func Day8Part1() int {
	input, err := utils.GetInput("day8")
	if err != nil {
		panic(err)
	}

	return countVisableTrees(input)
}

func countVisableTrees(input []string) int {
	total := 0

	var treeMatrix [][]byte
	columns := make([][]byte, len(input))
	rows := make([][]byte, len(input))

	//populate with 0s
	for i := 0; i < len(input); i++ {
		var row []byte
		for j := 0; j < len(input); j++ {
			if i == 0 || i == len(input)-1 {
				row = append(row, 1)
				continue
			}
			row = append(row, 0)
		}
		treeMatrix = append(treeMatrix, row)
	}

	for i, row := range input {
		for j, column := range row {
			rows[i] = append(rows[i], byte(column))
			columns[j] = append(columns[j], byte(column))
		}
	}

	for i := 0; i < len(input); i++ {
		if i == 0 || i == len(input)-1 {
			continue
		}
		checkVisability(rows[i], columns[i], i, treeMatrix)
	}

	for _, row := range treeMatrix {
		for _, column := range row {
			if column == 1 {
				total++
			}
		}
	}

	return total
}

func checkVisability(row, column []byte, inc int, matrix [][]byte) {
	//horizontal
	leftBiggest := row[0]
	rightBiggest := row[len(row)-1]
	for i := 0; i < len(row); i++ {
		curr := row[i]
		if i == 0 || i == len(row)-1 {
			matrix[inc][i] = 1
			continue
		}
		if curr > leftBiggest {
			matrix[inc][i] = 1
			leftBiggest = curr
			continue
		}
		if curr > rightBiggest {
			rightRow := row[i+1:]
			//sort bytes
			max := rightRow[0]
			for _, v := range rightRow {
				if v > max {
					max = v
				}
			}

			if curr > max {
				matrix[inc][i] = 1
			}
		}
	}

	//vertical
	topBiggest := column[0]
	bottomBiggest := column[len(column)-1]
	for i := 0; i < len(column); i++ {
		curr := column[i]
		if i == 0 || i == len(column)-1 {
			matrix[i][inc] = 1
			continue
		}
		if curr > topBiggest {
			matrix[i][inc] = 1
			topBiggest = curr
			continue
		}
		if curr > bottomBiggest {
			bottomRow := column[i+1:]
			max := bottomRow[0]
			for _, v := range bottomRow {
				if v > max {
					max = v
				}
			}
			if curr > max {
				matrix[i][inc] = 1
			}
		}
	}
}
