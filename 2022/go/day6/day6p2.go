package day6

import (
	"bufio"
	"os"
	"strings"
)

func Day6Part2() int {
	buff, err := os.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}
	data := bufio.NewReader(strings.NewReader(string(buff)))

	markLen := 14
  peek := 0
	for {
		mark, err := data.Peek(markLen)
    mark = mark[peek:markLen]
		if err != nil {
			return 0
		}
		if checkIfSliceUnique(mark) {
			break
		}
		markLen++
    peek++
	}

	return markLen
}
