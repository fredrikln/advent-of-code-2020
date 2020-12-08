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

// Instruction is a single instruction
type Instruction struct {
	Op       string
	Value    int
	TimesRun int
}

// VirtualMachine is a virtual machine
type VirtualMachine struct {
	Program        []Instruction
	ProgramCounter int
	Accumulator    int
}

// NewVirtualMachine parses a string slice into instructions and returns a virtual machine
func NewVirtualMachine(program []string) VirtualMachine {
	vm := VirtualMachine{}

	for _, line := range program {
		split := strings.Split(line, " ")
		val, _ := strconv.Atoi(split[1])

		ins := Instruction{split[0], val, 0}

		vm.Program = append(vm.Program, ins)
	}

	return vm
}

// Run runs the program and returns the Accumulator and ProgramCounter on completion
func (vm *VirtualMachine) Run() (int, int) {
	for {
		// Exit if we try to run instruction outside of program
		if vm.ProgramCounter >= len(vm.Program) {
			break
		}

		// Get pointer to current instruction and increase the amount of times we've run it
		ins := &vm.Program[vm.ProgramCounter]
		ins.TimesRun++

		// If ran more than one we exit
		if ins.TimesRun > 1 {
			break
		}

		// Implement instructions
		switch ins.Op {
		case "nop":
			vm.ProgramCounter++
		case "jmp":
			vm.ProgramCounter += ins.Value
		case "acc":
			vm.ProgramCounter++
			vm.Accumulator += ins.Value
		}
	}

	// On exit return current accumulator and program counter
	return vm.Accumulator, vm.ProgramCounter
}

func part1(program []string) int {
	vm := NewVirtualMachine(program)

	acc, _ := vm.Run()

	return acc
}

func part2(program []string) int {
	for i := 0; i < len(program); i++ {
		vm := NewVirtualMachine(program)

		ins := &vm.Program[i]

		if ins.Op == "nop" {
			ins.Op = "jmp"
		} else if ins.Op == "jmp" {
			ins.Op = "nop"
		}

		acc, ptr := vm.Run()

		if ptr == len(program) {
			return acc
		}
	}

	return 0
}
