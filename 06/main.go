package main

import (
	"fmt"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadInput("input.txt")

	fmt.Println("Part 1")
	result := part1(input)
	fmt.Println("Count:", result)

	fmt.Println("Part 2")
	result2 := part2(input)
	fmt.Println("Count:", result2)
}

type group struct {
	people    int
	questions map[string]int
}

func parseGroups(input []string) []group {
	groups := make([]group, 0)

	var countPeople int
	questions := make(map[string]int)

	for _, line := range input {
		if line == "" {
			// save
			groups = append(groups, group{countPeople, questions})

			// reset
			countPeople = 0
			questions = make(map[string]int)

			continue
		}

		countPeople++

		for _, letter := range line {
			questions[string(letter)]++
		}
	}

	groups = append(groups, group{countPeople, questions})

	return groups
}

func part1(input []string) int {
	groups := parseGroups(input)

	var result int

	for _, group := range groups {
		result += len(group.questions)
	}

	return result
}

func part2(input []string) int {
	groups := parseGroups(input)

	var result int

	for _, group := range groups {
		var countAll int
		for _, question := range group.questions {
			if question == group.people {
				countAll++
			}
		}

		result += countAll
	}

	return result
}
