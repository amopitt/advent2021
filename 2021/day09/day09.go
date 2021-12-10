package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*

 */
func main() {
	fmt.Println("Day 9, Hello.")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := make([][]int, 0)

	for scanner.Scan() {
		value := scanner.Text()
		intStr := strings.Split(value, "")
		rowNumbers := make([]int, 0)

		for _, v := range intStr {
			number, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			rowNumbers = append(rowNumbers, number)
		}
		numbers = append(numbers, rowNumbers)

	} // end for scanner.Scan()

	lowPoints, basinCounts := part1(numbers)

	// sort basinCounts
	sort.Ints(basinCounts)

	fmt.Println(lowPoints)
	fmt.Println(basinCounts)
	fmt.Println(basinCounts[len(basinCounts)-1] * basinCounts[len(basinCounts)-2] * basinCounts[len(basinCounts)-3])
}

func part1(numbers [][]int) (int, []int) {
	var lowPoints int
	basinCounts := make([]int, 0)
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {

			leftOf := 10
			rightOf := 10
			above := 10
			below := 10

			numberToCheck := numbers[i][j]

			// check if surrounding numbers are higher than numberToCheck
			if j-1 >= 0 {
				leftOf = numbers[i][j-1]
			}
			if j+1 < len(numbers[i]) {
				rightOf = numbers[i][j+1]
			}
			if i-1 >= 0 {
				above = numbers[i-1][j]
			}
			if i+1 < len(numbers) {
				below = numbers[i+1][j]
			}
			if numberToCheck < leftOf && numberToCheck < rightOf && numberToCheck < above && numberToCheck < below {
				fmt.Println("found a low point:", numberToCheck)

				count := getBasinCount(1, numbers, i, j, numberToCheck)
				fmt.Println("basin count:", count)
				basinCounts = append(basinCounts, count+1)
				lowPoints += (numberToCheck + 1)
			}
			//	fmt.Println(i, j, numberToCheck, leftOf, rightOf, above, below)

		}
	}
	return lowPoints, basinCounts
}

//step 2: check if left, right, above, below are higher than n
// - if any are higher and not equal to 9, add it to the basin
// repeat step 1
// 2199943210
// 3987894921
// 9856789892
// 8767896789
// 9899965678
func getBasinCount(sum int, numbers [][]int, i int, j int, numberToCheck int) (count int) {
	// check if surrounding numbers are higher than numberToCheck

	if j-1 >= 0 {
		leftOf := numbers[i][j-1]
		if leftOf > numberToCheck && leftOf != 9 && leftOf != -1 {
			count++
			numbers[i][j-1] = -1
			count += getBasinCount(sum, numbers, i, j-1, leftOf)
		}
	}

	if j+1 < len(numbers[i]) {
		rightOf := numbers[i][j+1]
		if rightOf > numberToCheck && rightOf != 9 && rightOf != -1 {
			count++
			numbers[i][j+1] = -1
			count += getBasinCount(sum, numbers, i, j+1, rightOf)
		}
	}

	if i-1 >= 0 {
		above := numbers[i-1][j]
		if above > numberToCheck && above != 9 && above != -1 {
			count++
			numbers[i-1][j] = -1
			count += getBasinCount(sum, numbers, i-1, j, above)
		}
	}

	if i+1 < len(numbers) {
		below := numbers[i+1][j]
		if below > numberToCheck && below != 9 && below != -1 {
			count++
			numbers[i+1][j] = -1
			count += getBasinCount(sum, numbers, i+1, j, below)
		}
	}

	fmt.Println("in method count for", numberToCheck, count)
	return count

}
