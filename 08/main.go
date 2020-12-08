package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadInput("input.txt")

	fmt.Println("Part 1")
	result := part1(input)
	fmt.Println("Accumulator:", result)

	fmt.Println("Part 2")
	result2 := part2(input)
	fmt.Println("Accumulator:", result2)
}

type instruction struct {
	instruction string
	value       int
	timesRun    int
}

func parseInstructions(data []string) []*instruction {
	instructions := make([]*instruction, 0)
	for _, line := range data {
		split := strings.Split(line, " ")
		val, _ := strconv.Atoi(split[1])

		ins := instruction{split[0], val, 0}

		instructions = append(instructions, &ins)
	}
	return instructions
}

func runProgram(instructions []*instruction) (int, int) {
	var acc int
	var ptr int

	for {
		if ptr == len(instructions) {
			break
		}

		ins := instructions[ptr]

		ins.timesRun++

		if ins.timesRun > 1 {
			break
		}

		switch ins.instruction {
		case "nop":
			ptr++
		case "jmp":
			ptr += ins.value
		case "acc":
			ptr++
			acc += ins.value
		}
	}

	return acc, ptr
}

func part1(program []string) int {
	instructions := parseInstructions(program)

	acc, _ := runProgram(instructions)

	return acc
}

func part2(program []string) int {
	for i := 0; i < len(program); i++ {
		instructions := parseInstructions(program)

		ins := instructions[i]

		if ins.instruction == "nop" {
			ins.instruction = "jmp"
		} else if ins.instruction == "jmp" {
			ins.instruction = "nop"
		}

		acc, ptr := runProgram(instructions)

		if ptr == len(program) {
			return acc
		}
	}

	return 0
}
