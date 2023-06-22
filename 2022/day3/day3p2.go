package day3

func Day3PartTwo() int {
	input := GetInput()

	groups := [][]string{}

	for i := 0; i < len(input); i += 3 {
		groups = append(groups, input[i:i+3])
	}

	intersections := []rune{}

	for _, group := range groups {
		intersection := findUniqueIntersection(group[0], group[1], group[2])
		intersections = append(intersections, intersection)
	}

	sum := 0

	for _, intersection := range intersections {
		value := GetValueFromLetter(intersection)
		sum += value
	}

	return sum
}

func findUniqueIntersection(one, two, three string) rune {
	oneMap := make(map[rune]bool)

	for _, char := range one {
		oneMap[char] = true
	}

	for _, char := range two {
		if oneMap[char] {
			for range three {
				if oneMap[char] {
					return char
				}
			}
		}
	}

	return 0
}
