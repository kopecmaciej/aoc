package day6

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Day6Part1() int {
	buff, err := os.ReadFile("input")
	if err != nil {
		os.Exit(1)
	}
	data := bufio.NewReader(strings.NewReader(string(buff)))

	return searchForUniqueMarker(data, 4)
}

func searchForUniqueMarker(data *bufio.Reader, markerLen int) int {
	var unique []rune
	curr := 0
	for {
		if c, _, err := data.ReadRune(); err == nil {
			if err == io.EOF {
				break
			} else {
				if len(unique) == markerLen {
					if checkIfSliceUnique(unique) {
						break
					}
					unique = unique[1:]
				}
				unique = append(unique, c)
				curr++
			}
		}
	}
	return curr
}

func checkIfSliceUnique[T string | byte | rune](unique []T) bool {
	uniqueMap := make(map[T]bool)
	for _, val := range unique {
		if _, ok := uniqueMap[val]; ok {
			return false
		}
		uniqueMap[val] = true
	}
	return true
}
