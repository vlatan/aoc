package day10

import "fmt"

// https://adventofcode.com/2023/day/10
func Part2() {
	matrix, startNode := parseFile("10/input.txt")
	loop := findLoop(startNode)
	result := castRays(matrix, loop)
	fmt.Println("Part 2:", result)
}

// Count points enclosed by the loop
// Cast a ray diagonally for every eligible point on the matrix
func castRays(matrix []string, loop Graph) (result int) {
	for x, line := range matrix {
		for y := range line {
			if _, ok := loop[Loc{x, y}]; ok {
				continue
			}
			count := 0 // cast the ray diagonally and count the intersections
			for i := 1; x+i < len(matrix) && y+i < len(matrix[0]); i++ {
				xi, yi := x+i, y+i
				// check if the point is on the loop
				node, ok := loop[Loc{xi, yi}]
				// point not on the loop, no intersection
				if !ok {
					continue
				}
				n1, n2 := node.neighbors[0], node.neighbors[1]
				// Point is on the loop but NOT on the corner of it.
				// Count the intersection and do nothing else.
				dx, dy := n1.loc.x-n2.loc.x, n1.loc.y-n2.loc.y
				if abs(dx) != 1 && abs(dy) != 1 {
					count++
					continue
				}
				// Check if the ray is grazing the corner on the outside.
				// If so do NOT count the intersection.
				left, down := Loc{xi, yi - 1}, Loc{xi + 1, yi}
				up, right := Loc{xi - 1, yi}, Loc{xi, yi + 1}
				if (n1.loc == left && n2.loc == down) ||
					(n2.loc == left && n1.loc == down) ||
					(n1.loc == up && n2.loc == right) ||
					(n2.loc == up && n1.loc == right) {
					continue
				}
				count++
			}
			// Odd number of interesections means
			// the point is inside the loop
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
