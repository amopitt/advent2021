package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 18")
	fmt.Println(input)
	fmt.Println("------------------")
}
