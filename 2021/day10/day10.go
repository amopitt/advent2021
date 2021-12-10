package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Day 10, Hello.")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	invalidKeys := make(map[string]int)
	completedScores := make([]int, 0)

	for scanner.Scan() {
		value := scanner.Text()
		keys := strings.Split(value, "")
		leftStack := make([]string, 0)
		isCorruptLine := false
		for _, key := range keys {
			if isOpeningBracket(key) {
				leftStack = append(leftStack, key)
			} else {
				if len(leftStack) == 0 {
					fmt.Println("No opening bracket found", key)
					os.Exit(1)
				}

				if key == ")" && leftStack[len(leftStack)-1] == "(" {
					leftStack = leftStack[:len(leftStack)-1]
				} else if key == "]" && leftStack[len(leftStack)-1] == "[" {
					leftStack = leftStack[:len(leftStack)-1]
				} else if key == "}" && leftStack[len(leftStack)-1] == "{" {
					leftStack = leftStack[:len(leftStack)-1]
				} else if key == ">" && leftStack[len(leftStack)-1] == "<" {
					leftStack = leftStack[:len(leftStack)-1]
				} else {
					fmt.Println("No matching bracket found", key)
					invalidKeys[key] = invalidKeys[key] + 1
					isCorruptLine = true
					break
				}
			}
		}
		if !isCorruptLine && len(leftStack) > 0 {
			fmt.Println(leftStack, "needs completed")
			score := 0
			// loop leftStack desecending
			for i := len(leftStack) - 1; i >= 0; i-- {
				key := leftStack[i]
				keyValue := 1
				if key == "[" {
					keyValue = 2
				} else if key == "{" {
					keyValue = 3
				} else if key == "<" {
					keyValue = 4
				}
				score = score*5 + keyValue
			}

			completedScores = append(completedScores, score)
		}
	} // end for scanner.Scan()
	// 			): 1 point.
	// ]: 2 points.
	// }: 3 points.
	// >: 4 points.
	// So, the last completion string above - ])}> - would be scored as follows:

	// Start with a total score of 0.
	// Multiply the total score by 5 to get 0, then add the value of ] (2) to get a new total score of 2.
	// Multiply the total score by 5 to get 10, then add the value of ) (1) to get a new total score of 11.
	// Multiply the total score by 5 to get 55, then add the value of } (3) to get a new total score of 58.
	// Multiply the total score by 5 to get 290, then add the value of > (4) to get a new total score of 294.
	sort.Ints(completedScores)
	fmt.Println("Completed scores:", completedScores)
	fmt.Println("Middle score:", completedScores[len(completedScores)/2])

	fmt.Println("Invalid keys:", invalidKeys)
	sum := 0
	for key, value := range invalidKeys {

		fmt.Println("Key:", key, "Value:", value)
		multiplier := 3 // ) == 3
		if key == "]" {
			multiplier = 57
		} else if key == "}" {
			multiplier = 1197
		} else if key == ">" {
			multiplier = 25137
		}

		sum = sum + value*multiplier
	}
	fmt.Println("Sum:", sum)
}

/*
If a chunk opens with (, it must close with ).
If a chunk opens with [, it must close with ].
If a chunk opens with {, it must close with }.
If a chunk opens with <, it must close with >.
*/

func isOpeningBracket(key string) bool {
	return key == "(" || key == "[" || key == "{" || key == "<"
}
