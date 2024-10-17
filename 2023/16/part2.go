package day16

import "fmt"

type Start struct {
	comingFrom string
	loc        Loc
}

// https://adventofcode.com/2023/day/16
func Part2() {
	matrix := parseFile("16/input.txt")
	startingPoints := findStarts(matrix)

	result := 0
	for _, start := range startingPoints {
		currentResult := solve(matrix, start.loc, start.comingFrom)
		if currentResult > result {
			result = currentResult
		}
	}

	fmt.Println("Part 2:", result)
}

// Find starting points on the matrix
func findStarts(matrix Matrix) (r []Start) {

	for i := range matrix[0] {
		r = append(r, Start{"up", Loc{0, i}})
		r = append(r, Start{"down", Loc{len(matrix[0]) - 1, i}})
	}

	for i := range len(matrix) {
		r = append(r, Start{"left", Loc{i, 0}})
		r = append(r, Start{"right", Loc{i, len(matrix[0]) - 1}})
	}

	return
}
