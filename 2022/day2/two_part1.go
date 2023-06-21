package day2

import (
	"bufio"
	"os"
	"strings"
)

type Symbol int

var (
	WinPair = [][]Symbol{
		{Rock, Scissors},
		{Scissors, Paper},
		{Paper, Rock},
	}
)

const (
	Rock Symbol = iota + 1
	Paper
	Scissors
)

type RoundResult int

const (
	Lose RoundResult = iota * 3
	Tie
	Win
)

func Day2PartOne() int {
	input := FetchInput()
	score := 0
	for _, pair := range input {
		getTotalScore(pair[0], pair[1], &score)
	}

	return score
}

// FetchInput reads the input file and returns a slice of pairs of strings
// represent first the opponent's choice and then my choice.
func FetchInput() [][]string {
	pwd, err := os.Getwd()
	file, err := os.Open(pwd + "/day2/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pairs := [][]string{}
	for scanner.Scan() {
		columns := strings.Split(scanner.Text(), " ")
		opponent := columns[0]
		me := columns[1]
		pairs = append(pairs, []string{opponent, me})
	}

	return pairs
}

func getTotalScore(opponent, me string, score *int) int {
	*score += compareAndSetScore(opponent, me)
	return *score
}

func compareAndSetScore(opponent, me string) int {
	opponentSymbol := ConvertToSymbol(opponent)
	mineSymbol := ConvertToSymbol(me)

	score := 0

	if opponentSymbol == mineSymbol {
		score = int(Tie)
	} else {
		for _, pair := range WinPair {
			if mineSymbol == pair[0] && opponentSymbol == pair[1] {
				score = int(Win)
			}
		}
	}
	score = score + int(mineSymbol)

	return score
}

func ConvertToSymbol(letter string) Symbol {
	switch letter {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		panic("invalid letter")
	}
}
