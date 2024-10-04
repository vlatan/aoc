package day16

import "fmt"

// This is a directed graph
// Find the terminal case for this graph.

type Loc struct {
	x, y int
}

// https://adventofcode.com/2023/day/16
func Part1() {
	matrix := parseFile("16/input.txt")
	result := solve(matrix, Loc{0, 0}, "left")
	fmt.Println(result)
}
