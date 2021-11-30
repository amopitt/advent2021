package day01

import (
	"advent2021/util"
	"fmt"
)

func Day1() {
	// read input and get it into an array of ints
	input, err := util.ReadFileInts("day01/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	a, b, c := findThreeThatSum2020(input)
	fmt.Println(a * b * c)

	// convert input to a hash map
	inputMap := make(map[int]int)
	for _, value := range input {
		inputMap[value] = value
	}

	// iterate through input
	var sum int
	for _, value := range input {
		missingValue := 2020 - value
		if _, ok := inputMap[missingValue]; ok {
			sum = value * missingValue
			break
		}
	}
	println(sum)
}

// findThreeThatSum2020 finds the three numbers that sum to 2020 - brute force approach O(n^3)
func findThreeThatSum2020(input []int) (int, int, int) {
	for a, value := range input {
		for b, value2 := range input {
			for c, value3 := range input {
				if a != b && b != c && value+value2+value3 == 2020 {
					return value, value2, value3
				}
			}
		}
	}
	return 0, 0, 0
}
