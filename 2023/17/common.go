package day17

import (
	"aoc/2023/utils"
	"bufio"
	"os"
)

type Matrix [][]int

func parseFile(path string) (r Matrix) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		slice := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			slice[i] = int(line[i] - '0')
		}
		r = append(r, slice)
	}
	return
}
