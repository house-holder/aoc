package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

type Line struct {
	start, end Coord
	pos        int
	isVert     bool
}

type Edge struct {
	x1, y1, x2, y2 int
}

func getCoords(input string) Coord {
	strCoord := strings.Split(input, ",")
	valX, err := strconv.Atoi(strCoord[0])
	if err != nil {
		log.Fatalf("strconv.Atoi failed. %v", err)
	}
	valY, err := strconv.Atoi(strCoord[1])
	if err != nil {
		log.Fatalf("strconv.Atoi failed. %v", err)
	}

	return Coord{
		X: valX,
		Y: valY,
	}
}

func convertCoords(input string) []Coord {
	output := []Coord{}
	for line := range strings.SplitSeq(input, "\n") {
		if len(line) > 0 {
			newCoords := getCoords(line)
			output = append(output, newCoords)
		}
	}
	return output
}

func findArea(topL Coord, botR Coord) int {
	width := 1 + botR.X - topL.X
	height := 1 + botR.Y - topL.Y
	return width * height
}

func buildEdges(coords []Coord) []Edge {
	edges := make([]Edge, 0, len(coords))
	for i := 0; i < len(coords)-1; i++ {
		edges = append(edges, Edge{
			coords[i].X, coords[i].Y,
			coords[i+1].X, coords[i+1].Y,
		})
	}
	edges = append(edges, Edge{
		coords[len(coords)-1].X, coords[len(coords)-1].Y,
		coords[0].X, coords[0].Y,
	})
	return edges
}

func crossesBoundary(minX, minY, maxX, maxY int, edges []Edge) bool {
	for _, e := range edges {
		eMinX, eMaxX := e.x1, e.x2
		if eMinX > eMaxX {
			eMinX, eMaxX = eMaxX, eMinX
		}
		eMinY, eMaxY := e.y1, e.y2
		if eMinY > eMaxY {
			eMinY, eMaxY = eMaxY, eMinY
		}

		if minX < eMaxX && maxX > eMinX && minY < eMaxY && maxY > eMinY {
			return true
		}
	}
	return false
}

func evalPart2(coords []Coord) int {
	maxArea := 0
	edges := buildEdges(coords)

	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			a, b := coords[i], coords[j]

			minX, maxX := a.X, b.X
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			minY, maxY := a.Y, b.Y
			if minY > maxY {
				minY, maxY = maxY, minY
			}

			if crossesBoundary(minX, minY, maxX, maxY, edges) {
				continue
			}

			area := (maxX - minX + 1) * (maxY - minY + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func evalPart1(coords []Coord) int {
	maxArea := 0
	for i, cornerA := range coords {
		for j, cornerB := range coords {
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

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	coords := convertCoords(body)

	result1 := evalPart1(coords)
	result2 := evalPart2(coords)
	fmt.Println()
	fmt.Printf("Part 1: %d\n", result1)
	fmt.Printf("Part 2: %d\n", result2)
}
