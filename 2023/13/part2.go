package day13

import (
	"fmt"
)

// https://adventofcode.com/2023/day/13
func Part2() {
	result := 0
	data, transposed := parseFile("13/input.txt")
	for i := 0; i < len(data); i++ {
		curr := calcRows(data[i], transposed[i], -1)
		result += fixSmudge(data[i], transposed[i], curr)
	}
	fmt.Println("Part 2:", result)
}

func fixSmudge(p, t Pattern, curr int) int {
	for x := 0; x < len(p); x++ {
		for y := 0; y < len(p[0]); y++ {
			switch p[x][y] {
			case '.':
				p[x][y], t[y][x] = '#', '#'
				if r := calcRows(p, t, curr); r != 0 {
					return r
				}
				p[x][y], t[y][x] = '.', '.'
			case '#':
				p[x][y], t[y][x] = '.', '.'
				if r := calcRows(p, t, curr); r != 0 {
					return r
				}
				p[x][y], t[y][x] = '#', '#'
			}
		}
	}
	return 0
}
