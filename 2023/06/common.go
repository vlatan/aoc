package day06

import (
	"aoc/2023/common"
	"bufio"
	"os"
	"strings"
)

func parseFile(path string) ([]string, []string) {
	file, err := os.Open(path)
	common.Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	times := strings.Fields(scanner.Text())[1:]
	scanner.Scan()
	distances := strings.Fields(scanner.Text())[1:]
	return times, distances
}
