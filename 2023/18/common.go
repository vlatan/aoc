package day18

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

type P struct{ x, y int }
type Graph map[P]struct{}
type BoundingBox struct{ xMin, yMin, xMax, yMax int }
type processLine func([]string) (string, uint64)

func parseFile(path string, fn processLine) (Graph, BoundingBox) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	graph, b := make(Graph), BoundingBox{}
	current := P{0, 0}
	graph[current] = struct{}{}

	// TODO: Work with just the edges of the polygon, not every point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		direction, steps := fn(fields)
		x, y := current.x, current.y

		switch direction {
		case "U":
			for range steps {
				x++
				current = P{x, y}
				graph[current] = struct{}{}
			}
			b.xMax = max(x, b.xMax)
		case "D":
			for range steps {
				x--
				current = P{x, y}
				graph[current] = struct{}{}
			}
			b.xMin = min(x, b.xMin)
		case "L":
			for range steps {
				y--
				current = P{x, y}
				graph[current] = struct{}{}
			}
			b.yMin = min(y, b.yMin)
		case "R":
			for range steps {
				y++
				current = P{x, y}
				graph[current] = struct{}{}
			}
			b.yMax = max(y, b.yMax)
		}
	}
	return graph, b
}

func (p P) castRay(graph Graph, b BoundingBox) int {
	// the point is on the polygon
	if _, ok := graph[p]; ok {
		return 1
	}

	count := 0
	// count diagonal ray interections with the polygon
	for i := 1; p.x+i <= b.xMax && p.y+i <= b.yMax; i++ {
		xi, yi := p.x+i, p.y+i

		// check if it's NOT an intersection with the polygon
		if _, ok := graph[P{xi, yi}]; !ok {
			continue
		}

		// Check left/down or up/right neighbours to see
		// if the ray is just grazing a corner on the outside.
		// If so, do not count as an intersection.
		_, left := graph[P{xi, yi - 1}]
		_, down := graph[P{xi + 1, yi}]
		_, right := graph[P{xi, yi + 1}]
		_, up := graph[P{xi - 1, yi}]
		if (left && down) || (up && right) {
			continue
		}

		count++
	}
	// Odd number of interesections means
	// the point is inside inside the polygon
	if count%2 != 0 {
		return 1
	}

	return 0
}
