package day06

import (
	"aoc/2023/common"
	"fmt"
)

// https://adventofcode.com/2023/day/6
func Part1() {
	times, distances := parseFile("06/input.txt")
	result := 1
	for i, t := range times {
		count, time := 0, common.ToInteger(t)
		distance := common.ToInteger(distances[i])
		for sec := 0; sec <= time; sec++ {
			if sec*(time-sec) > distance {
				count++
			}
		}
		result *= count
	}
	fmt.Println("Part 1:", result)
}
