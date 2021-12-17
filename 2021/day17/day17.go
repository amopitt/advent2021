package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 17")
	fmt.Println(input)
	fmt.Println("------------------")
	minX, maxX, minY, maxY := getInput(false)
	fmt.Println(minX, maxX, minY, maxY)
	p1 := part1(minX, maxX, minY, maxY)
	fmt.Println("Part 1:", p1)
	p2 := part2(minX, maxX, minY, maxY)
	fmt.Println("Part 2:", p2)
}

func step(x, y, dx, dy int) (int, int, int, int) {
	x += dx
	y += dy
	if dx > 0 {
		dx += -1
	}
	dy += -1
	return x, y, dx, dy
}

func runStep(dx, dy, minX, maxX, minY, maxY int) (bool, int) {
	x, y := 0, 0
	ans := 0
	for y >= minY {
		x, y, dx, dy = step(x, y, dx, dy)
		ans = max(ans, y)
		// if it's within the target area
		if minX <= x && x <= maxX && minY <= y && y <= maxY {
			return true, ans
		}
	}
	return false, ans
}

func part1(minX, maxX, minY, maxY int) int {
	ans := 0
	// brute force. yessir
	for i := -1000; i <= 1000; i++ {
		for j := -1000; j <= 1000; j++ {
			safe, maxHeight := runStep(i, j, minX, maxX, minY, maxY)
			if safe {
				ans = max(ans, maxHeight)
			}
		}
	}
	return ans
}

func part2(minX, maxX, minY, maxY int) int {
	ans := 0
	for i := -1000; i <= 1000; i++ {
		for j := -1000; j <= 1000; j++ {
			safe, _ := runStep(i, j, minX, maxX, minY, maxY)
			if safe {
				ans += 1
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// didnt feel lke parsing input today.
func getInput(isSample bool) (minX, maxX, minY, maxY int) {
	if isSample {
		return 20, 30, -10, -5
	} else {
		return 57, 116, -198, -148
	}
}
