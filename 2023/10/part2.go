package day10

import "fmt"

// https://adventofcode.com/2023/day/10
func Part2() {
	matrix, startLoc, graph := parseFile("10/input.txt")
	loop := findLoop(graph[startLoc])
	result := castRays(matrix, loop)
	fmt.Println(result)
}

// Count points that are enclosed by the loop
// Cast a reay diagonally for every eligible point on the matrix
func castRays(matrix []string, loop []*Node) (result int) {
	for x, line := range matrix {
		for y := range line {
			if _, onloop, _ := onLoop(Loc{x, y}, loop); onloop {
				continue
			}
			count := 0 // shoot the ray diagonally
			for i := 1; x+i < len(matrix) && y+i < len(matrix[0]); i++ {
				// check if the point is on the loop
				node, onloop, oncorner := onLoop(Loc{x + i, y + i}, loop)
				// point not on the loop, nothing to count
				if !onloop {
					continue
				}
				// point is on the loop but not on a corner of the loop
				// count the intersection and do nothing else
				if !oncorner {
					count++
					continue
				}
				// This is a corner. Check if the ray is grazing on the outside of the corner.
				// If the ray is not grazing on the outside of the corner count the intersection.
				n1, n2 := node.neighbors[0], node.neighbors[1]
				left, down := Loc{x + i, y + i - 1}, Loc{x + i + 1, y + i}
				up, right := Loc{x + i - 1, y + i}, Loc{x + i, y + i + 1}

				if (n1.loc == left || n1.loc == down) &&
					(n2.loc == left || n2.loc == down) {
					continue
				}
				if (n1.loc == up || n1.loc == right) &&
					(n2.loc == up || n2.loc == right) {
					continue
				}
				count++
			}
			// Odd interesections means the point is enclosed by the loop
			if count%2 != 0 {
				result++
			}
		}
	}
	return result
}

// Check if on the loop and if so if on corner
func onLoop(loc Loc, loop []*Node) (node *Node, onloop bool, oncorner bool) {
	for i, node := range loop {
		if node.loc != loc {
			continue
		}
		prevIndex := i - 1
		if prevIndex < 0 {
			prevIndex = len(loop) - 1
		}
		prev, next := loop[prevIndex], loop[(i+1)%len(loop)]
		dx := abs(prev.loc.x - next.loc.x)
		dy := abs(prev.loc.y - next.loc.y)
		if dx == 1 && dy == 1 {
			return node, true, true
		}
		return node, true, false
	}
	return node, onloop, oncorner
}

// Absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
