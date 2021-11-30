package util

import (
	"bufio"
	"os"
	"strconv"
)

// ReadFile reads a file and returns a slice of strings
func ReadFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadFileInts reads a file and returns a slice of ints
func ReadFileInts(fileName string) ([]int, error) {
	lines, err := ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var ints []int
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
