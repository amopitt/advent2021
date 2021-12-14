package main

import (
	_ "embed"
	"fmt"
)

//go:embed sample.txt
var input string

func main() {
	fmt.Println("Day 15")
	fmt.Println(input)
}
