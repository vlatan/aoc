package day15

import (
	"aoc/2023/common"
	"os"
	"strings"
)

func parseFile(path string) []string {
	content, err := os.ReadFile(path)
	common.Check(err)
	return strings.Split(string(content), ",")
}

func hash(s string) (r int) {
	for _, c := range s {
		r = (r + int(c)) * 17 % 256
	}
	return
}
