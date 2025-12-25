package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Box struct {
	dims    []int
	sfcArea int
	vol     int
}

func NewBox(line string) Box {
	strs := strings.Split(line, "x")
	length, err := strconv.Atoi(strs[0])
	if err != nil {
		log.Fatal("strconv failed:", err)
	}
	width, err := strconv.Atoi(strs[1])
	if err != nil {
		log.Fatal("strconv failed:", err)
	}
	height, err := strconv.Atoi(strs[2])
	if err != nil {
		log.Fatal("strconv failed:", err)
	}

	surfaceArea := 2 * ((length * width) +
		(width * height) + (length * height))

	dimensions := []int{length, width, height}
	slices.Sort(dimensions)

	return Box{
		dims:    dimensions,
		sfcArea: surfaceArea,
		vol:     length * width * height,
	}

}

func minDims(dims []int) (mins []int) {
	lowest := math.MaxInt
	lowest = min(dims[0], lowest)
	return []int{4, 5}
}

func evalPart1(lines []string) int {
	total := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		box := NewBox(line)
		total += box.sfcArea + (box.dims[0] * box.dims[1])
	}

	return total
}

func evalPart2(lines []string) int {
	total := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		box := NewBox(line)

		total += 2 * (box.dims[0] + box.dims[1])
		total += box.vol // forgot ribbon
	}

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
