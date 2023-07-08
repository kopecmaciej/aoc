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

type stack struct {
	data []int
}

func Day7Part1() int {
	input, err := utils.GetInput("day7")
	if err != nil {
		panic(err)
	}

	return findTotalDirsSize(input)
}

func findTotalDirsSize(input []string) int {
	total := 0
	dirsStack := stack{[]int{0}}
	for i, line := range input {
		switch {
		case strings.HasPrefix(line, lsPattern):
		case strings.HasPrefix(line, dirPattern):
		case strings.HasPrefix(line, cdPattern):
			if line == cdPattern+" .." {
				if dirsStack.len() == 0 {
					dirsStack.push(0)
				}
				lastDirSize := dirsStack.pop()
				if lastDirSize <= 100000 {
					total += lastDirSize
				}
        dirsStack.addToLast(lastDirSize)
			} else {
				dirsStack.push(0)
			}
		default:
			size, err := getFileSize(line)
			if err != nil {
				continue
			}
      dirsStack.addToLast(size)
			if i == len(input)-1 {
				total += dirsStack.peekLast()
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

func (s stack) len() int {
	return len(s.data)
}

func (s *stack) push(i int) {
	s.data = append(s.data, i)
}

func (s *stack) pop() int {
	i := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return i
}

func (s stack) addToLast(i int) {
	s.data[len(s.data)-1] += i
}

func (s stack) peekLast() int {
	return s.data[len(s.data)-1]
}
