package day07

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day7() {
	fmt.Println("Day 7")

	/*
		Each change of 1 step in horizontal position of a single crab costs 1 fuel.
		You could choose any horizontal position to align them all on, but the one that costs the least fuel is horizontal position 2:

		Move from 16 to 2: 14 fuel
		Move from 1 to 2: 1 fuel
		Move from 2 to 2: 0 fuel
		Move from 0 to 2: 2 fuel
		Move from 4 to 2: 2 fuel
		Move from 2 to 2: 0 fuel
		Move from 7 to 2: 5 fuel
		Move from 1 to 2: 1 fuel
		Move from 2 to 2: 0 fuel
		Move from 14 to 2: 12 fuel
		This costs a total of 37 fuel.

		This is the cheapest possible outcome; more expensive outcomes include aligning at position 1 (41 fuel), position 3 (39 fuel),
		or position 10 (71 fuel).

		16,1,2,0,4,2,7,1,2,14
	*/

	s, err := os.ReadFile("day07/input.txt")
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(s), ",")
	part1 := part1(arr)
	fmt.Println("Part 1:", part1)

	part2 := part2(arr)
	fmt.Println("Part 2:", part2)
}

func part1(arr []string) int {
	var input []int
	for _, v := range arr {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		input = append(input, i)
	}
	sort.Ints(input)

	median := CalcMedian(input)
	//fmt.Println("Median:", median)
	totalMoves := 0.0
	for i := 0; i < len(input); i++ {
		totalMoves += math.Abs(float64(input[i] - median))
	}
	//fmt.Println("Total moves:", totalMoves)
	return int(totalMoves)
}

func part2(arr []string) int {
	var input []int
	for _, v := range arr {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		input = append(input, i)
	}
	sort.Ints(input)

	var lowest int = math.MaxInt32
	mean := int(CalcMean(input))
	// the least amount of moves should be right around the mean so try 3 numbers above and below the mean
	attempts := []int{mean - 1, mean, mean + 1}
	for _, v := range attempts {
		//fmt.Println("Mean attempt:", v)
		totalMoves := 0.0
		for i := 0; i < len(input); i++ {
			// how many steps to move to the mean
			steps := math.Abs(float64(input[i] - v))

			// sequence equation is: n(n+1)/2
			// calculate how expensive the move is
			cost := ((steps) * (steps + 1.0)) / 2

			// sum the total moves
			totalMoves += float64(cost)
		}
		// if this is the lowest keep track of it
		if int(totalMoves) < lowest {
			lowest = int(totalMoves)
		}
		//fmt.Println("Total moves:", totalMoves, int(totalMoves), lowest)
	}
	return lowest
}

func CalcMean(input []int) float64 {
	total := 0.0

	for _, v := range input {
		total += float64(v)
	}

	// IMPORTANT: return was rounded!
	return math.Round(total / float64(len(input)))
}

func CalcMedian(input []int) int {
	mNumber := len(input) / 2

	if IsOdd(input) {
		return input[mNumber]
	}

	return (input[mNumber-1] + input[mNumber]) / 2
}

// check if the length of array is even or odd
func IsOdd(input []int) bool {
	return len(input)%2 != 0
}
