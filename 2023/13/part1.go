package day13

import "fmt"

// https://adventofcode.com/2023/day/13
func Part1() {
	data, result := parseFile("13/input.txt"), 0
	for _, p := range data {
		if rows := p.Process(100); rows != 0 {
			result += rows
			continue
		}
		result += p.Transpose().Process(1)
	}
	fmt.Println(result)
}

// Find the line of symetry if any and return the prev number of rows
func (p Pattern) Process(delta int) int {
	prevLine, result := "", 0
	for i, line := range p {
		if line == prevLine {
			if result := p.Symmetry(i, delta); result != 0 {
				return result
			}
		}
		prevLine = p[i]
	}
	return result
}

func (p Pattern) Symmetry(start, delta int) int {
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
