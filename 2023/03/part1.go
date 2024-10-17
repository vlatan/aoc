package day03

import (
	"aoc/2023/utils"
	"fmt"
)

// https://adventofcode.com/2023/day/3
func Part1() {
	result, matrix := 0, utils.ParseFile("03/input.txt")
	// iterate over the two-dimensional array
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if utils.IsDigit(matrix[x][y]) {
				num := number{string(matrix[x][y]), pos{x, y}, pos{x, y}}
				for i := y + 1; i < len(matrix[x]); i++ {
					if !utils.IsDigit(matrix[x][i]) {
						break
					}
					num.value += string(matrix[x][i])
					num.end = pos{x, i}
					y = i
				}
				// check here if num is valid and note its neighbouring stars
				if _, ok := inspectAroundNumber(num, matrix); ok {
					result += utils.ToInteger(num.value)
				}
			}
		}
	}
	fmt.Println("Part 1:", result)
}
