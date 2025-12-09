package main

import (
	"fmt"
	"log"
	"os"
)

type Pos struct {
	X int
	Y int
	Z int
}

var allConns = make(map[float64]*Pos)

func evalPart1(input string) int {
	result := 1

	return result
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)

	result1 := evalPart1(body)
	fmt.Printf("Part 1: %d\n", result1)
}
