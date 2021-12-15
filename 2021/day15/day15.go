package main

import (
	"advent2021/util"
	"container/heap"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample.txt
var input string

type Node struct {
	pos      coord
	vertexes map[coord]int
	visited  bool
	distance int
}

func main() {
	fmt.Println("Day 15")

	fmt.Println(input)
	fmt.Println("------------------")
	values := strings.Split(input, "\r\n")

	width := 0
	height := 0

	coords := make(map[coord]int)
	doPart2 := true

	for rowNum, line := range values {
		numbers := strings.Split(line, "")
		for x, number := range numbers {
			coords[coord{x, rowNum}] = util.GetInt(number)
		}
		if len(numbers) > width {
			width = len(numbers)
		}
		height++
	}

	if doPart2 {
		nr := make(map[coord]int)
		for rx := 0; rx <= 4; rx++ {
			for ry := 0; ry <= 4; ry++ {

				// loop through x/y of coords width
				for pos, risk := range coords {
					x, y := pos.x, pos.y
					nextRisk := (risk + rx + ry)
					if nextRisk > 9 {
						nextRisk -= 9
					}

					nr[coord{x: x + rx*width, y: y + ry*height}] = nextRisk
				}
			}
		}
		coords = nr
		width = width * 5
		height = height * 5
	}

	// create a map of source/destination and weight of all possible directions
	nodez := make(map[coord]Node)

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

		nodez[c] = f
	}

	// part 1
	// visited := make(map[coord]bool)
	printGraph(coords, width)

	fmt.Println(width, height)
	part1 := part1(coords, width, height, coord{0, 0}, nodez, coord{width - 1, height - 1})
	fmt.Println("part 1:", part1)

}

func printGraph(graph map[coord]int, width int) {
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(graph[coord{x, y}])
		}
		fmt.Println()
	}
}

func printGraphDistance(graph map[coord]Node, width int) {
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(graph[coord{x, y}].distance)
			fmt.Print(",")
		}
		fmt.Println()
	}
}

func part1(coords map[coord]int, width int, height int, start coord, graph map[coord]Node, end coord) int {
	fmt.Println(start, end)

	// distance to self is 0
	if n, ok := graph[start]; ok {
		n.distance = 0
		graph[start] = n
	}

	// nope, queue doesn't work, researched into how to do a priority queue
	// queue := make([]coord, 0)
	// queue = append(queue, start)

	// make a priority queue
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{value: start, priority: 0}
	heap.Init(&pq)

	// while queue is not empty
	for pq.Len() > 0 {
		// pop the highest priority from queue
		item := heap.Pop(&pq).(*Item)
		current := item.value

		// skip if already visited
		if graph[current].visited {
			continue
		}

		// mark as visited using janky go syntax
		if n, ok := graph[current]; ok {
			n.visited = true
			graph[current] = n
		}

		// the promised land
		if current == end {
			printGraphDistance(graph, width)
			return graph[end].distance
		}

		// loop through each vertex to find the minimum distance
		for neighbor, risk := range graph[current].vertexes {
			// keep going if we've already been there
			if graph[neighbor].visited {
				continue
			}

			// add the risk to the current distance
			moveTo := graph[current].distance + risk

			if gn, ok := graph[neighbor]; ok {
				if moveTo < graph[neighbor].distance {
					// update to the new distance
					gn.distance = moveTo
					graph[neighbor] = gn
				}

				// add it to the priority queue using example code
				// priority would be to move to the smallest distance
				heap.Push(&pq, &Item{value: neighbor, priority: gn.distance})
				pq.update(&Item{value: neighbor, priority: gn.distance}, neighbor, gn.distance)
			}
		}
	}

	// dijkstra's algorithm for shortest path
	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
	// Create map to track distances from source vertex
	return 911 // help i suck
}

type coord struct {
	x int
	y int
}

// i janked this code from https://go.dev/pkg/container/heap/#example__priorityQueue

// An Item is something we manage in a priority queue.
type Item struct {
	value    coord // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the LOWEST, not highest, priority so we use less than here.
	return pq[i].priority < pq[j].priority
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
func (pq *PriorityQueue) update(item *Item, value coord, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
