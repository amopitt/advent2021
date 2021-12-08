package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
	  0:      1:      2:      3:      4:
	aaaa    ....    aaaa    aaaa    ....
	b    c  .    c  .    c  .    c  b    c
	b    c  .    c  .    c  .    c  b    c
	....    ....    dddd    dddd    dddd
	e    f  .    f  e    .  .    f  .    f
	e    f  .    f  e    .  .    f  .    f
	gggg    ....    gggg    gggg    ....

	5:      6:      7:      8:      9:
	aaaa    aaaa    aaaa    aaaa    aaaa
	b    .  b    .  .    c  b    c  b    c
	b    .  b    .  .    c  b    c  b    c
	dddd    dddd    ....    dddd    dddd
	.    f  e    f  .    f  e    f  .    f
	.    f  e    f  .    f  e    f  .    f
	gggg    gggg    ....    gggg    gggg
*/
func main() {
	fmt.Println("Day 8")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numberMap := make(map[int]string)
	numberMap[0] = "abcefg"  // len 6, matches 4 = 3, matches 1 = 2
	numberMap[1] = "cf"      // len 2
	numberMap[2] = "acdeg"   // len 5, matches 4 = 2, matches 1 = 1
	numberMap[3] = "acdfg"   // len 5, matches 4 = 3, matches 1 = 2
	numberMap[4] = "bcdf"    // len 4
	numberMap[5] = "abdfg"   // len 5, matches 4 = 3, matches 1 = 1
	numberMap[6] = "abdefg"  // len 6, matches 4 = 3, matches 1 = 1
	numberMap[7] = "acf"     // len 3
	numberMap[8] = "abcdefg" // len 7
	numberMap[9] = "abcdfg"  // len 6, matches 4 = 4, matches 1 = 2

	sum := 0
	for scanner.Scan() {
		value := scanner.Text()
		kv := strings.Split(value, " | ")
		input := kv[0]
		output := kv[1]
		inputEntries := strings.Split(input, " ")
		finalEntries := strings.Split(output, " ")

		// loop over inputEntries
		for _, v := range inputEntries {
			numberMap[len(v)] = v
		}
		d := ""
		for _, v := range finalEntries {
			len := len(v)
			match4 := getMatchCount(v, numberMap[4])
			match1 := getMatchCount(v, numberMap[2])
			switch {
			// unique values where length directly maps to a number
			case len == 2:
				d += "1"
			case len == 3:
				d += "7"
			case len == 4:
				d += "4"
			case len == 7:
				d += "8"
			// for these two cases, check the number of chars that match the 4th and 1st numbers
			case len == 5:
				if match4 == 2 {
					d += "2"
				} else {
					if match1 == 1 {
						d += "5"
					} else {
						d += "3"
					}
				}
			case len == 6:
				if match4 == 4 {
					d += "9"
				} else {
					if match1 == 1 {
						d += "6"
					} else {
						d += "0"
					}
				}

			}
		}

		fmt.Println("iteration sum", d)
		intValue, _ := strconv.Atoi(d)
		sum += intValue
	} // end for scanner.Scan()
	fmt.Println("total", sum)
}

func getMatchCount(s string, s2 string) int {
	arrS1 := strings.Split(s, "")

	// loop over arrS1
	count := 0
	for _, v := range arrS1 {
		if strings.Contains(s2, v) {
			count++
		}
	}
	return count
}

// ended up not using these after dealing with part 2 but leaving them for utils later maybe

func getUniqueChar(s1, s2 string) string {
	arrS1 := strings.Split(s1, "")
	arrS2 := strings.Split(s2, "")
	var foundChar string
	for _, s := range arrS2 {
		if !contains(arrS1, s) {
			foundChar += s
		}
	}
	return foundChar
}

func contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func sortArrayByLength(arr []string) {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
}

func isUnique(s string) bool {
	strLength := len(s)
	if strLength == 2 || strLength == 4 || strLength == 3 || strLength == 7 {
		return true
	}
	return false
}
