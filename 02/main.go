package main

import (
	"fmt"
	"strconv"
	"strings"

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

func part1(input []string) int {
	countValid := 0

	for _, line := range input {
		parts := strings.Split(line, " ")

		minMax := strings.Split(parts[0], "-")

		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		letter := string(parts[1][0])

		password := parts[2]

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
		parts := strings.Split(line, " ")

		minMax := strings.Split(parts[0], "-")

		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		letter := string(parts[1][0])

		password := parts[2]

		if (string(password[min-1]) == letter || string(password[max-1]) == letter) && string(password[min-1]) != string(password[max-1]) {
			countValid++
		}
	}

	return countValid
}
