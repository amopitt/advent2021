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

	for scanner.Scan() {
		value := scanner.Text()
		fmt.Println(value)
	} // end for scanner.Scan()

}
