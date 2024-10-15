package day18

import (
	"fmt"
	"strconv"
)

// https://adventofcode.com/2023/day/18
func Part1() {
	polygon, perimeter := parseFile("18/input.txt", processLine1)
	fmt.Println(area(polygon, perimeter))
}

func processLine1(fields []string) (string, int) {
	steps, _ := strconv.ParseInt(fields[1], 10, 64)
	return fields[0], int(steps)
}
