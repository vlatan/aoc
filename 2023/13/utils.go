package day13

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"slices"
)

type Pattern [][]byte

// Parse file to a list of patterns
func parseFile(path string) (r, t []Pattern) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := Pattern{}
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			pattern = append(pattern, []byte(line))
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
		row := []byte{}
		for x := 0; x < len(p); x++ {
			row = append(row, p[x][y])
		}
		result[y] = row
	}
	return result
}

// Find the line of reflection.
// Return the prev number of rows.
// Return zero if no true line of reflection.
// If there is result to ignore, ignore it.
func (p Pattern) Process(delta, ignore int) int {
	prevLine := []byte{}
	for i, line := range p {
		if slices.Equal(line, prevLine) {
			r := p.Reflect(i, delta)
			if r != 0 && r != ignore {
				return r
			}
		}
		prevLine = line
	}
	return 0
}

// Check if a given line of reflection (start) is valid.
// If valid reflection return the num of prev rows.
// If invalid reflection return zero.
func (p Pattern) Reflect(start, delta int) int {
	for i, j := start, 1; i < len(p); i, j = i+1, j+1 {
		if start-j < 0 {
			break
		}
		if !slices.Equal(p[i], p[start-j]) {
			return 0
		}
	}
	return len(p[:start]) * delta
}

// Return the number of rows before a line of reflection.
// Ignore result if ignore >= 0
func calcRows(p, t Pattern, ignore int) int {
	if rows := p.Process(100, ignore); rows != 0 {
		return rows
	}
	return t.Process(1, ignore)
}
