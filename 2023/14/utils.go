package day14

import (
	"aoc/2023/utils"
	"bufio"
	"os"
)

type Matrix [][]byte

func parseFile(path string) (result Matrix) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, []byte(scanner.Text()))
	}
	return
}
