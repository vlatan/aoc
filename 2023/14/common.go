package day14

import (
	"aoc/2023/common"
	"bufio"
	"os"
)

type Matrix [][]byte

func parseFile(path string) (result Matrix) {
	file, err := os.Open(path)
	common.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, []byte(scanner.Text()))
	}
	return
}

func (m Matrix) North() {
	for move := true; move; {
		move = false
		for x := 1; x < len(m); x++ {
			for y := 0; y < len(m[0]); y++ {
				if m[x][y] == 'O' && m[x-1][y] == '.' {
					m[x-1][y], m[x][y] = 'O', '.'
					move = true
				}
			}
		}
	}
}

func (m Matrix) Count() (sum int) {
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[0]); y++ {
			if m[x][y] == 'O' {
				sum += len(m) - x
			}
		}
	}
	return
}
