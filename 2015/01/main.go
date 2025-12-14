package main

import (
	"fmt"
	"log"
	"os"
)

func evalPart1(body string) int {
	value := 0
	for _, char := range body {
		switch c := char; c {
		case '(':
			value++
		case ')':
			value--
		}
	}
	return value
}

func evalPart2(body string) int {
	value := 0
	for i, char := range body {
		switch c := char; c {
		case '(':
			value++
		case ')':
			value--
			if value == -1 {
				return i + 1
			}
		}
	}
	return value
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)

	r1 := evalPart1(body)
	fmt.Println("Part 1:", r1)
	r2 := evalPart2(body)
	fmt.Println("Part 2:", r2)
}
