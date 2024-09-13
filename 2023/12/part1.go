package day12

import "fmt"

// Every '?' can be '.' or '#'.
// Determine the indexes of the '?' in the string
// Produce the combinations of them being populated with '.' or/and '#'
// The number of the combinations should be 2^n, whene n is the number of '?'
// Try every combination and check if the string satisfies the contiguous group criteria.
// If so count the combination as a valid arrangement.
// Take the sum of all posible arangements.

// https://adventofcode.com/2023/day/12
func Part1() {
	data := parseFile("12/input.txt")
	fmt.Println(data)
}
