package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample.txt
var input string

func main() {
	fmt.Println("Day 13, Hello.")
	fmt.Println("Part 1 input", input)
	lines := strings.Split(input, "\r\n")
	fmt.Println("Part 1 lines", lines, len(lines))
}
