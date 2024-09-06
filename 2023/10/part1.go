package day10

import (
	"fmt"
	"slices"
)

// https://adventofcode.com/2023/day/10
func Part1() {
	startLoc, graph := parseFile("10/input.txt")
	loop := findLoop(graph[startLoc])
	fmt.Println((len(loop) + 1) / 2)
}

// Returns a slice of nodes if loop possible.
// If not returns an empty slice.
func findLoop(start *Node) []*Node {

	if len(start.neighbors) != 2 {
		return []*Node{}
	}

	visited := []*Node{start}
	next, end := start.neighbors[0], start.neighbors[1]
	var recurse func(start *Node) []*Node

	recurse = func(start *Node) []*Node {
		if start == end {
			return []*Node{start}
		}
		visited = append(visited, start)
		for _, node := range start.neighbors {
			if !slices.Contains(visited, node) {
				path := recurse(node)
				if len(path) > 0 {
					return append(path, start)
				}
				return path
			}
		}
		return []*Node{}
	}

	result := recurse(next)
	if len(result) > 0 {
		result = append(result, start)
		slices.Reverse(result)
	}
	return result
}
