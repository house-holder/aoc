package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
)

type Network struct {
	devices map[string][]string
	cache   map[string]map[string]int
	cycles  int
}

func NewNetwork(lines []string) Network {
	d := make(map[string][]string)
	for _, ln := range lines {
		parts := strings.Split(ln, ":")
		key := parts[0]
		val := strings.Split(strings.Trim(parts[1], " "), " ")
		d[key] = val
	}
	return Network{
		devices: d,
		cache:   make(map[string]map[string]int),
		cycles:  0,
	}
}

func (n *Network) cacheDump() {
	nodes := make([]string, 0, len(n.cache))
	for node := range n.cache {
		nodes = append(nodes, node)
	}
	sort.Strings(nodes)

	for _, node := range nodes {
		fmt.Printf("> %s\n", node)

		stopKeys := make([]string, 0, len(n.cache[node]))
		for key := range n.cache[node] {
			stopKeys = append(stopKeys, key)
		}
		sort.Strings(stopKeys)

		for _, stopKey := range stopKeys {
			result := n.cache[node][stopKey]
			stopsDisplay := stopKey
			if stopsDisplay == "" {
				stopsDisplay = "-"
			}
			fmt.Printf("  %-9s %d paths\n", stopsDisplay, result)
		}
	}
}

func (n *Network) anyPath(curr string, tgt string) int {
	if curr == tgt {
		return 1
	}
	acc := 0

	for _, each := range n.devices[curr] {
		acc += n.anyPath(each, tgt)
	}

	return acc
}

func (n *Network) pathStops(
	curr string,
	tgt string,
	stops []string,
	path []string,
	visited map[string]bool,
) int {
	if visited[curr] {
		return 0
	}
	visitedKey := ""
	for _, stop := range stops {
		for _, p := range path {
			if p == stop {
				visitedKey += stop + ","
			}
		}
	}

	if n.cache[curr] != nil {
		if cached, ok := n.cache[curr][visitedKey]; ok {
			// fmt.Printf("! node=%s stops=%q result=%d\n", curr, visitedKey, cached)
			return cached
		}
	}

	// fmt.Printf("node=%s stops=%q d=%d\n", curr, visitedKey, len(path))
	n.cycles++

	visited[curr] = true
	newPath := append([]string{}, path...)
	newPath = append(path, curr)

	// if len(newPath)%10 == 0 {
	// 	fmt.Printf("Depth: %d, Current: %s\n", len(newPath), curr)
	// }

	if curr == tgt {
		result := 0
		if allVisited(newPath, stops) {
			// fmt.Println(">>> Found valid path")
			result = 1
		}
		if n.cache[curr] == nil {
			n.cache[curr] = make(map[string]int)
		}
		n.cache[curr][visitedKey] = result
		return result
	}

	acc := 0
	for _, cnxns := range n.devices[curr] {
		newVisited := make(map[string]bool, len(visited))
		for each := range visited {
			newVisited[each] = true
		}
		acc += n.pathStops(cnxns, tgt, stops, newPath, newVisited)
	}

	if n.cache[curr] == nil {
		n.cache[curr] = make(map[string]int)
	}
	n.cache[curr][visitedKey] = acc

	return acc
}

func allVisited(path []string, stops []string) bool {
	for _, stop := range stops {
		if !slices.Contains(path, stop) {
			return false
		}
	}
	return true
}

func evalPart1(filename string) int {
	lines := readAOC(filename, 1)
	n := NewNetwork(lines)

	return n.anyPath("you", "out")
}

func evalPart2(filename string) int {
	lines := readAOC(filename, 2)
	n := NewNetwork(lines)

	path := []string{}
	stops := []string{"dac", "fft"}
	visited := make(map[string]bool)
	n.cache = make(map[string]map[string]int)

	result := n.pathStops("svr", "out", stops, path, visited)
	n.cacheDump()
	fmt.Printf("(Part 2 Cycles: %d)\n", n.cycles)
	return result
}

func main() {
	result1 := evalPart1(os.Args[1])
	fmt.Println("Part 1:", result1)

	result2 := evalPart2(os.Args[1])
	fmt.Println("Part 2:", result2)
}

func readAOC(filename string, part int) []string {
	fn := "input.txt"
	if filename == "example.txt" {
		fn = fmt.Sprintf("example%d.txt", part)
	}
	bytes, err := os.ReadFile(fn)
	if err != nil {
		log.Fatalf("os.ReadFile failed. %v", err)
	}
	body := string(bytes)
	return strings.Split(strings.Trim(body, "\n"), "\n")
}
