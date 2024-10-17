package day12

import (
	"fmt"
)

// https://adventofcode.com/2023/day/12
func Part1() {
	lines, groups := parseFile("12/input.txt")
	result := 0
	for i := 0; i < len(lines); i++ {
		result += solve(lines[i], groups[i], Cache{})
	}
	fmt.Println("Part 1:", result)
}
