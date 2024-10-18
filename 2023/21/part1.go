package day21

import (
	"fmt"
)

type Points map[P]struct{}

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
		// step around
		around := []P{
			{point.x, point.y - 1},
			{point.x, point.y + 1},
			{point.x - 1, point.y},
			{point.x + 1, point.y},
		}

		for _, p := range around {
			// check the bounds
			if p.x < 0 || p.x > len(matrix)-1 ||
				p.y < 0 || p.y > len(matrix[0])-1 {
				continue
			}

			// check if a garden plot
			if matrix[p.x][p.y] == '.' {
				r[p] = struct{}{}
			}
		}
	}
	return r
}
