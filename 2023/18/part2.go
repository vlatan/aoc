package day18

import (
	"fmt"
	"strconv"
)

// https://adventofcode.com/2023/day/18
func Part2() {
	polygon, perimeter := parseFile("18/input.txt", processLine2)
	fmt.Println(area(polygon, perimeter))
}

func processLine2(fields []string) (string, int) {
	steps, _ := strconv.ParseInt(fields[2][2:7], 16, 64)
	direction := ""
	switch fields[2][7] {
	case '0':
		direction = "R"
	case '1':
		direction = "D"
	case '2':
		direction = "L"
	case '3':
		direction = "U"
	}
	return direction, int(steps)
}
