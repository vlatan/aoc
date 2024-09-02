package wait

import (
	"aoc/2023/utils"
	"fmt"
)

// https://adventofcode.com/2023/day/6
func WaitForItPartOne() {
	times, distances := parseFile("06/input.txt")
	result := 1
	for i, t := range times {
		count, time := 0, utils.ToInteger(t)
		distance := utils.ToInteger(distances[i])
		for sec := 0; sec <= time; sec++ {
			if sec*(time-sec) > distance {
				count++
			}
		}
		result *= count
	}
	fmt.Println(result)
}
