package day06

import (
	"aoc/2023/common"
	"fmt"
	"strings"
)

// https://adventofcode.com/2023/day/6
func Part2() {
	times, distances := parseFile("06/input.txt")
	time := common.ToInteger(strings.Join(times, ""))
	distance := common.ToInteger(strings.Join(distances, ""))

	count := 0
	for sec := 0; sec <= time; sec++ {
		if sec*(time-sec) > distance {
			count++
		}
	}
	fmt.Println("Part 2:", count)
}
