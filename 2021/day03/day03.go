package day03

import (
	"advent2021/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day3() {
	// read input of bit strings
	input, err := util.ReadFile("day03/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// create a new list
	result := part1(input)
	part2a := part2(input, 0, true)
	part2b := part2(input, 0, false)
	fmt.Println("")
	fmt.Printf("part1 = %d \n", result)
	fmt.Printf("part2 = %d \n", part2a)
	fmt.Printf("part2 = %d \n", part2b)

	max := "100111110011"
	min := "001011100001"
	intMax, err := strconv.ParseInt(max, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	intMin, err := strconv.ParseInt(min, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(intMax)
	fmt.Println(intMin)
	fmt.Println(intMax * intMin)
}

type Common struct {
	Zero       int
	One        int
	ZeroValues []string
	OneValues  []string
}

func part1(input []string) int64 {

	dict := make(map[int]Common)
	// find the most common bit in each slot

	// loop over the input
	for _, line := range input {
		// split line into bits
		bits := strings.Split(line, "")

		// loop over the bits
		for i := 0; i < len(bits); i++ {
			// check if the key exists
			if _, ok := dict[i]; !ok {
				// if not create a new entry
				dict[i] = Common{0, 0, []string{}, []string{}}
			}
			if entry, ok := dict[i]; ok {
				// check if the bit is 0 or 1
				if bits[i] == "0" {
					// if 0 increment the zero count
					entry.Zero = entry.Zero + 1
				} else {
					// if 1 increment the one count
					entry.One = entry.One + 1
				}
				dict[i] = entry
			}
		}
	}
	keys := make([]int, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// loop over the dict
	var max string
	var min string
	for _, key := range keys {
		if dict[key].Zero > dict[key].One {
			max += "0"
			min += "1"
		} else {
			max += "1"
			min += "0"
		}
	}
	//var min string

	fmt.Println(max)
	fmt.Println(min)
	intMax, err := strconv.ParseInt(max, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	intMin, err := strconv.ParseInt(min, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	fmt.Println(intMax)
	fmt.Println(intMin)
	return intMax * intMin
}

/**
Start with all 12 numbers and consider only the first bit of each number. There are more 1 bits (7) than 0 bits (5),
so keep only the 7 numbers with a 1 in the first position: 11110, 10110, 10111, 10101, 11100, 10000, and 11001.
Then, consider the second bit of the 7 remaining numbers: there are more 0 bits (4) than 1 bits (3), so keep only the 4 numbers with a 0 in the second position: 10110, 10111, 10101, and 10000.
In the third position, three of the four numbers have a 1, so keep those three: 10110, 10111, and 10101.
In the fourth position, two of the three numbers have a 1, so keep those two: 10110 and 10111.
In the fifth position, there are an equal number of 0 bits and 1 bits (one each). So, to find the oxygen generator rating, keep the number with a 1 in that position: 10111.
As there is only one number left, stop; the oxygen generator rating is 10111, or 23 in decimal.
*/
func part2(input []string, i int, findMax bool) int {
	dict := make(map[int]Common)
	// find the most common bit in each slot

	// loop over the input
	for _, line := range input {
		// split line into bits
		bits := strings.Split(line, "")

		// check if the key exists
		if _, ok := dict[i]; !ok {
			// if not create a new entry
			dict[i] = Common{0, 0, []string{}, []string{}}
		}
		if entry, ok := dict[i]; ok {
			// check if the bit is 0 or 1
			if bits[i] == "0" {
				// if 0 increment the zero count
				entry.Zero = entry.Zero + 1
				entry.ZeroValues = append(entry.ZeroValues, line)
			} else {
				// if 1 increment the one count
				entry.One = entry.One + 1
				entry.OneValues = append(entry.OneValues, line)
			}
			dict[i] = entry
		}
	}
	for {
		// yea........ rush job
		if findMax {
			if dict[i].Zero == dict[i].One || dict[i].One > dict[i].Zero {
				// use the values in
				if len(dict[i].OneValues) == 1 {
					fmt.Println(dict[i].OneValues)
					break
				} else {
					return part2(dict[i].OneValues, i+1, findMax)
				}
			} else {
				if len(dict[i].ZeroValues) == 1 {
					fmt.Println(dict[i].ZeroValues)
					break
				} else {
					return part2(dict[i].ZeroValues, i+1, findMax)
				}
			}
		} else {
			if dict[i].Zero == dict[i].One || dict[i].Zero < dict[i].One {
				// use the values in
				if len(dict[i].ZeroValues) == 1 {
					fmt.Println(dict[i].ZeroValues)
					break
				} else {
					return part2(dict[i].ZeroValues, i+1, findMax)
				}
			} else {
				if len(dict[i].OneValues) == 1 {
					fmt.Println(dict[i].OneValues)
					break
				} else {
					return part2(dict[i].OneValues, i+1, findMax)
				}
			}
		}
	}
	return 0
}
