package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ReadInput reads file's lines into array of strings
func ReadInput() []string {
	var input []string

	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

// ReadInputLinesAsNumbers reads file's lines into an array of numbers
func ReadInputLinesAsNumbers() []int {
	var input []int

	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var n int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &n)
		if err != nil {
			panic(err)
		}

		input = append(input, n)
	}

	return input
}

// ReadInputAsNumbers reads file as line of numbers
func ReadInputAsNumbers() []int {
	var input []int

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	text := strings.Trim(string(content), "\r\n")

	tokens := strings.Split(text, ", ")

	for _, i := range tokens {
		val, _ := strconv.Atoi(i)
		input = append(input, val)
	}

	return input
}
