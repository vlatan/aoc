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
			// check above and/or left of the current node
			switch symbol {
			case '|', 'L':
				Check(node, Loc{x - 1, y}, graph, "|7F")
			case '-', '7':
				Check(node, Loc{x, y - 1}, graph, "-LF")
			case 'J':
				Check(node, Loc{x - 1, y}, graph, "|7F")
				Check(node, Loc{x, y - 1}, graph, "-LF")
			case 'S':
				start = Loc{x, y}
			}
		}
	}
	return start, graph
}

// If neighbour found connect the two nodes
func Check(node *Node, target Loc, graph Graph, symbols string) {
	if nn, ok := graph[target]; ok {
		if strings.ContainsRune(symbols, nn.symbol) {
			node.neighbors = append(node.neighbors, nn)
			nn.neighbors = append(nn.neighbors, node)
		}
	}
}
