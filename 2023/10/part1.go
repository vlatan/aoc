package day10

import (
	"fmt"
)

// https://adventofcode.com/2023/day/10
func Part1() {
	_, startNode := parseFile("10/input.txt")
	loop := findLoop(startNode)
	fmt.Println((len(loop) + 1) / 2)
}
