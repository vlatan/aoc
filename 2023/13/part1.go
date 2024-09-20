package day13

import (
	"fmt"
)

// https://adventofcode.com/2023/day/13
func Part1() {
	result := 0
	data, transposed := parseFile("13/input.txt")
	for i := 0; i < len(data); i++ {
		result += calcRows(data[i], transposed[i], -1)
	}
	fmt.Println(result)
}
