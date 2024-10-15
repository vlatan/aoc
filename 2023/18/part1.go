package day18

import (
	"fmt"
	"strconv"
)

// https://adventofcode.com/2023/day/18
func Part1() {
	polygon, box := parseFile("18/input.txt", processLine1)
	// fmt.Println(polygon)
	result := box.Count(polygon)
	fmt.Println(result)
}

func processLine1(fields []string) (string, int) {
	steps, _ := strconv.ParseUint(fields[1], 10, 32)
	return fields[0], int(steps)
}
