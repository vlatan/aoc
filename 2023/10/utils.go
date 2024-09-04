package day10

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

type Loc struct {
	x, y int
}

type Node struct {
	symbol    rune
	neighbors []*Node
}

type Graph map[Loc]*Node

func parseFile(path string) (Loc, Graph) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	start, graph := Loc{}, Graph{}
	for x := 0; scanner.Scan(); x++ {
		line := scanner.Text()
		for y, symbol := range line {
			node := &Node{symbol, []*Node{}}
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
				start = Loc{x, y}
			}
		}
	}
	// resolve the neighbours of "S" the starting point
	node, x, y := graph[start], start.x, start.y
	left, right := Loc{x, y - 1}, Loc{x, y + 1}
	up, down := Loc{x - 1, y}, Loc{x + 1, y}
	Connect(node, left, graph, "-LF")
	Connect(node, right, graph, "-J7")
	Connect(node, up, graph, "|7F")
	Connect(node, down, graph, "|JL")

	return start, graph
}

// If neighbour found connect the two nodes
func Connect(node *Node, target Loc, graph Graph, symbols string) {
	if nn, ok := graph[target]; ok {
		if strings.ContainsRune(symbols, nn.symbol) {
			node.neighbors = append(node.neighbors, nn)
			nn.neighbors = append(nn.neighbors, node)
		}
	}
}
