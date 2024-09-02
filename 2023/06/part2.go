package wait

import (
	"aoc/2023/utils"
	"fmt"
	"strings"
)

// https://adventofcode.com/2023/day/6
func WaitForItPartTwo() {
	times, distances := parseFile("06/input.txt")
	time := utils.ToInteger(strings.Join(times, ""))
	distance := utils.ToInteger(strings.Join(distances, ""))

	count := 0
	for sec := 0; sec <= time; sec++ {
		if sec*(time-sec) > distance {
			count++
		}
	}
	fmt.Println(count)
}
