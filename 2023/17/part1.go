package day17

import "fmt"

// https://adventofcode.com/2023/day/17
func Part1() {
	m := parseFile("17/input.txt")
	start, end := P{0, 0}, P{len(m) - 1, len(m[0]) - 1}
	result := solve(m, start, end, endState1, CreateNextStates1)
	fmt.Println("Part 1:", result)
}

func endState1(s State, end P) bool {
	return s.loc == end
}

func CreateNextStates1(m Matrix, s State) (r States) {
	// turn 90° on both sides
	r.Push(s.dy, -s.dx, 1, s.loc, m)
	r.Push(-s.dy, s.dx, 1, s.loc, m)
	// continue in the same direction if allowed
	if s.streak < 3 {
		r.Push(s.dx, s.dy, s.streak+1, s.loc, m)
	}
	return
}
