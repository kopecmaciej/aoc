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

	return searchForUniqueMarker(data, 14)
}
