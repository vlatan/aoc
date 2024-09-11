package day11

import "fmt"

// https://adventofcode.com/2023/day/11
func Part1() {
	matrix := parseFile("11/input.txt")
	galaxies := getGalaxies(matrix)
	ignore, result := map[P]struct{}{}, 0
	for _, node := range galaxies {
		result += shortestPathsSum(node, ignore, matrix)
		ignore[node] = struct{}{}
	}
	fmt.Println(result)
}

// Sum of steps to shortest paths from start to all the other '#'
func shortestPathsSum(start P, ignore map[P]struct{}, matrix Matrix) (result int) {
	visited, queue := map[P]int{start: 0}, []P{start}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		for _, node := range getNeighbours(current, matrix) {
			if _, ok := visited[node]; ok {
				continue
			}
			queue = append(queue, node)
			visited[node] = visited[current] + 1
			if matrix[node.x][node.y] == '#' {
				if _, ok := ignore[node]; !ok {
					result += visited[current] + 1
				}
			}
		}
	}
	return result
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
