package day11

import "fmt"

// https://adventofcode.com/2023/day/11
func Part1() {
	matrix := parseFile("11/input.txt")
	// for _, line := range matrix {
	// 	fmt.Println(string(line))
	// }
	pairs := galaxyPairs(matrix)
	// fmt.Println(len(pairs))
	result := 0
	for _, pair := range pairs {
		result += shortestPath(pair.a, pair.b, matrix)
	}
	fmt.Println(result)
}

func shortestPath(start P, end P, matrix Matrix) int {
	distance := map[P]int{start: 0}
	queue := []P{start}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		neighbours := getNeighbours(current, matrix)
		for _, n := range neighbours {
			if _, ok := distance[n]; !ok {
				queue = append(queue, n)
				distance[n] = distance[current] + 1
			}
			if n == end {
				return distance[n]
			}
		}
	}
	return distance[start]
}

func getNeighbours(point P, matrix Matrix) (neighbours []P) {
	x, y := point.x, point.y
	left, right := P{x, y - 1}, P{x, y + 1}
	up, down := P{x - 1, y}, P{x + 1, y}

	if y-1 >= 0 {
		neighbours = append(neighbours, left)
	}
	if y+1 < len(matrix[0]) {
		neighbours = append(neighbours, right)
	}
	if x-1 >= 0 {
		neighbours = append(neighbours, up)
	}
	if x+1 < len(matrix) {
		neighbours = append(neighbours, down)
	}
	return neighbours
}
