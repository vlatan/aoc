package day12

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

func parseFile(path string) map[string][]int {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()
	scanner, result := bufio.NewScanner(file), map[string][]int{}
	for x := 0; scanner.Scan(); x++ {
		fields := strings.Fields(scanner.Text())
		groups := []int{}
		for i := 0; i < len(fields[1]); i++ {
			if utils.IsDigit(fields[1][i]) {
				groups = append(groups, int(fields[1][i]-'0'))
			}
		}
		result[fields[0]] = groups
	}
	return result
}
