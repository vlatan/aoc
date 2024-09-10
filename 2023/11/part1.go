package day11

import "fmt"

// https://adventofcode.com/2023/day/11
func Part1() {
	matrix := parseFile("11/input.txt")
	pairs := galaxyPairs(matrix)

	// for _, line := range matrix {
	// 	fmt.Println(string(line))
	// }

	fmt.Println(pairs)

}
