package day15

import (
	"aoc/2023/utils"
	"os"
	"strings"
)

func parseFile(path string) []string {
	content, err := os.ReadFile(path)
	utils.Check(err)
	return strings.Split(string(content), ",")
}
