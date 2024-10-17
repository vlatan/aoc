package day05

import "fmt"

// https://adventofcode.com/2023/day/5
func Part1() {
	seeds, maps := parseFile("05/input.txt")
	minLoc := 1<<63 - 1
	for _, s := range seeds {
		for _, m := range maps {
			s = m.Convert(s)
		}
		minLoc = min(minLoc, s)
	}
	fmt.Println("Part 1:", minLoc)
}

func (m Map) Convert(num int) int {
	for _, line := range m {
		if line.src <= num && num < line.src+line.len {
			return line.dest + (num - line.src)
		}
	}
	return num
}
