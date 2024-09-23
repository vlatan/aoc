package day14

import "fmt"

// https://adventofcode.com/2023/day/14
func Part1() {
	matrix := parseFile("14/input.txt")
	matrix.North()
	fmt.Println(matrix.Count())
}
