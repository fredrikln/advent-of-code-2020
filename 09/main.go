package main

import (
	"fmt"
	"math"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadInputLinesAsNumbers()

	fmt.Println("Part 1")
	result := part1(input, 25)
	fmt.Println("First invalid:", result)

	fmt.Println("Part 2")
	result2 := part2(input, result)
	fmt.Println("Sum:", result2)
}

func firstInvalid(numbers []int, preamble int) int {
	numbersSoFar := make([]int, preamble, len(numbers))
	copy(numbersSoFar, numbers[:preamble])

	for _, n := range numbers[preamble:] {
		if !isValidNumber(numbersSoFar[len(numbersSoFar)-preamble:], n) {
			return n
		}

		numbersSoFar = append(numbersSoFar, n)
	}

	return 0
}

func isValidNumber(numbersToConsider []int, number int) bool {
	for i := 0; i < len(numbersToConsider); i++ {
		for j := i + 1; j < len(numbersToConsider); j++ {
			if numbersToConsider[i]+numbersToConsider[j] == number {
				return true
			}
		}
	}

	return false
}

func findContiguous(numbers []int, number int) []int {
	runningSum := make([]int, len(numbers))

	var sum int
	for i, n := range numbers {
		sum += n
		runningSum[i] = sum
	}

	for i := 2; i < number/2; i++ {
		for j := i; j < len(numbers)-i; j++ {
			if runningSum[j]-runningSum[j-i] == number {
				// j-i+1 : j-i is the number before the range of numbers that sum up to the one I want
				// j+1 : I want the numbers up to and including one I'm standing on, Go end-index indexing is exclusive
				// Example: I want to know what the sum numbers of index 3 to 5 (three numbers)
				//          Sum at index 5 is 15, sum at index 5-(three)=2 is 3, 15 - 3 is 12
				//          This is the number we want, let's get the actual numbers:
				//          5-3 is before the first number we want, so we want to start at (j-i)+1
				//          Also we want to include the one we're at, Go's indexing for the second parameter is exclusive, so we need to get j+1
				//          This gives the numbers 3,4 and 5, which summed gives 12.
				// 1 2 3  4  5
				// 1 3 6 10 15
				//   |       |
				//   j       i
				return numbers[j-i+1 : j+1]
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

func part1(numbers []int, preambleLength int) int {
	return firstInvalid(numbers, preambleLength)
}

func part2(numbers []int, number int) int {
	contiguous := findContiguous(numbers, number)

	min, max := minMax(contiguous)

	return min + max
}
