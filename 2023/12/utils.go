package day12

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

func parseFile(path string) ([]string, [][]int) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines, groups := []string{}, [][]int{}
	for x := 0; scanner.Scan(); x++ {
		fields := strings.Fields(scanner.Text())
		lines = append(lines, fields[0])
		list := strings.Split(fields[1], ",")
		group := make([]int, len(list))
		for i := 0; i < len(list); i++ {
			group[i] = utils.ToInteger(list[i])
		}
		groups = append(groups, group)
	}
	return lines, groups
}
