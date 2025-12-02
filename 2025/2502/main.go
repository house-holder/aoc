package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func hasPattern(id int) bool {
	idStr := strconv.Itoa(id)
	length := len(idStr)
	half := length / 2

	for i := range half {
		if idStr[i] != idStr[half+i] {
			return false
		}
	}

	return length == half*2
}

func evaluateRange(firstID int, lastID int, accumulator *int) {
	for i := firstID; i <= lastID; i++ {
		if hasPattern(i) {
			fmt.Printf("   > Found ID %v has pattern\n", i)
			*accumulator += i
		}
	}
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 && os.Args[1] == "-t" {
		filename = "example.txt"
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open: %v", err)
	}
	defer file.Close()

	body, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}

	accumulator := 0

	for ID := range strings.SplitSeq(string(body), ",") {
		possRange := strings.Split(ID, "-")
		fmt.Printf("Evaluating range: %v\n", possRange)

		firstID, err := strconv.Atoi(possRange[0])
		if err != nil {
			log.Fatalf("Failed to convert: %v", err)
		}

		lastID, err := strconv.Atoi(possRange[1])
		if err != nil {
			log.Fatalf("Failed to convert: %v", err)
		}

		evaluateRange(firstID, lastID, &accumulator)
		fmt.Printf("Accumulator: %v\n", accumulator)
	}
}
