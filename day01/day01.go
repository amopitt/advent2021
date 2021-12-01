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
	part1(input)
	part2(input)

}

/*
199  A
200  A B
208  A B C
210    B C D
200  E   C D
207  E F   D
240  E F G
269    F G H
260      G H
263        H
*/
func part2(input []int) {
	start := 0
	compare := 1
	// compare 0 - 2 to 1 - 3
	// compare 1 - 3 to 2 - 4
	// compare 2 - 4 to 3 - 5
	numIncreased := 0
	for {
		// at the end here
		if compare+2 >= len(input) {
			break
		}

		if getSum(input, compare, 3) > getSum(input, start, 3) {
			numIncreased++
		}
		start++
		compare++
	}
	fmt.Println(numIncreased)
}
func getSum(input []int, start int, num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		sum += input[start+i]
	}
	return sum
}

func part1(input []int) {
	start := 0
	compare := 1
	numIncreased := 0
	for {
		if compare >= len(input) {
			break
		}

		if input[compare] > input[start] {
			numIncreased++
		}

		start++
		compare++
	}
	fmt.Println(numIncreased)
}
