package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample.txt
var input string

func main() {
	fmt.Println("Day 14, Hello.")
	lines := strings.Split(input, "\r\n")

	for _, line := range lines {
		fmt.Println(line)
	}
}
