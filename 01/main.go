package main

import (
	"fmt"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func part1(expenses []int) int {
	fmt.Println("Part 1")
	for _, i := range expenses {
		for _, j := range expenses {
			if (i + j) == 2020 {
				return i * j
			}
		}
	}

	return 0
}

func part2(expenses []int) int {
	fmt.Println("Part 2")

	for _, i := range expenses {
		for _, j := range expenses {
			for _, k := range expenses {
				if (i + j + k) == 2020 {
					return i * j * k
				}
			}
		}
	}

	return 0
}

func main() {
	input := utils.ReadInputLinesAsNumbers()

	result1 := part1(input)
	fmt.Println(result1)

	result2 := part2(input)
	fmt.Println(result2)
}
