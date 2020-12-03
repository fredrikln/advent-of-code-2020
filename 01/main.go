package main

import (
	"fmt"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadInputLinesAsNumbers()

	result1 := part1(input)
	fmt.Println(result1)

	result2 := part2(input)
	fmt.Println(result2)
}

func part1(expenses []int) int {
	fmt.Println("Part 1")

	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			first := expenses[i]
			second := expenses[j]

			if first+second == 2020 {
				return first * second
			}
		}
	}

	return 0
}

func part2(expenses []int) int {
	fmt.Println("Part 2")

	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			for k := j + 1; k < len(expenses); k++ {
				first := expenses[i]
				second := expenses[j]
				third := expenses[k]

				if first+second+third == 2020 {
					return first * second * third
				}
			}
		}
	}

	return 0
}
