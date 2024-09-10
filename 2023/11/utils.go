package day11

import (
	"aoc/2023/utils"
	"bufio"
	"os"
)

type Matrix [][]byte

type P struct{ x, y int }

type Pair struct{ a, b P }

func parseFile(path string) Matrix {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	// add empty rows
	firstStage := Matrix{}
	scanner := bufio.NewScanner(file)
	for x := 0; scanner.Scan(); x++ {
		line := scanner.Text()
		byteArray, galaxyRow := make([]byte, len(line)), false
		for y := 0; y < len(line); y++ {
			byteArray[y] = line[y]
			if line[y] == '#' {
				galaxyRow = true
			}
		}
		firstStage = append(firstStage, byteArray)
		if !galaxyRow {
			byteArray := make([]byte, len(line))
			for i := 0; i < len(line); i++ {
				byteArray[i] = '.'
			}
			firstStage = append(firstStage, byteArray)
		}
	}

	// add empty columns
	matrix := make(Matrix, len(firstStage))
	for y := 0; y < len(firstStage[0]); y++ {
		galaxyColumn := false
		for x := 0; x < len(firstStage); x++ {
			matrix[x] = append(matrix[x], firstStage[x][y])
			if firstStage[x][y] == '#' {
				galaxyColumn = true
			}
		}
		if !galaxyColumn {
			for i := 0; i < len(firstStage); i++ {
				matrix[i] = append(matrix[i], '.')
			}
		}
	}

	return matrix
}

func galaxyPairs(matrix Matrix) (pairs []Pair) {
	galaxies := []P{}
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			if matrix[x][y] == '#' {
				galaxies = append(galaxies, P{x, y})
			}
		}
	}
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			pairs = append(pairs, Pair{g1, g2})
		}
	}
	return pairs
}
