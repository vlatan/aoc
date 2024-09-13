package day11

import (
	"aoc/2023/utils"
	"bufio"
	"os"
)

type Matrix [][]byte

type P struct{ x, y int }

type Space map[int]bool

func parseFile(path string) (Matrix, []P, Space, Space) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	matrix, galaxies := Matrix{}, []P{}
	eRows, eCols := Space{}, Space{}

	scanner := bufio.NewScanner(file)
	for x := 0; scanner.Scan(); x++ {
		line := []byte(scanner.Text())
		matrix = append(matrix, line)
		emptyRow := true
		for y := 0; y < len(line); y++ {
			if line[y] == '#' {
				emptyRow = false
				galaxies = append(galaxies, P{x, y})
			}
		}
		if emptyRow {
			eRows[x] = true
		}
	}

	for y := 0; y < len(matrix[0]); y++ {
		emptyColumn := true
		for x := 0; x < len(matrix); x++ {
			if matrix[x][y] == '#' {
				emptyColumn = false
				break
			}
		}
		if emptyColumn {
			eCols[y] = true
		}
	}
	return matrix, galaxies, eRows, eCols
}

// Sum of steps to shortest paths between
// all the pairs of galaxies in the universe
func shortestPathsSum(expansion int) (result int) {

	matrix, galaxies, eRows, eCols := parseFile("11/input.txt")
	done := map[P]struct{}{}

	dijkstra := func(start P) (currentResult int) {
		visited, queue := map[P]int{start: 0}, []P{start}
		for len(queue) != 0 {
			current := queue[0]
			queue = queue[1:]
			for _, node := range getNeighbours(current, matrix) {
				if _, ok := visited[node]; ok {
					continue
				}
				step := 1
				if eRows[current.x] || eCols[current.y] {
					step = expansion
				}
				queue = append(queue, node)
				visited[node] = visited[current] + step
				if matrix[node.x][node.y] == '#' {
					if _, ok := done[node]; !ok {
						currentResult += visited[node]
					}
				}
			}
		}
		return
	}

	for _, node := range galaxies {
		result += dijkstra(node)
		done[node] = struct{}{}
	}
	return

}

// Get neighbours of a coordinate on the matrix
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
	return
}
