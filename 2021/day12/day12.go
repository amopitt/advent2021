package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 12, Hello.")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// e.g. A -> start, b, e
	paths := make(map[string][]string)
	for scanner.Scan() {
		value := scanner.Text()
		path := strings.Split(value, "-")
		// forward and backward
		paths[path[0]] = append(paths[path[0]], path[1])
		paths[path[1]] = append(paths[path[1]], path[0])
	}

	part1 := countPaths(paths, "start", map[string]bool{"start": true})
	fmt.Println("Part 1 paths:", part1)

	path2 := countPathsPart2FancyRules(paths, "start", map[string]int{"start": 1}, false)
	fmt.Println("Path 2 paths:", path2)
}

func countPaths(paths map[string][]string, current string, visited map[string]bool) int {
	count := 0
	if current == "end" {
		return 1
	}
	for _, to := range paths[current] {

		// if it was visited or the value is lowercase,skip it
		if visited[to] && strings.ToLower(to) == to {
			continue
		}
		visited[current] = true
		count += countPaths(paths, to, visited)

		// reset for next iteration
		if visited[to] {
			fmt.Println("Resetting", to)
		}
		visited[to] = false
	}
	return count
}

/**
After reviewing the available paths, you realize you might have time to visit a single small cave twice.
Specifically, big caves can be visited any number of times, a single small cave can be visited at most twice,
and the remaining small caves can be visited at most once.
However, the caves named start and end can only be visited exactly once each: once you leave the start cave,
you may not return to it, and once you reach the end cave, the path must end immediately.
*/
func countPathsPart2FancyRules(paths map[string][]string, current string, visited map[string]int, visitedTwice bool) int {
	count := 0
	if current == "end" {
		return 1
	}
	visited[current]++

	for _, to := range paths[current] {
		// nope, skip it
		if to == "start" {
			continue
		}

		// if it was visited or the value is lowercase, check double logic
		if visited[to] > 0 && strings.ToLower(to) == to {
			// already had a double visit, skip it
			if visitedTwice {
				continue
			} else {
				visitedTwice = true
			}
		}
		count += countPathsPart2FancyRules(paths, to, visited, visitedTwice)

		// reset for next iteration
		visited[to]--
		if strings.ToUpper(to) != to && visited[to] == 1 {
			visitedTwice = false
		}
	}
	return count
}
