package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 3")
	fmt.Println(input)
	fmt.Println("------------------")

	lines := strings.Split(input, "\r\n")
	trees := make(map[coord]bool)
	width := len(lines[0])
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				trees[coord{x, y}] = true
			}
		}
	}
	fmt.Println(trees)

	fmt.Println("Part 1:", getTreesHit(lines, trees, 3, 1, width, coord{0, 0}))

	slopes := []coord{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	prod := 1
	for _, slope := range slopes {
		prod *= getTreesHit(lines, trees, slope.x, slope.y, width, coord{0, 0})
	}

	fmt.Println("Part 2:", prod)
}

func getTreesHit(lines []string, trees map[coord]bool, dx, dy int, width int, start coord) int {
	treesHit := 0
	for start.y < len(lines) {
		if trees[coord{start.x % width, start.y}] {
			treesHit++
		}
		start.x += dx
		start.y += dy
	}
	return treesHit
}

type coord struct {
	x, y int
}
