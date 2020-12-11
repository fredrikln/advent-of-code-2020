package main

import (
	"fmt"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func main() {
	filename := "input.txt"
	input := utils.ReadInput(filename)

	var numChanged int
	for {
		input, numChanged = step(input)
		if numChanged == 0 {
			break
		}
	}

	fmt.Println("Part 1:", countOccupied(input))

	input = utils.ReadInput(filename)

	numChanged = 0
	for {
		input, numChanged = step2(input)
		if numChanged == 0 {
			break
		}
	}

	fmt.Println("Part 2:", countOccupied(input))
}

func step(theMap []string) ([]string, int) {
	newMap := make([]string, len(theMap))

	var numChanged int

	for y, line := range theMap {
		newLine := make([]rune, len(line))

		for x, char := range line {
			if char == '.' {
				newLine[x] = '.'
				continue
			}

			var countOccupied int

			for i := -1; i <= 1; i++ {
				if y+i < 0 || y+i >= len(theMap) {
					continue
				}

				for j := -1; j <= 1; j++ {
					if x+j < 0 || x+j >= len(line) {
						continue
					}

					if i == 0 && j == 0 {
						continue
					}

					if theMap[y+i][x+j] == '#' {
						countOccupied++
					}
				}
			}

			if char == 'L' && countOccupied == 0 {
				newLine[x] = '#'
				numChanged++
			} else if char == '#' && countOccupied >= 4 {
				newLine[x] = 'L'
				numChanged++
			} else {
				newLine[x] = char
			}
		}

		newMap[y] = string(newLine)
	}

	return newMap, numChanged
}

func step2(theMap []string) ([]string, int) {
	newMap := make([]string, len(theMap))

	var numChanged int

	for y, line := range theMap {
		newLine := make([]rune, len(line))

		for x, char := range line {
			if char == '.' {
				newLine[x] = '.'
				continue
			}

			countOccupied := 0

			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}

					if canSeeOccupied(theMap, x, y, j, i) {
						countOccupied++
					}
				}
			}

			if char == 'L' && countOccupied == 0 {
				newLine[x] = '#'
				numChanged++
			} else if char == '#' && countOccupied >= 5 {
				newLine[x] = 'L'
				numChanged++
			} else {
				newLine[x] = char
			}
		}

		newMap[y] = string(newLine)
	}

	return newMap, numChanged
}

func canSeeOccupied(theMap []string, x, y, xDir, yDir int) bool {
	nextX := x
	nextY := y

	for {
		nextX += xDir
		nextY += yDir

		if nextY < 0 || nextY >= len(theMap) {
			return false
		}
		if nextX < 0 || nextX >= len(theMap[0]) {
			return false
		}

		if theMap[nextY][nextX] == 'L' {
			return false
		}

		if theMap[nextY][nextX] == '#' {
			return true
		}
	}
}

func printMap(theMap []string) {
	for _, line := range theMap {
		fmt.Println(line)
	}
}

func countOccupied(theMap []string) (numChanged int) {
	for _, line := range theMap {
		for _, char := range line {
			if char == '#' {
				numChanged++
			}
		}
	}

	return
}
