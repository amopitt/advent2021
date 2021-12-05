package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMappedCoordinates(fileName string, allowDiagonal bool) (map[string]int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create hash map of mappedCoords and counts
	mappedCoords := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := scanner.Text()
		coords := strings.Split(path, " -> ")
		from := coords[0]
		to := coords[1]
		updateMappedCoords(mappedCoords, from, to, allowDiagonal)
	}

	return mappedCoords, getOverlapCount(mappedCoords)
}

func getOverlapCount(mappedCoords map[string]int) int {
	overlapCount := 0
	for _, v := range mappedCoords {
		if v > 1 {
			// fmt.Printf("%s appears %d times\n", k, v)
			overlapCount++
		}
	}
	return overlapCount
}

func Day5() {
	fileName := "day05/input.txt"

	// part 1 -
	// Consider only horizontal and vertical lines. At how many points do at least two lines overlap?
	_, overlapCount := getMappedCoordinates(fileName, false)
	fmt.Printf("part 1 (horiz/vertical): %d\n", overlapCount)

	// part 2 -
	// Now consider diagonals. At how many points do at least two lines overlap?
	_, overlapCount = getMappedCoordinates(fileName, true)
	fmt.Printf("part 2 (horiz/vertical/diagonal): %d\n", overlapCount)
}

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) toString() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func updateMappedCoords(mappedCoords map[string]int, from string, to string, allowDiagonal bool) {
	fromCoord := getCoordinate(from)
	toCoord := getCoordinate(to)

	// horizontal
	if fromCoord.x == toCoord.x {
		min := min(fromCoord.y, toCoord.y)
		max := max(fromCoord.y, toCoord.y)

		for y := min; y <= max; y++ {
			mappedCoords[Coordinate{fromCoord.x, y}.toString()]++
		}
	} else if fromCoord.y == toCoord.y {
		min := min(fromCoord.x, toCoord.x)
		max := max(fromCoord.x, toCoord.x)

		// vertical
		for x := min; x <= max; x++ {
			mappedCoords[Coordinate{x, fromCoord.y}.toString()]++
		}
	} else if allowDiagonal {
		// diagonal line

		// get the slope
		slope := float64(toCoord.y-fromCoord.y) / float64(toCoord.x-fromCoord.x)

		// get the intercept
		intercept := float64(fromCoord.y) - slope*float64(fromCoord.x)

		// get the min and max x and y
		minX := min(fromCoord.x, toCoord.x)
		maxX := max(fromCoord.x, toCoord.x)
		minY := min(fromCoord.y, toCoord.y)
		maxY := max(fromCoord.y, toCoord.y)

		// loop through the x and y values
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				// check if the point is on the line
				if float64(y) == slope*float64(x)+intercept {
					mappedCoords[Coordinate{x, y}.toString()]++
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getCoordinate(coord string) Coordinate {
	split := strings.Split(coord, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return Coordinate{x, y}
}
