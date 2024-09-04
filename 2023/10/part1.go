package day10

import (
	"fmt"
	"slices"
)

// https://adventofcode.com/2023/day/10
func Part1() {
	startLoc, graph := parseFile("10/input.txt")
	start, end := graph[startLoc].neighbors[0], graph[startLoc]
	loop := findLoop(start, end)
	fmt.Println((len(loop) + 1) / 2)
}

func findLoop(start *Node, end *Node) []*Node {
	visited := []*Node{}
	var recurse func(start *Node) []*Node
	recurse = func(start *Node) []*Node {
		if start == end {
			return []*Node{start}
		}
		visited = append(visited, start)
		for _, node := range start.neighbors {
			if !slices.Contains(visited, node) {
				return append(recurse(node), start)
			}
		}
		panic(fmt.Sprintf("There's no path from %v to %v", start, end))
	}
	return recurse(start)
}
