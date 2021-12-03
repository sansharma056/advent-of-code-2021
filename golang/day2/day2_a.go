package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mulPosDepth() int {
	pos := 0
	depth := 0

	file, _ := os.Open("./input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		unit, _ := strconv.Atoi(instruction[1])

		switch direction := instruction[0]; direction {
		case "forward":
			pos += unit
		case "down":
			depth += unit
		case "up":
			depth -= unit
		}
	}

	return pos * depth
}

func main() {
	fmt.Println(mulPosDepth())
}
