package day05

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

func parseFile(path string) [][]int {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()
	scanner, result := bufio.NewScanner(file), [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		current := make([]int, len(fields))
		for i, str := range fields {
			current[i] = utils.ToInteger(str)
		}
		result = append(result, current)
	}
	return result
}
