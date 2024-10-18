package day21

import (
	"aoc/2023/common"
	"bufio"
	"os"
)

type P struct{ x, y int }
type M [][]byte

func parseFile(path string) (start P, matrix M) {
	file, err := os.Open(path)
	common.Check(err)
	defer file.Close()
	scanner, matrix := bufio.NewScanner(file), M{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		row := make([]byte, len(line))
		for j := 0; j < len(line); j++ {
			if line[j] == 'S' {
				start = P{i, j}
				row[j] = '.'
				continue
			}
			row[j] = line[j]
		}
		matrix = append(matrix, row)
	}
	return
}
