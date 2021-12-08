package day02

import (
	"advent2021/util"
	"fmt"
	"strconv"
	"strings"
)

func Day2() {
	// read input and get it into an array of ints
	input, err := util.ReadFile("day02/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	part1 := part1(input)
	part2 := part2(input)
	fmt.Println("")
	fmt.Printf("part1 = %d \n", part1)
	fmt.Printf("part2 = %d \n", part2)
}

// Your horizontal position and depth both start at 0
// After following these instructions, you would have a horizontal position of 15 and a depth of 10.
// (Multiplying these together produces 150.)
func part1(input []string) int {
	horizontal := 0
	depth := 0
	// loop over the input
	for _, line := range input {
		items := strings.Split(line, " ")
		instruction := items[0]

		count, err := strconv.Atoi(items[1])
		if err != nil {
			panic(err)
		}

		//
		switch instruction {
		case "forward":
			horizontal += count
		case "down":
			depth += count
		case "up":
			depth -= count
		}
	}
	sum := horizontal * depth
	//fmt.Printf("%d - %d = %d", horizontal, depth, sum)
	return sum
}

// down X increases your aim by X units.
// up X decreases your aim by X units.
// forward X does two things:
// It increases your horizontal position by X units.
// It increases your depth by your aim multiplied by X.
func part2(input []string) int {
	horizontal := 0
	depth := 0
	aim := 0
	// loop over the input
	for _, line := range input {
		items := strings.Split(line, " ")
		instruction := items[0]

		count, err := strconv.Atoi(items[1])
		if err != nil {
			panic(err)
		}

		switch instruction {
		case "forward":
			horizontal += count
			depth += (aim * count)
		case "down":
			aim += count
		case "up":
			aim -= count
		}
		// fmt.Printf("horiz = %d, depth= %d, aim = %d", horizontal, depth, aim)
		// fmt.Println()
	}
	sum := horizontal * depth
	//fmt.Printf("%d - %d = %d", horizontal, depth, sum)
	return sum
}
