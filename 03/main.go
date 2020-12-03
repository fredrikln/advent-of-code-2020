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

type coord struct {
	x int
	y int
}

func part1(input []string, xSpeed int, ySpeed int) int {
	trees := make(map[coord]bool)

	for y, line := range input {
		for x, tile := range strings.Split(line, "") {
			if tile == "#" {
				trees[coord{x, y}] = true
			}
		}
	}

	x := 0
	treesHit := 0
	for y := 0; y < len(input); y += ySpeed {
		coord := coord{x: x % len(input[0]), y: y}
		_, ok := trees[coord]
		if ok {
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
		product = product * part1(input, slope.xSpeed, slope.ySpeed)
	}

	return product
}
