package day18

import (
	"fmt"
	"strconv"
)

// https://adventofcode.com/2023/day/18
func Part1() {
	graph, bBox := parseFile("18/input.txt", processLine1)
	result := 0
	for x := bBox.xMin; x <= bBox.xMax; x++ {
		for y := bBox.yMin; y <= bBox.yMax; y++ {
			if (P{x, y}).castRay(graph, bBox) {
				result++
			}
		}
	}
	fmt.Println(result)
}

func processLine1(fields []string) (string, uint64) {
	steps, _ := strconv.ParseUint(fields[1], 10, 32)
	return fields[0], steps
}
