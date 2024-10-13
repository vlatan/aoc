package day18

import (
	"fmt"
	"strconv"
)

// https://adventofcode.com/2023/day/18
func Part1() {
	graph, bBox := parseFile("18/input.txt", processLine1)
	result := bBox.Count(&bBox, graph)
	fmt.Println(result)
}

func processLine1(fields []string) (string, uint64) {
	steps, _ := strconv.ParseUint(fields[1], 10, 32)
	return fields[0], steps
}
