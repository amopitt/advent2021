package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 13, Hello.")
	lines := strings.Split(input, "\r\n")

	var foldDirections []fold
	points := make(map[coord]bool)

	for _, line := range lines {
		// if line starts with fold along
		if strings.HasPrefix(line, "fold along") {
			foldParts := strings.Split(line, "fold along ")
			foldLine := foldParts[1]
			f := strings.Split(foldLine, "=")
			distance, _ := strconv.Atoi(f[1])
			foldDirections = append(foldDirections, fold{f[0], distance})
		} else {
			if line == "" {
				continue
			}
			// if line starts with a point
			p := strings.Split(line, ",")

			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[1])
			points[coord{x, y}] = true
		}
	}
	part1 := part1(points, foldDirections)
	fmt.Println("Part 1", part1)

	part2 := part2(points, foldDirections)
	draw(part2)
}

func draw(part2 map[coord]bool) {
	// The manual says the code is always eight capital letters.
	for lineNum := 0; lineNum < 6; lineNum++ {
		for x := 0; x < 8*4+7; x++ { // 8 letters * width 4 + 7 total spaces between letters
			if part2[coord{x, lineNum}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (f fold) inFold(p coord) bool {
	if f.direction == "x" {
		return p.x == f.distance
	} else {
		return p.y == f.distance
	}
}

func part1(points map[coord]bool, foldDirections []fold) int {

	// loop over fold directions
	for _, f := range foldDirections {
		// loop over points
		for p := range points {
			// if point is in fold direction
			if f.inFold(p) {
				// remove point
				points[p] = false
				continue
			}

			if f.direction == "x" {
				if p.x > f.distance {
					// remove point
					points[p] = false

					// get fold location
					transposed := f.distance - (p.x - f.distance)
					// if transposed < 0 {
					// 	continue
					// }

					// add point to fold
					points[coord{transposed, p.y}] = true
					//fmt.Println("adding", coord{f.distance - (p.x - f.distance), p.y}, "from fold", f, "prev point=", p)
				}
			} else if f.direction == "y" {
				if p.y > f.distance {
					// remove point
					points[p] = false

					transposed := f.distance - (p.y - f.distance)
					// if transposed < 0 {
					// 	continue
					// }

					// add point to fold
					points[coord{p.x, transposed}] = true

					//fmt.Println("adding", coord{p.x, f.distance - (p.y - f.distance)}, "from fold", f, "prev point=", p)
				}
			}

		}
		break
	}

	// loop over points
	count := 0
	for _, v := range points {
		if v {
			count++
		}
	}
	return count
}

func part2(points map[coord]bool, foldDirections []fold) (output map[coord]bool) {

	// loop over fold directions
	for _, f := range foldDirections {
		// loop over points
		for p := range points {
			// if point is in fold direction
			if f.inFold(p) {
				// remove point
				points[p] = false
				continue
			}

			if f.direction == "x" {
				if p.x > f.distance {
					// remove point
					points[p] = false

					// add point to fold
					transposed := f.distance - (p.x - f.distance)
					points[coord{transposed, p.y}] = true
					//	fmt.Println("adding", coord{f.distance - (p.x - f.distance), p.y}, "from fold", f, "prev point=", p)
				}
			} else if f.direction == "y" {
				if p.y > f.distance {
					// remove point
					points[p] = false

					// add point to fold
					transposed := f.distance - (p.y - f.distance)
					points[coord{p.x, transposed}] = true

					//	fmt.Println("adding", coord{p.x, f.distance - (p.y - f.distance)}, "from fold", f, "prev point=", p)
				}
			}

		}
	}

	return points
}

type fold struct {
	direction string
	distance  int
}

type coord struct {
	x int
	y int
}
