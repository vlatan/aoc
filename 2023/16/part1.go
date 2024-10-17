package day16

import "fmt"

// https://adventofcode.com/2023/day/16
func Part1() {
	matrix := parseFile("16/input.txt")
	result := solve(matrix, Loc{0, 0}, "left")
	fmt.Println("Part 1:", result)
}
