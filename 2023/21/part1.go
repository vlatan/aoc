package day21

import (
	"fmt"
)

type Points map[P]struct{}

// https://adventofcode.com/2023/day/21
func Part1() {
	start, matrix := parseFile("21/input.txt")
	result := walk(start, matrix, 64)
	fmt.Println("Part 1:", result)
}

func walk(start P, matrix M, steps int) int {
	points := Points{start: struct{}{}}
	for range steps {
		currentPoints := make(Points)
		for point := range points {
			// steps around
			around := []P{
				{point.x, point.y - 1},
				{point.x, point.y + 1},
				{point.x - 1, point.y},
				{point.x + 1, point.y},
			}

			// check around
			for _, p := range around {
				// check if a garden plot
				if value, ok := matrix[p]; ok && value == '.' {
					currentPoints[p] = struct{}{}
				}
			}
		}
		points = currentPoints
	}
	return len(points)
}
