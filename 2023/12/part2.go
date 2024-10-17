package day12

import "fmt"

// https://adventofcode.com/2023/day/12
func Part2() {
	lines, groups := parseFile("12/input.txt")
	for i := 0; i < len(lines); i++ {
		line, group := lines[i], groups[i]
		for range 4 {
			lines[i] += "?" + line
			groups[i] = append(groups[i], group...)
		}
	}

	result := 0
	for i := 0; i < len(lines); i++ {
		result += solve(lines[i], groups[i], Cache{})
	}
	fmt.Println("Part 2:", result)
}
