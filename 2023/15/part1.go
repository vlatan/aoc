package day15

import "fmt"

// https://adventofcode.com/2023/day/15
func Part1() {
	data := parseFile("15/input.txt")
	sum := 0
	for _, s := range data {
		sum += hash(s)
	}
	fmt.Println("Part 1:", sum)
}
