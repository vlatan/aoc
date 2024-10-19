package day21

import (
	"fmt"
)

type Points map[P]struct{}

// https://adventofcode.com/2023/day/21
func Part1() {
	start, matrix := parseFile("21/input.txt")
	points := Points{start: struct{}{}}
	for range 64 {
		points = walk(points, matrix)
	}
	fmt.Println("Part 1:", len(points))
}

func walk(points Points, matrix M) Points {
	r := make(Points)
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
				r[p] = struct{}{}
			}
		}
	}
	return r
}
