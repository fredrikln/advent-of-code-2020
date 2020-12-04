package main

import (
	"fmt"
	"strings"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadInput()

	fmt.Println("Part 1")
	result1 := part1(input, 3, 1)
	fmt.Println(result1)

	fmt.Println("Part 2")
	result2 := part2(input)
	fmt.Println(result2)
}

func part1(input []string, xSpeed int, ySpeed int) int {
	treesHit := 0
	x := 0
	for y := 0; y < len(input); y += ySpeed {
		line := strings.Split(input[y], "")
		if line[x%len(input[y])] == "#" {
			treesHit++
		}
		x += xSpeed
	}

	return treesHit
}

type slope struct {
	xSpeed int
	ySpeed int
}

func part2(input []string) int {
	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product := 1
	for _, slope := range slopes {
		product *= part1(input, slope.xSpeed, slope.ySpeed)
	}

	return product
}
