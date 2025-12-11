package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func getCoords(input string) Coordinate {
	strPair := strings.Split(input, ",")
	valX, err := strconv.Atoi(strPair[0])
	if err != nil {
		log.Fatalf("strconv.Atoi failed. %v", err)
	}
	valY, err := strconv.Atoi(strPair[1])
	if err != nil {
		log.Fatalf("strconv.Atoi failed. %v", err)
	}

	return Coordinate{
		X: valX,
		Y: valY,
	}
}

func findGreatestArea(input []Coordinate) int {
	maxArea := 0
	for i, cornerA := range input {
		for j, cornerB := range input {
			if i != j {
				thisArea := findArea(cornerA, cornerB)
				if maxArea < thisArea {
					maxArea = thisArea
				}
			}
		}
	}
	return maxArea
}

func findArea(cornerA Coordinate, cornerB Coordinate) int {
	latSize := 1 + cornerB.X - cornerA.X
	vertSize := 1 + cornerB.Y - cornerA.Y
	return int(math.Abs(float64(latSize * vertSize)))
}

func convertCoords(input string) []Coordinate {
	output := []Coordinate{}
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) > 0 {
			newCoords := getCoords(line)
			output = append(output, newCoords)
		}
	}
	return output
}

func evalPart1(input []Coordinate) int {
	return findGreatestArea(input)
}

func evalPart2(input []Coordinate) int {

	return len(input)
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	coords := convertCoords(body)

	fmt.Println()
	result1 := evalPart1(coords)
	result2 := evalPart2(coords)
	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}
