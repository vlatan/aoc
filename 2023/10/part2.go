package day10

import "fmt"

// https://adventofcode.com/2023/day/10
func Part2() {
	matrix, startNode := parseFile("10/input.txt")
	loop := findLoop(startNode)
	result := castRays(matrix, loop)
	fmt.Println(result)
}

// Count points enclosed by the loop
// Cast a reay diagonally for every eligible point on the matrix
func castRays(matrix []string, loop Graph) (result int) {
	for x, line := range matrix {
		for y := range line {
			if _, ok := loop[Loc{x, y}]; ok {
				continue
			}
			count := 0 // cast the ray diagonally and count the intersections
			for i := 1; x+i < len(matrix) && y+i < len(matrix[0]); i++ {
				// check if the point is on the loop
				node, ok := loop[Loc{x + i, y + i}]
				// point not on the loop, no intersection
				if !ok {
					continue
				}
				n1, n2 := node.neighbors[0], node.neighbors[1]
				// Point is NOT on a corner of the loop.
				// Count the intersection and do nothing else.
				dx, dy := abs(n1.loc.x-n2.loc.x), abs(n1.loc.y-n2.loc.y)
				if !(dx == 1 && dy == 1) {
					count++
					continue
				}
				// Check if the ray is grazing on the outside of the corner.
				// If so do NOT count the intersection.
				left, down := Loc{x + i, y + i - 1}, Loc{x + i + 1, y + i}
				up, right := Loc{x + i - 1, y + i}, Loc{x + i, y + i + 1}
				c1 := (n1.loc == left && n2.loc == down) || (n2.loc == left && n1.loc == down)
				c2 := (n1.loc == up && n2.loc == right) || (n2.loc == up && n1.loc == right)
				if c1 || c2 {
					continue
				}
				count++
			}
			// Odd number of interesections means
			// the point is enclosed by the loop
			if count%2 != 0 {
				result++
			}
		}
	}
	return result
}

// Absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
