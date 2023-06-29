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

  var unique []rune
	curr := 0
	for {
		if c, _, err := data.ReadRune(); err == nil {
			if err == io.EOF {
				break
			} else {
				if len(unique) == 4 {
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

func checkIfSliceUnique(unique []rune) bool {
	for i := 0; i < len(unique); i++ {
		for j := i + 1; j < len(unique); j++ {
			if unique[i] == unique[j] {
				return false
			}
		}
	}
	return true
}
