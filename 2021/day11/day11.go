package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 11, Hello.")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := make(map[Point]int)

	row := 0
	for scanner.Scan() {
		value := scanner.Text()
		numbers := stringToNumbers(value)
		for y, count := range numbers {
			grid[Point{x: row, y: y}] = count
		}
		row++
	} // end for scanner.Scan()

	totalFlashes := 0

	for step := 1; step < 100000; step++ {
		flashed := make(map[Point]bool)
		var flash func(p Point)
		flash = func(p Point) {
			// do not flash if already flashed
			if flashed[p] {
				return
			}
			flashed[p] = true
			totalFlashes++

			// flash all neighbors
			for _, n := range p.neighbors() {
				if _, ok := grid[n]; !ok {
					continue
				}
				if flashed[n] {
					continue
				}
				// increment their count
				grid[n] = grid[n] + 1

				// check if this causes another flash
				if grid[n] > 9 {
					flash(n)
				}
			}
		}

		for p, count := range grid {
			grid[p] = count + 1
			if grid[p] > 9 {
				flash(p)
			}
		}

		if len(flashed) == len(grid) {
			// all points flashed
			fmt.Printf("part 2 - Step %d: %d\n", step, totalFlashes)
			break
		}

		// set grid[p] to 0 for all flashed points
		for p := range flashed {
			grid[p] = 0
		}

		if step == 100 {
			fmt.Println("part1 - 100th step:", totalFlashes)
			//	break
		}

	}

	fmt.Println(grid)

}

func stringToNumbers(s string) []int {
	numbers := make([]int, 0)
	for _, c := range s {
		numbers = append(numbers, int(c-'0'))
	}
	return numbers
}

type Point struct {
	x int
	y int
}

func (p Point) neighbors() []Point {
	return []Point{
		{x: p.x - 1, y: p.y},
		{x: p.x - 1, y: p.y - 1},
		{x: p.x - 1, y: p.y + 1},
		{x: p.x + 1, y: p.y},
		{x: p.x + 1, y: p.y - 1},
		{x: p.x + 1, y: p.y + 1},
		{x: p.x, y: p.y - 1},
		{x: p.x, y: p.y + 1},
	}
}
