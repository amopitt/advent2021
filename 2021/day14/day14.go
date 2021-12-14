package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample.txt
var input string

func main() {
	fmt.Println("Day 14, Hello.")
	lines := strings.Split(input, "\r\n")

	/*
			NNCB

		CH -> B
		HH -> N
		CB -> H
		NH -> C
		HB -> C
		HC -> B
		HN -> C
		NN -> C
		BH -> H
		NC -> B
		NB -> B
		BN -> B
		BB -> N
		BC -> B
		CC -> N
		CN -> C
	*/
	startingInput := ""
	directions := make(map[string]string)
	for c, line := range lines {
		if c == 0 {
			startingInput = line
		} else if line == "" {
			continue
		} else {
			parts := strings.Split(line, " -> ")
			directions[parts[0]] = parts[1]
		}
	}
	fmt.Println("startingInput", startingInput)
	fmt.Println("p1 final answer sir", part2(startingInput, directions, 10))
	fmt.Println("p2 final answer sir", part2(startingInput, directions, 40))

	// leaving the below in so all can see this fail
	// start := 0
	// next := 1
	// newString := ""

	// steps := make(map[int]string)

	// totalSteps := 40
	// for step := 1; step <= totalSteps; step++ {
	// 	index := 0
	// 	for {
	// 		leftString := (string)(startingInput[start])
	// 		rightString := (string)(startingInput[next])
	// 		key := (string)(leftString + rightString)

	// 		//	fmt.Println(key)

	// 		if index == 0 {
	// 			newString = newString[:index] + leftString + newString[index:]
	// 		}

	// 		//fmt.Println("addding left string", index, leftString)

	// 		if repl, ok := directions[key]; ok {
	// 			// insert key between start and next
	// 			index++
	// 			//fmt.Println("addding key string", index, repl)
	// 			newString = newString[:index] + repl + newString[index:]
	// 		}
	// 		index++
	// 		newString = newString[:index] + rightString + newString[index:]
	// 		//fmt.Println("addding right string", index, rightString)

	// 		//	fmt.Println("updated newString:", newString)
	// 		next++
	// 		start++
	// 		if next+1 > len(startingInput) {
	// 			// step completed
	// 			startingInput = newString
	// 			start = 0
	// 			next = 1
	// 			steps[step] = newString
	// 			break
	// 		}
	// 	}
	// 	//	fmt.Println("step", step, "newString:", newString)
	// 	fmt.Println("step completed", step)
	// 	newString = ""
	// }
	// //	fmt.Println("steps", steps)
	// lastKey := steps[totalSteps]
	// occurrences := make(map[string]int)
	// // find least and most common character in lastKey
	// for _, r := range lastKey {
	// 	occurrences[(string)(r)]++
	// }

	// min := 9223372036854775807
	// max := 0

	// for _, v := range occurrences {
	// 	if v < min {
	// 		min = v
	// 	}
	// 	if v > max {
	// 		max = v
	// 	}
	// }
	// fmt.Println("min", min, "max", max, "difference", max-min)
	//fmt.Println("occurrences", occurrences)

	/*
			Template:     NNCB
		After step 1: NCNBCHB
		After step 2: NBCCNBBBCBHCB
		After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
		After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB
	*/
	// fmt.Println(newString)
	// fmt.Println(directions)
	// fmt.Println(startingInput)
}

func part2(startingInput string, directions map[string]string, steps int) int {
	counter := make(map[string]int)

	// loop over the input
	for i := 0; i < len(startingInput)-1; i++ {
		counter[startingInput[i:i+2]]++
	}

	fmt.Println("counter", counter)

	// loop for each step
	for i := 0; i < steps; i++ {
		newCounter := make(map[string]int)
		for k, v := range counter {
			if ch, ok := directions[k]; ok {
				newCounter[(string)(k[0])+ch] += v
				newCounter[ch+(string)(k[1])] += v
			}
		}
		counter = newCounter
	}

	/**
	This polymer grows quickly. After step 5, it has length 97; After step 10, it has length 3073.
	After step 10, B occurs 1749 times, C occurs 298 times, H occurs 161 times, and N occurs 865 times; taking the quantity of the most common element (B, 1749) and subtracting the quantity of the least common element (H, 161) produces 1749 - 161 = 1588.
	*/
	answer := make(map[string]int)

	for k, v := range counter {
		// loop over k
		for i := 0; i < len(k); i++ {
			answer[(string)(k[i])] += v
		}
	}

	// beginning and end wont be duplicated so add 1 to them
	answer[(string)(startingInput[0])]++
	answer[(string)(startingInput[len(startingInput)-1])]++

	// divide by 2 to remove duplicate counts
	for k := range answer {
		answer[k] = answer[k] / 2
	}
	fmt.Println("answer", answer)

	min := 9223372036854775807
	max := 0

	for _, v := range answer {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	fmt.Println("min", min, "max", max, "difference", max-min)

	return max - min

}
