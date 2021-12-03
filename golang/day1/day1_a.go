package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countNumOfIncDepths() int {
	count := 0
	file, _ := os.Open("./input")

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	prev, _ := strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		cur, _ := strconv.Atoi(scanner.Text())
		if cur > prev {
			count++
		}

		prev = cur
	}

	return count
}

func main() {
	fmt.Println(countNumOfIncDepths())
}
