package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordMatch struct {
	min      int
	max      int
	key      string
	password string
}

func main() {
	/*
			1-3 a: abcde
		1-3 b: cdefg
		2-9 c: ccccccccc
	*/
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var p []PasswordMatch
	//var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		comp := strings.Split(value, ":")     // 1-3 a, abcde
		policy := strings.Split(comp[0], " ") // 1-3, a

		minStr := strings.Split(policy[0], "-")[0]
		maxStr := strings.Split(policy[0], "-")[1]
		min, _ := strconv.Atoi(minStr)
		max, _ := strconv.Atoi(maxStr)

		var pw PasswordMatch = PasswordMatch{
			min:      min,
			max:      max,
			key:      policy[1],
			password: strings.TrimSpace(comp[1]),
		}

		p = append(p, pw)
	}
	fmt.Println(p)
	//fmt.Println(part1(p))
	fmt.Println(part2(p))
}

func part1(p []PasswordMatch) int {
	var validPasswords int
	for _, v := range p {

		// count the number of times the key appears in the password
		count := strings.Count(v.password, v.key)

		// if the count is less than the min or greater than the max, then the password doesn't match
		if count >= v.min && count <= v.max {
			validPasswords++
		}
	}
	return validPasswords
}

func part2(p []PasswordMatch) int {
	var validPasswords int
	for _, v := range p {
		matches := 0
		vals := strings.Split(v.password, "")
		if vals[v.min-1] == v.key {
			matches++
		}
		if vals[v.max-1] == v.key {
			matches++
		}

		if matches == 1 {
			validPasswords++
		}
	}
	return validPasswords
}
