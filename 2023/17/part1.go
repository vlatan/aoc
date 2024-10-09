package day17

import (
	"fmt"
)

// https://adventofcode.com/2023/day/17
func Part1() {
	m := parseFile("17/input.txt")
	start, end := P{0, 0}, P{len(m) - 1, len(m[0]) - 1}
	result := solve(m, start, end, CreateNextStates)
	fmt.Println(result)
}

func CreateNextStates(m Matrix, s State) (r States) {
	// turn 90° on both sides
	r.Push(s.dy, -s.dx, 1, s.loc, m)
	r.Push(-s.dy, s.dx, 1, s.loc, m)
	// continue in the same direction if allowed
	if s.streak < 3 {
		r.Push(s.dx, s.dy, s.streak+1, s.loc, m)
	}
	return
}
