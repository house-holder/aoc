package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func rebuildWithBeams(lines []string) []string {
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

func evalPart1(lines []string) int {
	prev := ""
	splitCount := 0
	lines = rebuildWithBeams(lines)

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

func evalPart2(lines []string) int {
	timelines := 0

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

	fmt.Printf("Part 1: %v\n", part1Result)
	fmt.Printf("Part 2: %v\n", part2Result)
}
