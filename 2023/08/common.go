package day08

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

func parseFile(path string) (string, map[string][]string) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	commands := scanner.Text()

	graph := map[string][]string{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			continue
		}
		root, nodes, _ := strings.Cut(line, " = ")
		n1, n2, _ := strings.Cut(nodes[1:len(nodes)-1], ", ")
		graph[root] = []string{n1, n2}
	}

	return commands, graph
}
