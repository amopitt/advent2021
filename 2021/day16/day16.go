package main

import (
	_ "embed"
	"fmt"
)

//go:embed sample.txt
var input string

func main() {
	fmt.Println("Day 16")

	fmt.Println(input)
	fmt.Println("------------------")
}
