package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadInput reads file's lines into array of strings
func ReadInput() []string {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	contents := string(bytes)

	lines := strings.Split(contents, "\r\n")
	lines = lines[:len(lines)-1]

	return lines
}

// ReadInputLinesAsNumbers reads file's lines into an array of numbers
func ReadInputLinesAsNumbers() []int {
	var input []int

	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	contents := strings.Trim(string(bytes), "\r\n")

	lines := strings.Split(contents, "\r\n")

	for _, s := range lines {
		n, err := strconv.Atoi(s)
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

	tokens := strings.Split(text, ",")

	for _, i := range tokens {
		val, _ := strconv.Atoi(i)
		input = append(input, val)
	}

	return input
}
