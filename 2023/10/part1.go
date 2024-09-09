package day10

import (
	"fmt"
)

// https://adventofcode.com/2023/day/10
func Part1() {
	_, startLoc, graph := parseFile("10/input.txt")
	loop := findLoop(graph[startLoc])
	fmt.Println((len(loop) + 1) / 2)
}
