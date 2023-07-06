package day7

import (
	"aoc-2022-1/utils"
	"strconv"
	"strings"
)

const (
	lsPattern  = "$ ls"
	cdPattern  = "$ cd"
	dirPattern = "dir "
)

func Day7Part1() int {
	input, err := utils.GetInput("day7")
	if err != nil {
		panic(err)
	}

	return findTotalDirsSize(input)
}

func findTotalDirsSize(input []string) int {
	total := 0
	dirsStack := []int{0}
	for i, line := range input {
		switch {
		case strings.HasPrefix(line, lsPattern):
		case strings.HasPrefix(line, dirPattern):
		case strings.HasPrefix(line, cdPattern):
			if line == cdPattern+" .." {
				if len(dirsStack) == 0 {
					dirsStack = append(dirsStack, 0)
				}
				lastDirSize := dirsStack[len(dirsStack)-1]
				if lastDirSize <= 100000 {
					total += lastDirSize
				}
				dirsStack = dirsStack[:len(dirsStack)-1]
				dirsStack[len(dirsStack)-1] += lastDirSize
			} else {
				dirsStack = append(dirsStack, 0)
			}
		default:
			size, err := getFileSize(line)
			if err != nil {
				continue
			}
			dirsStack[len(dirsStack)-1] += size
			if i == len(input)-1 {
				total += dirsStack[len(dirsStack)-1]
			}
		}
	}
	return total
}

func getFileSize(line string) (int, error) {
	sizeStr := strings.Split(line, " ")[0]
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return 0, err
	}
	return size, nil
}
