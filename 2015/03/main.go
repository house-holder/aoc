package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

func visitHouses(input string) map[Coord]int {
	pos := Coord{0, 0}
	visited := make(map[Coord]int)
	visited[pos] = 1

	for _, char := range input {
		switch char {
		case '^':
			pos.y++
		case 'v':
			pos.y--
		case '>':
			pos.x++
		case '<':
			pos.x--
		}

		_, ok := visited[pos]
		if ok {
			visited[pos]++
		} else {
			visited[pos] = 1
		}
	}

	return visited
}

func evalPart1(input string) int {
	visited := visitHouses(input)
	return len(visited)
}

func evalPart2(input string) int {
	var santaBuild, roboBuild strings.Builder

	for i, char := range input {
		if i%2 != 0 {
			santaBuild.WriteRune(char)
		} else {
			roboBuild.WriteRune(char)
		}
	}

	santaInstructions := santaBuild.String()
	santaVisited := visitHouses(santaInstructions)

	roboInstructions := roboBuild.String()
	roboVisited := visitHouses(roboInstructions)

	outputMap := make(map[Coord]int)
	maps.Copy(outputMap, santaVisited)

	for key, val := range roboVisited {
		if existingVal, ok := outputMap[key]; ok {
			outputMap[key] = existingVal + val
		} else {
			outputMap[key] = val
		}
	}

	return len(outputMap)
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)

	result1 := evalPart1(body)
	fmt.Println("Part 1:", result1)

	result2 := evalPart2(body)
	fmt.Println("Part 2:", result2)
}
