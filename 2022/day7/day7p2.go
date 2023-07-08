package day7

import (
	"aoc-2022-1/utils"
	"regexp"
	"strings"
)

const (
	fileSystemSpace = 70000000
	minNeededSpace  = 30000000
)

func Day7Part2() int {
	input, err := utils.GetInput("day7")
	if err != nil {
		panic(err)
	}

	return findLargestDirToDelete(input)
}

func findLargestDirToDelete(input []string) int {
	dirsStack := []int{0}
	rootDirSize := 0
	for _, line := range input {
		numeric := regexp.MustCompile(`\d`).MatchString(line)
		if !numeric {
			continue
		}
		fileSize, err := getFileSize(line)
		if err != nil {
			continue
		}
		rootDirSize += fileSize
	}
	freeSpace := fileSystemSpace - rootDirSize
	closestSize := rootDirSize
	for _, line := range input {
		switch {
		case strings.HasPrefix(line, lsPattern):
		case strings.HasPrefix(line, dirPattern):
		case strings.HasPrefix(line, cdPattern):
			if line == cdPattern+" .." {
				if len(dirsStack) == 0 {
					dirsStack = append(dirsStack, 0)
				}
				lastDirSize := dirsStack[len(dirsStack)-1]
				if freeSpace+lastDirSize > minNeededSpace && lastDirSize < closestSize {
					closestSize = lastDirSize
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
		}
	}

	return closestSize
}
