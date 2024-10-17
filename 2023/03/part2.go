package day03

import (
	"aoc/2023/utils"
	"fmt"
)

// https://adventofcode.com/2023/day/3
func Part2() {
	stars := make(map[pos][]string)
	matrix := utils.ParseFile("03/input.txt")
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
				if starNeighbours, ok := inspectAroundNumber(num, matrix); ok {
					for pos, number := range starNeighbours {
						stars[pos] = append(stars[pos], number)
					}
				}
			}
		}
	}
	result := 0
	for _, v := range stars {
		if len(v) == 2 {
			result += utils.ToInteger(v[0]) * utils.ToInteger(v[1])
		}
	}
	fmt.Println("Part 2:", result)
}
