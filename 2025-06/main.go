package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mathProcess(intArray []int, symbol string) int {
	output := 0

	switch symbol {
	case "+":
		for idx := range intArray {
			// fmt.Printf("%d + ", intArray[idx])
			output += intArray[idx]
		}
		// fmt.Printf(" = %d\n", output)
	case "*":
		output = 1
		for idx := range intArray {
			// fmt.Printf("%d * ", intArray[idx])
			output *= intArray[idx]
		}
		// fmt.Printf(" = %d\n", output)
	default:
		log.Fatalf("Invalid symbol: %s", symbol)
	}

	return output
}

func getIntsAndSymbols(lines []string) ([][]int, []string) {
	intArrays := [][]int{}
	symbolArray := []string{}
	for idx := range lines {
		tempArray := []int{}
		for element := range strings.SplitSeq(lines[idx], " ") {
			if element == "" || element == " " {
				continue
			}
			if idx < len(lines)-1 {
				temp, err := strconv.Atoi(element)
				if err != nil {
					log.Fatalf("strconv.Atoi failed. %v", err)
				}
				tempArray = append(tempArray, temp)
			} else {
				symbolArray = append(symbolArray, element)
			}
		}
		intArrays = append(intArrays, tempArray)
	}
	return intArrays, symbolArray
}

func eval1(lines []string) int {
	sum := 0
	intArrays, symbolArray := getIntsAndSymbols(lines)

	for symIndex, symbol := range symbolArray {
		tempArray := []int{}
		for array := range intArrays {
			if len(intArrays[array]) == 0 {
				continue
			}
			tempArray = append(tempArray, intArrays[array][symIndex])
		}
		sum += mathProcess(tempArray, symbol)
	}

	return sum
}

func stringDataToInt(lines []string, idx int) int {
	temp := ""
	symbolRow := len(lines) - 1

	for i := range symbolRow {
		if string(lines[i][idx]) == " " {
			continue
		}
		temp += string(lines[i][idx])
	}
	if temp == "" {
		return 0
	}

	tempInt, err := strconv.Atoi(temp)
	if err != nil {
		log.Fatalf("strconv.Atoi failed. %v", err)
	}
	return tempInt
}

func eval2(lines []string) int {
	sum := 0
	accum := []int{}
	symbolLine := lines[len(lines)-1]

	for i := len(symbolLine) - 1; i >= 0; i-- {
		char := symbolLine[i]
		val := stringDataToInt(lines, i)
		if val != 0 {
			accum = append(accum, val)
		}
		// fmt.Printf("accum: %v\n", accum)
		if char != ' ' {
			// fmt.Printf("Stop on symbol: %s\n", string(char))
			sum += mathProcess(accum, string(char))
			accum = []int{}
		}
	}
	return sum
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	lines := strings.Split(body, "\n")

	result1 := eval1(lines)
	result2 := eval2(lines)
	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
}
