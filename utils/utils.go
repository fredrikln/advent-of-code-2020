package utils

import (
	"bufio"
	"os"
)

// ReadInput reads input file lines into an array
func ReadInput() []string {

	var input []string

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
