package day13

import "fmt"

// https://adventofcode.com/2023/day/13
func Part1() {
	result := 0
	data, transposed := parseFile("13/input.txt")
	for i := 0; i < len(data); i++ {
		result += calcRows(data[i], transposed[i], -1)
	}
	fmt.Println(result)
}

// Return the number of rows before a line of reflection.
func calcRows(p, t Pattern, ignore int) int {
	if rows := p.Process(100, ignore); rows != 0 {
		return rows
	}
	return t.Process(1, ignore)
}

// Find the line of reflection.
// Return the prev number of rows.
// Return zero if no true line of reflection.
func (p Pattern) Process(delta, ignore int) int {
	prevLine := ""
	for i, line := range p {
		if line == prevLine && i != ignore {
			if r := p.Reflect(i, delta); r != 0 {
				return r
			}
		}
		prevLine = p[i]
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
		if p[i] != p[start-j] {
			return 0
		}
	}
	return len(p[:start]) * delta
}
