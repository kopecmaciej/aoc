package utils

import (
	"bufio"
	"os"
)

func GetInput() ([]string, error) {
  dir, _ := os.Getwd()
  file, err := os.Open(dir + "/input")
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
