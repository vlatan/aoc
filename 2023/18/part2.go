package day18

import (
	"fmt"
	"strconv"
)

// https://adventofcode.com/2023/day/18
func Part2() {
	polygon, perimeter := parseFile("18/input.txt", color)
	fmt.Println("Part 2:", area(polygon, perimeter))
}

func color(fields []string) (string, int) {
	steps, _ := strconv.ParseInt(fields[2][2:7], 16, 64)
	direction := map[byte]string{'0': "R", '1': "D", '2': "L", '3': "U"}
	return direction[fields[2][7]], int(steps)
}
