package day10

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"slices"
	"strings"
)

type Loc struct {
	x, y int
}

type Node struct {
	symbol    rune
	loc       Loc
	neighbors []*Node
}

type Graph map[Loc]*Node

// Create Graph from matrix
func parseFile(path string) ([]string, *Node) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix, start, graph := []string{}, &Node{}, Graph{}
	for x := 0; scanner.Scan(); x++ {
		line := scanner.Text()
		matrix = append(matrix, line)
		for y, symbol := range line {
			node := &Node{symbol, Loc{x, y}, []*Node{}}
			graph[Loc{x, y}] = node
			up, left := Loc{x - 1, y}, Loc{x, y - 1}
			switch symbol {
			case '|', 'L':
				Connect(node, up, graph, "|7F")
			case '-', '7':
				Connect(node, left, graph, "-LF")
			case 'J':
				Connect(node, up, graph, "|7F")
				Connect(node, left, graph, "-LF")
			case 'S':
				start = node
				Connect(node, up, graph, "|7F")
				Connect(node, left, graph, "-LF")
			}
		}
	}
	// finish resolving neighbours of "S"
	x, y := start.loc.x, start.loc.y
	right, down := Loc{x, y + 1}, Loc{x + 1, y}
	Connect(start, right, graph, "-J7")
	Connect(start, down, graph, "|JL")
	return matrix, start
}

// Connect two nodes if found to be neighbours
func Connect(node *Node, target Loc, graph Graph, symbols string) {
	if nn, ok := graph[target]; ok {
		if strings.ContainsRune(symbols, nn.symbol) {
			node.neighbors = append(node.neighbors, nn)
			nn.neighbors = append(nn.neighbors, node)
		}
	}
}

// Returns a slice of nodes if loop possible.
// If not returns an empty slice.
func findLoop(start *Node) []*Node {

	if len(start.neighbors) != 2 {
		return []*Node{}
	}

	visited := map[Loc]struct{}{start.loc: {}}
	next, end := start.neighbors[0], start.neighbors[1]
	var recurse func(start *Node) []*Node

	recurse = func(start *Node) []*Node {
		if start == end {
			return []*Node{start}
		}
		visited[start.loc] = struct{}{}
		for _, node := range start.neighbors {
			if _, ok := visited[node.loc]; !ok {
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
