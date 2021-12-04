package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binaryDiagnostic() {
	file, _ := os.Open("./input")
	scanner := bufio.NewScanner(file)

	var nums [1000][12]int

	i := 0
	for scanner.Scan() {
		var row [12]int
		for index, val := range scanner.Text() {
			num, _ := strconv.Atoi(string(val))
			row[index] = num
		}
		nums[i] = row
		i++
	}

	var counter [len(nums[0])]int
	for _, row := range nums {
		for index, val := range row {
			counter[index] += val
		}
	}

	var gammaArray [len(nums[0])]string
	var epsilonArray [len(nums[0])]string

	for index, val := range counter {
		if val > len(nums)/2 {
			gammaArray[index] = "1"
			epsilonArray[index] = "0"
		} else {
			gammaArray[index] = "0"
			epsilonArray[index] = "1"
		}
	}

	gamma, _ := strconv.ParseInt(strings.Join(gammaArray[:], ""), 2, 64)
	epsilon, _ := strconv.ParseInt(strings.Join(epsilonArray[:], ""), 2, 64)

	fmt.Println(gamma * epsilon)
}

func main() {
	binaryDiagnostic()
}
