package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(file), "\r\n")
	input := map[int]int{}
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		input[i] = i
	}
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

/*
1721
979
366
299
675
1456

Find the two entries that sum to 2020; what do you get if you multiply them together?
*/
func part1(input map[int]int) int {
	for k := range input {
		if input[2020-k] != 0 {
			fmt.Println("values", k, 2020-k)
			return k * input[2020-k]
		}
	}
	return 0
}

func part2(input map[int]int) int {
	// find 3 entries that sum to 2020
	for k := range input {
		for k2 := range input {
			if k2 == k {
				continue
			}
			if input[2020-k-k2] != 0 {
				fmt.Println("values", k, k2, 2020-k-k2)
				return k * k2 * input[2020-k-k2]
			}
		}
	}
	return 0
}
