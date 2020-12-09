package main

import (
	"fmt"
	"math"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadInputLinesAsNumbers()

	fmt.Println("Part 1")
	result := part1(input)
	fmt.Println("First invalid:", result)

	fmt.Println("Part 2")
	result2 := part2(input, result)
	fmt.Println("Sum:", result2)
}

func firstInvalid(numbers []int, preamble int) int {
	numbersSoFar := make([]int, preamble, len(numbers))
	copy(numbersSoFar, numbers[:preamble])

	for _, n := range numbers[preamble:] {
		if !isValidNumber(numbersSoFar, n, preamble) {
			return n
		}

		numbersSoFar = append(numbersSoFar, n)
	}

	return 0
}

func isValidNumber(previous []int, number int, preamble int) bool {
	numbersToConsider := previous[len(previous)-preamble:]

	for i := 0; i < len(numbersToConsider); i++ {
		for j := i + 1; j < len(numbersToConsider); j++ {
			if numbersToConsider[i]+numbersToConsider[j] == number {
				return true
			}
		}
	}

	return false
}

func sum(numbers []int) int {
	var sum int

	for _, i := range numbers {
		sum += i
	}

	return sum
}

func findContiguous(numbers []int, number int) []int {
	for i := 2; i < number/2; i++ {
		for j := 0; j < len(numbers)-i; j++ {
			if sum(numbers[j:j+i]) == number {
				return numbers[j : j+i]
			}
		}
	}

	return make([]int, 0)
}

func minMax(numbers []int) (int, int) {
	min, max := math.MaxInt64, math.MinInt64

	for _, i := range numbers {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	return min, max
}

func part1(numbers []int) int {
	return firstInvalid(numbers, 25)
}

func part2(numbers []int, number int) int {
	contiguous := findContiguous(numbers, number)

	min, max := minMax(contiguous)

	return min + max
}
