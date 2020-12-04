package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	passports := parseInput(input)

	fmt.Println("Part 1")
	countValid := part1(passports)
	fmt.Println(countValid)

	fmt.Println("Part 2")
	countValid2 := part2(passports)
	fmt.Println(countValid2)
}

func parseInput(data []string) []map[string]string {
	passports := make([]map[string]string, 0)

	curPassport := make(map[string]string)
	for _, line := range data {
		if len(line) == 0 {
			passports = append(passports, curPassport)
			curPassport = make(map[string]string)
			continue
		}

		fields := strings.Split(line, " ")

		for _, field := range fields {
			data := strings.Split(field, ":")

			curPassport[data[0]] = data[1]
		}
	}
	passports = append(passports, curPassport)

	return passports
}

func part1(passports []map[string]string) int {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	numValid := 0
	for _, passport := range passports {
		valid := true
		for _, requiredField := range requiredFields {
			_, ok := passport[requiredField]
			if !ok {
				valid = false
			}
		}
		if valid {
			numValid++
		}
	}

	return numValid
}

func checkYr(data string, min, max int) bool {
	year, err := strconv.Atoi(data)
	if err != nil {
		panic(err)
	}

	if year < min || year > max {
		return false
	}

	return true
}

func checkHeight(data string) bool {
	if len(data) < 3 {
		return false
	}

	height, err := strconv.Atoi(data[:len(data)-2])

	if err != nil {
		panic(err)
	}
	unit := data[len(data)-2:]

	if unit == "cm" && (height < 150 || height > 193) {
		return false
	} else if unit == "in" && (height < 59 || height > 76) {
		return false
	} else if unit != "in" && unit != "cm" {
		return false
	}

	return true
}

func checkHcl(data string) bool {
	re := regexp.MustCompile(`^#[0-9a-z]{6}$`)
	match := re.MatchString(data)

	if !match {
		return false
	}

	return true
}

func checkEcl(data string) bool {
	re := regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)`)
	match := re.MatchString(data)

	if !match {
		return false
	}

	return true
}

func checkPid(data string) bool {
	re := regexp.MustCompile(`^[0-9]{9}$`)
	match := re.MatchString(data)

	if !match {
		return false
	}

	return true
}

func part2(passports []map[string]string) int {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	numValid := 0

	for _, passport := range passports {
		valid := true

		for _, requiredField := range requiredFields {
			data, ok := passport[requiredField]
			if !ok {
				valid = false
				continue
			}

			switch requiredField {
			case "byr":
				if !checkYr(data, 1920, 2002) {
					valid = false
				}
			case "iyr":
				if !checkYr(data, 2010, 2020) {
					valid = false
				}
			case "eyr":
				if !checkYr(data, 2020, 2030) {
					valid = false
				}
			case "hgt":
				if !checkHeight(data) {
					valid = false
				}
			case "hcl":
				if !checkHcl(data) {
					valid = false
				}
			case "ecl":
				if !checkEcl(data) {
					valid = false
				}
			case "pid":
				if !checkPid(data) {
					valid = false
				}
			}
		}
		if valid {
			numValid++
		}
	}

	return numValid
}
