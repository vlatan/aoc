package day18

import (
	"fmt"
	"strconv"
)

// https://adventofcode.com/2023/day/18
func Part1() {
	polygon, perimeter := parseFile("18/input.txt", noColor)
	fmt.Println("Part 1:", area(polygon, perimeter))
}

func noColor(fields []string) (string, int) {
	steps, _ := strconv.ParseInt(fields[1], 10, 64)
	return fields[0], int(steps)
}
