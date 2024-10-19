package day21

import (
	"aoc/2023/common"
	"bufio"
	"os"
)

type P struct{ x, y int }
type M map[P]byte

func parseFile(path string) (start P, matrix M) {
	file, err := os.Open(path)
	common.Check(err)
	defer file.Close()
	scanner, matrix := bufio.NewScanner(file), M{}
	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		for c := 0; c < len(line); c++ {
			if line[c] == 'S' {
				start = P{r, c}
				matrix[start] = '.'
				continue
			}
			matrix[P{r, c}] = line[c]
		}
	}
	return
}
