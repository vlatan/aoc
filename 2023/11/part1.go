package day11

import "fmt"

// https://adventofcode.com/2023/day/11
func Part1() {
	matrix := parseFile("11/input.txt")
	galaxies := galaxies(matrix)
	ignore, result := map[P]struct{}{}, 0
	for _, node := range galaxies {
		current := shortestPaths(node, ignore, matrix)
		for _, num := range current {
			result += num
		}
		ignore[node] = struct{}{}
	}
	fmt.Println(result)
}

func shortestPaths(start P, ignore map[P]struct{}, matrix Matrix) map[P]int {
	distance := map[P]int{start: 0}
	galaxies := map[P]int{start: 0}

	queue := []P{start}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		neighbours := getNeighbours(current, matrix)
		for _, node := range neighbours {
			if _, ok := distance[node]; !ok {
				queue = append(queue, node)
				distance[node] = distance[current] + 1
				if matrix[node.x][node.y] == '#' {
					if _, ok := ignore[node]; !ok {
						galaxies[node] = distance[current] + 1
					}
				}
			}
		}
	}
	return galaxies
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
