package day08

import "fmt"

// https://adventofcode.com/2023/day/8
func Part1() {
	commands, graph := parseFile("08/input.txt")
	idx := map[rune]int{'L': 0, 'R': 1}
	steps, current := 0, "AAA"
	for {
		for _, cmd := range commands {
			if current == "ZZZ" {
				fmt.Println("Part 1:", steps)
				return
			}
			current = graph[current][idx[cmd]]
			steps++
		}
	}
}
