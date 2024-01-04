package day2

var (
	mapSymbolToResult = map[string]RoundResult{
		"X": Lose,
		"Y": Tie,
		"Z": Win,
	}
)

func Day2PartTwo() int {
	input := FetchInput()
	finalScore := 0
	for _, pair := range input {
		desiredResult := mapSymbolToResult[pair[1]]
		opponentSymbol := ConvertToSymbol(pair[0])

		score := getScore(opponentSymbol, desiredResult)
		finalScore += score
	}

	return finalScore
}

func getScore(opponentSymbol Symbol, desiredResult RoundResult) int {
	score := 0
	switch desiredResult {
	case Win:
		for _, pair := range WinPair {
			if pair[1] == opponentSymbol {
				score += int(Win)
				score += int(pair[0])
			}
		}
		return score
	case Lose:
		for _, pair := range WinPair {
			if pair[0] == opponentSymbol {
				score += int(pair[1])
			}
		}
		return score
	case Tie:
		score += int(Tie)
		score += int(opponentSymbol)
		return score
	default:
		panic("Invalid result")
	}
}
