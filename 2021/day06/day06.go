package day06

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day6() {
	fileName := "day06/sample.txt"
	input, err := getFishMap(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(input)

	// Part 1
	//result := part1(input, 256)
	//fmt.Println("Part 1: ", result)

	// Part 2
	result := part2(input, 256)
	fmt.Println("Part 2: ", result)
}

func getFishMap(fileName string) ([]int, error) {
	s, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	fish := strings.Split(string(s), ",")
	// day 8 + 1
	fishMap := make([]int, 9)
	for i := 0; i < len(fish); i++ {
		n, err := strconv.Atoi(fish[i])
		if err != nil {
			return nil, err
		}
		fishMap[n]++
	}
	return fishMap, nil
}

// original brute force bad way
func part1(fishes []int, days int) int {
	for i := 0; i < days; i++ {

		// each day, decrement the fish by
		addFish := make([]int, 0)
		// if fish == 0, set it to 8 and then append a fish at the end
		for j := 0; j < len(fishes); j++ {
			if fishes[j] == 0 {
				fishes[j] = 6
				addFish = append(addFish, 8)
			} else {
				fishes[j]--
			}
		}
		fishes = append(fishes, addFish...)
	}
	return len(fishes)
}

/*
Initial state: 3,4,3,1,2
After  1 day:  2,3,2,0,1
After  2 days: 1,2,1,6,0,8
After  3 days: 0,1,0,6,0,8,8
After  4 days: 6,0,6,4,5,6,7,8,8
3, 4, 3, 1, 2........

after 13 days we would add 5 fish

all new fish have a growth rate of 8

*/

// Day 6 Part 2
// https://adventofcode.com/2021/day/6
func part2(fishes []int, days int) int {
	// loop through the days
	for i := 1; i <= days; i++ {
		// anything at 0 index will cause new fish to be added
		newFishes := fishes[0]
		// reset the value
		fishes[0] = 0

		// for the rest of the array, decrement the fish until population by moving them down a slot
		for j := 1; j < len(fishes); j++ {
			fishes[j-1] = fishes[j]
		}

		// newFishes spawn fish that have a timer of 6 but also duplicate themselves at the 8 timer
		fishes[6], fishes[8] = fishes[6]+newFishes, newFishes
	}

	sum := 0
	for _, c := range fishes {
		sum += c
	}
	return sum
}
