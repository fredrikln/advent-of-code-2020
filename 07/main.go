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

	bags := parseBags(input)

	fmt.Println("Part 1")
	result := part1(bags, "shiny gold")
	fmt.Println("Count:", result)

	fmt.Println("Part 2")
	result2 := part2(bags, "shiny gold")
	fmt.Println("Count:", result2)
}

type content struct {
	count int
	bag   *bag
}
type bag struct {
	contains map[string]*content
}

func parseBags(input []string) map[string]*bag {
	bags := make(map[string]*bag)

	for _, line := range input {
		firstBags := strings.Index(line, " bags")

		color := line[:firstBags]
		re := regexp.MustCompile("([0-9]+) ([a-z ]+) bag")

		contains := re.FindAllStringSubmatch(line, -1)

		curBag, ok := bags[color]
		if !ok {
			bags[color] = &bag{make(map[string]*content, 0)}
			curBag = bags[color]
		}

		for _, bagmatch := range contains {
			theBag, ok := bags[bagmatch[2]]
			if !ok {
				bags[bagmatch[2]] = &bag{make(map[string]*content, 0)}
				theBag = bags[bagmatch[2]]
			}

			num, err := strconv.Atoi(bagmatch[1])
			if err != nil {
				panic(err)
			}

			curBag.contains[bagmatch[2]] = &content{num, theBag}
		}
	}

	return bags
}

func contains(curBag *bag, color string) bool {
	_, ok := curBag.contains[color]
	if ok {
		return true
	}

	for _, content := range curBag.contains {
		if contains(content.bag, color) {
			return true
		}
	}

	return false
}

func part1(bags map[string]*bag, color string) int {
	var count int
	for _, bag := range bags {
		if contains(bag, color) {
			count++
		}
	}
	return count
}

func countSubBags(bag *bag) int {
	count := 0
	for _, contains := range bag.contains {
		count += contains.count + contains.count*countSubBags(contains.bag)
	}

	return count
}

func part2(bags map[string]*bag, color string) int {
	bag := bags[color]

	return countSubBags(bag)
}
