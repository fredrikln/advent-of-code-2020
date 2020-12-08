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
	Opcode  string
	Operand int
}

// VM is a virtual machine
type VM struct {
	Program        []Instruction
	ProgramCounter int
	Accumulator    int
}

// Load parses a string slice into instructions and returns a virtual machine
func (vm *VM) Load(program []string) {
	for _, line := range program {
		split := strings.Split(line, " ")
		opcode := split[0]
		operand, _ := strconv.Atoi(split[1])

		ins := Instruction{opcode, operand}

		vm.Program = append(vm.Program, ins)
	}
}

// Step step the VM forward 1 pc
func (vm *VM) Step() {
	ins := &vm.Program[vm.ProgramCounter]

	switch ins.Opcode {
	case "nop":
		vm.ProgramCounter++
	case "jmp":
		vm.ProgramCounter += ins.Operand
	case "acc":
		vm.ProgramCounter++
		vm.Accumulator += ins.Operand
	}
}

// PrintCurrentState prints current operation
func (vm *VM) PrintCurrentState() {
	pc := vm.ProgramCounter
	instruction := vm.Program[pc]

	fmt.Printf("pc: %5d :: instruction: %s % 5d :: accumulator: % 5d\n", pc, instruction.Opcode, instruction.Operand, vm.Accumulator)
}

// RunVM runs the program and returns the Accumulator and ProgramCounter on completion
func runVM(vm VM, printState bool) (int, int) {
	timesRun := make(map[int]int)

	for {
		// Exit if we try to run instruction outside of program
		if vm.ProgramCounter >= len(vm.Program) {
			break
		}

		// Get pointer to current instruction and increase the amount of times we've run it
		timesRun[vm.ProgramCounter]++

		// If instruction ran more than once we exit
		if timesRun[vm.ProgramCounter] > 1 {
			break
		}

		if printState {
			vm.PrintCurrentState()
		}

		// Implement instructions
		vm.Step()
	}

	// On exit return current accumulator and program counter
	return vm.Accumulator, vm.ProgramCounter
}

func part1(program []string) int {
	vm := VM{}
	vm.Load(program)

	acc, _ := runVM(vm, false)

	return acc
}

func part2(program []string) int {
	var numProgramsTested int
	for i := 0; i < len(program); i++ {
		if program[i][:3] != "nop" && program[i][:3] != "jmp" {
			continue
		}

		numProgramsTested++

		vm := VM{}
		vm.Load(program)

		ins := &vm.Program[i]

		if ins.Opcode == "nop" {
			ins.Opcode = "jmp"
		} else if ins.Opcode == "jmp" {
			ins.Opcode = "nop"
		}

		acc, ptr := runVM(vm, false)

		if ptr == len(program) {
			return acc
		}
	}

	return 0
}
