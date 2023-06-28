package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
)

func GetInput(year string) ([]string, error) {
	_, b, _, _ := runtime.Caller(0)
  root := filepath.Join(filepath.Dir(b), "../")
	file, err := os.Open(root + "/" + year + "/input")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input, nil
}
