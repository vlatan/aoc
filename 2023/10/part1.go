package day10

import (
	"fmt"
	"slices"
	"strings"
)

// https://adventofcode.com/2023/day/10
func Part1() {
	start, graph := parseFile("10/input.txt")
	neighbours := startNeighbours(start, graph)
	loop := findLoop(neighbours[0], neighbours[1])
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
				path := recurse(node)
				return append(path, start)
			}
		}
		visited = visited[:len(visited)-2]
		return []*Node{}
	}

	return recurse(start)
}

func startNeighbours(start Loc, graph Graph) []*Node {
	valid, x, y := &[]*Node{}, start.x, start.y
	left, right := Loc{x, y - 1}, Loc{x, y + 1}
	up, down := Loc{x - 1, y}, Loc{x + 1, y}

	Check(valid, left, graph, "-LF")
	Check(valid, right, graph, "-J7")
	Check(valid, up, graph, "|7F")
	Check(valid, down, graph, "|JL")

	return *valid
}

func Check(valid *[]*Node, loc Loc, graph Graph, symbols string) {
	if n, ok := graph[loc]; ok {
		if strings.ContainsRune(symbols, n.symbol) {
			*valid = append(*valid, n)
		}
	}
}
