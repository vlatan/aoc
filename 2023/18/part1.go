package day18

import (
	"aoc/2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type P struct{ x, y int }
type Graph map[P]struct{}

// https://adventofcode.com/2023/day/18
func Part1() {
	file, err := os.Open("18/input.txt")
	utils.Check(err)
	defer file.Close()

	graph := make(Graph)
	xMax, xMin, yMax, yMin := 0, 0, 0, 0

	current := P{0, 0}
	graph[current] = struct{}{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		steps := utils.ToInteger(fields[1])
		x, y := current.x, current.y

		switch fields[0] {
		case "U":
			for range steps {
				x++
				current = P{x, y}
				graph[current] = struct{}{}
			}
			xMax = max(x, xMax)
		case "D":
			for range steps {
				x--
				current = P{x, y}
				graph[current] = struct{}{}
			}
			xMin = min(x, xMin)
		case "L":
			for range steps {
				y--
				current = P{x, y}
				graph[current] = struct{}{}
			}
			yMin = min(y, yMin)
		case "R":
			for range steps {
				y++
				current = P{x, y}
				graph[current] = struct{}{}
			}
			yMax = max(y, yMax)
		}
	}

	result := 0
	// cast rays diagonally
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			// count if point is on the boundary
			if _, ok := graph[P{x, y}]; ok {
				result++
				continue
			}

			count := 0
			// count ray interections with boundary
			for i := 1; x+i <= xMax && y+i <= yMax; i++ {
				xi, yi := x+i, y+i

				// check if it's intersection with the boundary
				if _, ok := graph[P{xi, yi}]; !ok {
					continue
				}

				// check left/down or up/right neighbours to see
				// if the ray is just grazing a corner on the outside
				// do not count as intersection
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
			// the point is inside the border
			if count%2 != 0 {
				result++
			}
		}
	}

	fmt.Println(result)

	// for x := xMax; x >= xMin; x-- {
	// 	for y := yMin; y <= yMax; y++ {
	// 		if _, ok := graph[P{x, y}]; ok {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}
