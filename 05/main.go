package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadInput("input.txt")

	fmt.Println("Part 1")
	result := part1(input)
	fmt.Println("Highest seat id:", result)

	fmt.Println("Part 2")
	result2 := part2(input)
	fmt.Println("My seat id:", result2)
}

func getID(row string) int {
	row = strings.ReplaceAll(row, "F", "0")
	row = strings.ReplaceAll(row, "B", "1")
	row = strings.ReplaceAll(row, "L", "0")
	row = strings.ReplaceAll(row, "R", "1")

	number, err := strconv.ParseInt(row, 2, 32)
	if err != nil {
		panic(err)
	}

	return int(number)
}

func part1(input []string) int {
	highest := 0
	for _, line := range input {
		id := getID(line)

		if id > highest {
			highest = id
		}
	}
	return highest
}

func part2(input []string) int {
	seats := make([]int, 0)

	for _, line := range input {
		id := getID(line)

		seats = append(seats, id)
	}

	sort.Ints(seats)

	for i, num := range seats {
		if seats[i+1] != num+1 {
			return num + 1
		}
	}

	return 0
}
