package day13

import (
	"aoc/2023/utils"
	"bufio"
	"os"
)

type Pattern []string

// Parse file to a list of patterns
func parseFile(path string) (r, t []Pattern) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := Pattern{}
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			pattern = append(pattern, line)
			continue
		}
		r = append(r, pattern)
		t = append(t, pattern.Transpose())
		pattern = Pattern{}
	}
	r = append(r, pattern)
	t = append(t, pattern.Transpose())
	return
}

// Transpose a matrix
func (p Pattern) Transpose() Pattern {
	result := make(Pattern, len(p[0]))
	for y := 0; y < len(p[0]); y++ {
		row := ""
		for x := 0; x < len(p); x++ {
			row += string(p[x][y])
		}
		result[y] = row
	}
	return result
}
