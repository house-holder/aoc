package main

// begin time: 1764825035
//
//
// stop time:

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func evaluatePart1(body string) int {
	accessibleRollCount := 0
	maxAdjacentRolls := 3
	allRows := strings.Split(body, "\n")

	for i := range allRows {
		// fmt.Printf("Evaluating row %d\n", i+1)
		rowStartIdx := max(0, i-1)
		rowStopIdx := min(len(allRows)-1, i+1)
		evalRows := allRows[rowStartIdx : rowStopIdx+1]

		for j := range allRows[i] {
			colStartIdx := max(0, j-1)
			colStopIdx := min(len(allRows[i])-1, j+1)

			if allRows[i][j] == '.' {
				continue
			}
			// fmt.Printf("  Roll=(%d, %d)", i, j)

			adjacentCount := 0
			// fmt.Printf(" reset adj=0\n")
			for k, row := range evalRows {
				// fmt.Printf("    evalRow %d\n", k+rowStartIdx)
				for l, char := range row[colStartIdx : colStopIdx+1] {
					// fmt.Printf("      coords=(%d, %d) ", k+rowStartIdx, l+colStartIdx)
					if k+rowStartIdx == i && l+colStartIdx == j {
						// fmt.Printf("(self! skip)\n")
						continue
					}
					if char == '@' {
						// fmt.Printf("@")
						adjacentCount++
						// fmt.Printf(" (adj++, %d)", adjacentCount)
					}
				}
			}
			if adjacentCount <= maxAdjacentRolls {
				// fmt.Printf(" increment accessible\n")
				accessibleRollCount++
			}
		}
	}
	// fmt.Printf("accessibleRollCount: %d\n", accessibleRollCount)
	return accessibleRollCount
}

func evaluatePart2(body string) int {
	rollsRemoved := 0
	removedCount := 0
	canRemoveRolls := true
	maxAdjacentRolls := 3

	for canRemoveRolls {
		accessibleRollCount := 0
		coordAccum := [][]int{}
		allRows := strings.Split(body, "\n")

		for i := range allRows {
			rowStartIdx := max(0, i-1)
			rowStopIdx := min(len(allRows)-1, i+1)
			evalRows := allRows[rowStartIdx : rowStopIdx+1]

			for j := range allRows[i] {
				adjacentCount := 0
				coords := []int{}
				colStartIdx := max(0, j-1)
				colStopIdx := min(len(allRows[i])-1, j+1)

				if allRows[i][j] == '.' || allRows[i][j] == 'x' {
					continue
				}

				for k, row := range evalRows {
					for l, char := range row[colStartIdx : colStopIdx+1] {
						coords = []int{i, j}
						if k+rowStartIdx == i && l+colStartIdx == j {
							continue
						}
						if char == '@' {
							adjacentCount++
						}
					}
				}
				if adjacentCount <= maxAdjacentRolls {
					accessibleRollCount++
					coordAccum = append(coordAccum, coords)
				}
			}
		}
		if accessibleRollCount == 0 {
			canRemoveRolls = false
		} else {
			body, removedCount = replaceRolls(body, coordAccum)
			// fmt.Printf("Removed %d rolls\n", removedCount)
			rollsRemoved += removedCount
		}
		//prompt to continue for testing
		// fmt.Printf("Press Enter to continue...")
		// fmt.Scanln()
	}
	return rollsRemoved
}

func replaceRolls(body string, coords [][]int) (string, int) {
	if len(coords) == 0 {
		return body, 0
	}
	removedCount := 0
	lines := strings.Split(body, "\n")
	outputLines := []string{}

	lineIndices := []int{}

	for lineIdx, _ := range lines {
		for _, coord := range coords {
			if coord[0] == lineIdx {
				lineIndices = append(lineIndices, coord[1])
			}
		}
		// fmt.Printf("Replace in line %d at positions %v\n", lineIdx, lineIndices)

		charAccum := ""
		for charIdx, char := range lines[lineIdx] {
			if slices.Contains(lineIndices, charIdx) {
				charAccum += "x"
				removedCount++
			} else {
				charAccum += string(char)
			}
			// fmt.Printf("Built line: %s\n", charAccum)
		}
		lineIndices = []int{}
		outputLines = append(outputLines, charAccum)
	}
	outputBody := strings.Join(outputLines, "\n")
	// fmt.Printf("Output body:\n%s\n", outputBody)
	return outputBody, removedCount
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)

	result1 := evaluatePart1(body)
	result2 := evaluatePart2(body)
	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}
