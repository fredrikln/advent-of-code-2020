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

func convertRow(row string) int {
	part := row[:7]
	part = strings.ReplaceAll(part, "F", "0")
	part = strings.ReplaceAll(part, "B", "1")

	number, err := strconv.ParseInt(part, 2, 8)
	if err != nil {
		panic(err)
	}

	return int(number)
}

func convertColumn(column string) int {
	part := column[7:]
	part = strings.ReplaceAll(part, "L", "0")
	part = strings.ReplaceAll(part, "R", "1")

	number, err := strconv.ParseInt(part, 2, 8)
	if err != nil {
		panic(err)
	}

	return int(number)
}

func part1(input []string) int {
	highest := 0
	for _, line := range input {
		row := convertRow(line)
		col := convertColumn(line)
		id := row*8 + col

		if id > highest {
			highest = id
		}
	}
	return highest
}

func part2(input []string) int {
	seats := make([]int, 0)

	for _, line := range input {
		row := convertRow(line)
		col := convertColumn(line)
		id := row*8 + col

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
