package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	pos      coord
	vertexes map[coord]int
	visited  bool
	distance int
}

type Graph struct {
	nodes []Node
}

func main() {
	fmt.Println("Day 15")
	fmt.Println(input)

	values := strings.Split(input, "\r\n")

	width := 0
	height := 0

	coords := make(map[coord]int)

	//graph.nodes = make([]Node, 0)

	for rowNum, line := range values {
		numbers := strings.Split(line, "")
		for x, number := range numbers {
			coords[coord{x, rowNum}] = getInt(number)
		}
		if len(numbers) > width {
			width = len(numbers)
		}
		height++
	}

	nr := make(map[coord]int)
	for rx := 0; rx <= 4; rx++ {
		for ry := 0; ry <= 4; ry++ {
			//for pos, risk := range coords {

			// loop through x/y of coords width
			for pos, risk := range coords {
				x, y := pos.x, pos.y
				nextRisk := (risk + rx + ry)
				if nextRisk > 9 {
					nextRisk -= 9
				}

				nr[coord{x: x + rx*width, y: y + ry*height}] = nextRisk
			}

			//	}
		}
	}
	coords = nr

	// create a map of source/destination and weight of all possible directions
	nodez := make(map[coord]Node)

	width = width * 5
	height = height * 5
	for c := range coords {
		f := Node{}
		f.distance = 99999999999999
		f.pos = c
		f.vertexes = make(map[coord]int)

		// add up and left
		if c.y-1 >= 0 {
			f.vertexes[coord{c.x, c.y - 1}] = coords[coord{c.x, c.y - 1}]

		}
		if c.x-1 >= 0 {
			f.vertexes[coord{c.x - 1, c.y}] = coords[coord{c.x - 1, c.y}]
		}
		if c.y+1 < height {
			f.vertexes[coord{c.x, c.y + 1}] = coords[coord{c.x, c.y + 1}]
		}
		if c.x+1 < width {
			f.vertexes[coord{c.x + 1, c.y}] = coords[coord{c.x + 1, c.y}]
		}

		//	mw.nodes = append(mw.nodes, f)
		nodez[c] = f
	}

	// part 1
	// visited := make(map[coord]bool)
	printGraph(coords)

	fmt.Println(width, height)
	part1 := part1(coords, width, height, coord{0, 0}, nodez, coord{width - 1, height - 1})
	fmt.Println("part 1:", part1)

	// sum part 1
	// sum := 0
	// for _, val := range part1 {
	// 	sum += val
	// }
	// fmt.Println("sum:", sum)

}

func printGraph(graph map[coord]int) {
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			fmt.Print(graph[coord{x, y}])
		}
		fmt.Println()
	}
}

func part1(coords map[coord]int, width int, height int, start coord, graph map[coord]Node, end coord) int {
	fmt.Println(start, end)
	distances := make(map[coord]int)
	distances[start] = 0

	if n, ok := graph[start]; ok {
		n.distance = 0
		graph[start] = n
	}

	// make a queue
	queue := make([]coord, 0)
	queue = append(queue, start)

	// while queue is not empty
	for len(queue) > 0 {
		// pop from queue
		current := queue[0]
		queue = queue[1:]

		if graph[current].visited {
			continue
		}

		if n, ok := graph[current]; ok {
			n.visited = true
			graph[current] = n
		}

		if current == end {
			return graph[end].distance
		}

		// for each neighbor
		for neighbor, weight := range graph[current].vertexes {
			// if neighbor is not visited
			if graph[neighbor].visited {
				continue
			}
			t := graph[current].distance + weight

			if gn, ok := graph[neighbor]; ok {
				if t < graph[neighbor].distance {
					gn.distance = t
					graph[neighbor] = gn
				}
				if gn.distance != 99999999999999 {
					queue = append(queue, neighbor)
				}
			} else {
				fmt.Println("not ok", neighbor)
			}
		}
	}

	// dijkstra's algorithm for shortest path
	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
	// Create map to track distances from source vertex
	return 35333
}

// func oldPart1(coords map[coord]int, width int, height int, position coord, visited map[coord]bool, minJumps int) (count int, completed bool) {

// 	// part 1
// 	if visited[position] {
// 		return 0, false
// 	}

// 	// reached the end
// 	if position.x == width-1 && position.y == height-1 {
// 		fmt.Println("reached the end", len(visited), position.x, position.y)
// 		return coords[position], true
// 	}

// 	visited[position] = true
// 	count += coords[position]

// 	// find lowest neighbor to move to
// 	lowest := -1
// 	lowestNeighbor := []coord{{-1, -1}}

// 	// find lowest neighbor not-diagonally
// 	for x := position.x - 1; x <= position.x+1; x++ {
// 		for y := position.y - 1; y <= position.y+1; y++ {
// 			if x == position.x && y == position.y {
// 				continue
// 			}
// 			if x < 0 || y < 0 || x >= width || y >= height {
// 				continue
// 			}
// 			if visited[coord{x, y}] {
// 				continue
// 			}

// 			// if the move is diagonal, skip it
// 			if x == position.x-1 && y == position.y-1 || x == position.x+1 && y == position.y+1 {
// 				continue
// 			} else if x == position.x-1 && y == position.y+1 || x == position.x+1 && y == position.y-1 {
// 				continue
// 			} else if x == position.x-1 || x == position.x+1 {
// 				if y == position.y-1 || y == position.y+1 {
// 					continue
// 				}
// 			} else if y == position.y-1 || y == position.y+1 {
// 				if x == position.x-1 || x == position.x+1 {
// 					continue
// 				}
// 			}

// 			if coords[coord{x, y}] < lowest || lowest == -1 {
// 				lowest = coords[coord{x, y}]
// 				lowestNeighbor = []coord{{x, y}}
// 			} else if coords[coord{x, y}] == lowest {
// 				lowestNeighbor = append(lowestNeighbor, coord{x, y})
// 			}

// 		}
// 	}
// 	end := false
// 	var add int
// 	// try each lowest neighbor
// 	for _, neighbor := range lowestNeighbor {
// 		if visited[neighbor] {
// 			continue
// 		}
// 		if neighbor.x == -1 {
// 			continue
// 		}
// 		if neighbor.y == -1 {
// 			continue
// 		}
// 		add, end = part1(coords, width, height, neighbor, visited, minJumps)
// 		count += add
// 		if end {
// 			// loop over visited
// 			jumps := 0
// 			for _, v := range visited {
// 				if v {
// 					jumps++
// 				}
// 			}
// 			if jumps <= minJumps {
// 				minJumps = jumps
// 				//fmt.Println("loop", jumps, count)
// 			}
// 			fmt.Println("count", count, len(visited), jumps, minJumps)

// 			end = false
// 		}
// 	}
// 	fmt.Println("returning", count, end)
// 	return count, end
// }

func getInt(val string) int {
	if i, err := strconv.Atoi(val); err == nil {
		return i
	}
	return -1
}

type coord struct {
	x int
	y int
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
