package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func rebuild(lines []string) []string {
	output := []string{}

	for lineIdx, line := range lines {
		if line == "" {
			continue
		}
		if lineIdx == 0 && strings.Contains(line, "S") {
			output = append(output, line)
			continue
		}
		newLine := ""
		prevLine := string(output[len(output)-1])

		for charIdx, char := range line {
			prevChar := prevLine[charIdx]
			incrIdx := int(math.Min(float64(charIdx+1), float64(len(line))))
			decrIdx := int(math.Max(float64(charIdx-1), 0))
			before := charIdx < len(line)-1
			after := charIdx > 0

			if before && line[incrIdx] == '^' && prevLine[incrIdx] == '|' {
				newLine += "|"
				continue
			}
			if after && line[decrIdx] == '^' && prevLine[decrIdx] == '|' {
				newLine += "|"
				continue
			}
			if (prevChar == '|' && char != '^') || prevChar == 'S' {
				newLine += "|"
				continue
			} else if char == '^' {
				newLine += "^"
			} else {
				newLine += "."
			}
		}
		output = append(output, newLine)
	}
	return output
}

func evalPart1(rawLines []string) int {
	prev := ""
	splitCount := 0
	lines := rebuild(rawLines)

	for i, line := range lines {
		if i == 0 || line == "" || !strings.Contains(line, "^") {
			continue
		}
		prev = lines[i-1]
		for j, char := range line {

			if string(char) == "^" && string(prev[j]) == "|" {
				splitCount++
			}
		}
	}
	return splitCount
}

func evalPart2(rawLines []string) int {
	timelines := 0
	lines := rebuild(rawLines)

	width := len(lines[0])
	paths := make([]int, width)
	paths[strings.Index(lines[0], "S")] = 1
	prevPaths := paths

	for _, line := range lines[1:] {
		paths = make([]int, width)
		for charIdx, char := range line {
			pathsToPPOS := 0
			if char == '|' {
				pathsToPPOS += prevPaths[charIdx]
				if charIdx < width-1 && line[charIdx+1] == '^' {
					pathsToPPOS += prevPaths[charIdx+1]
				}
				if charIdx > 0 && line[charIdx-1] == '^' {
					pathsToPPOS += prevPaths[charIdx-1]
				}
			}
			paths[charIdx] = pathsToPPOS
		}
		prevPaths = paths
	}
	for _, val := range paths {
		timelines += val
	}
	return timelines
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	lines := strings.Split(body, "\n")

	part1Result := evalPart1(lines)
	part2Result := evalPart2(lines)
	fmt.Printf("\nPart 1: %v", part1Result)
	fmt.Printf("\tPart 2: %v\n", part2Result)
}
