package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getDepths() []int {
	var depths []int

	file, _ := os.Open("./input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		cur, _ := strconv.Atoi(scanner.Text())

		depths = append(depths, cur)
	}

	return depths
}

func countNumOfIncDepths() int {
	count := 0
	depths := getDepths()

	windowSize := 0
	curWindowsSum := 0
	prevWindowsSum := 0

	for i := 0; i < len(depths); i++ {
		if windowSize < 3 {
			windowSize++
			curWindowsSum += depths[i]
		} else {
			prevWindowsSum = curWindowsSum
			curWindowsSum = curWindowsSum + depths[i] - depths[i-3]

			if curWindowsSum > prevWindowsSum {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countNumOfIncDepths())
}
