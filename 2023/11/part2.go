package day11

import "fmt"

// https://adventofcode.com/2023/day/11
func Part2() {
	expansion := 1_000_000
	result := allShortestPathsSum(expansion)
	fmt.Println(result)
}
