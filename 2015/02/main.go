package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func lineToDims(line string) [3]int {
	strs := strings.Split(line, "x")
	fmt.Println(strs)
	return [3]int{1, 2, 3}
}

func sfcArea(x, y int) int {
	return x * y
}

func minDims(dims [3]int) (mins [2]int) {
	return [2]int{4, 5}
}

func evalPart1(lines []string) int {
	total := 0

	for _, line := range lines {
		dims := lineToDims(line)

	}

	return total
}

func evalPart2(lines []string) int {
	total := 0

	return total
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	lines := strings.Split(body, "\n")

	result1 := evalPart1(lines)
	fmt.Println("Part 1:", result1)
	result2 := evalPart2(lines)
	fmt.Println("Part 2:", result2)
}
