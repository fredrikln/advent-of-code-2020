package main

import (
	"fmt"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadInput()

	fmt.Println("Part 01")
	result1 := part1(input)
	fmt.Println(result1)

	fmt.Println("Part 02")
	result2 := part2(input)
	fmt.Println(result2)
}

func parseLine(line string) (min int, max int, letter string, password string) {
	fmt.Sscanf(line, "%d-%d %1s: %s", &min, &max, &letter, &password)

	return min, max, letter, password
}

func part1(input []string) int {
	countValid := 0

	for _, line := range input {
		min, max, letter, password := parseLine(line)

		counts := make(map[string]int)
		for i := 0; i < len(password); i++ {
			curLetter := string(password[i])

			counts[curLetter]++
		}

		if counts[letter] >= min && counts[letter] <= max {
			countValid++
		}
	}

	return countValid
}

func part2(input []string) int {
	countValid := 0

	for _, line := range input {
		min, max, letter, password := parseLine(line)

		if (string(password[min-1]) == letter || string(password[max-1]) == letter) && string(password[min-1]) != string(password[max-1]) {
			countValid++
		}
	}

	return countValid
}
