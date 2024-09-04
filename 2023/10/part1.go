package day10

import "fmt"

// https://adventofcode.com/2023/day/10
func Part1() {
	start, graph := parseFile("10/input.txt")
	fmt.Println(start)
	fmt.Println()
	for loc, v := range graph {
		fmt.Println(loc, v)
	}
}
