package day17

import "fmt"

// https://adventofcode.com/2023/day/17
func Part2() {
	m := parseFile("17/input.txt")
	start, end := P{0, 0}, P{len(m) - 1, len(m[0]) - 1}
	result := solve(m, start, end, endState2, CreateNextStates2)
	fmt.Println("Part 2:", result)
}

func endState2(s State, end P) bool {
	return s.loc == end && !(s.streak >= 1 && s.streak <= 3)
}

func CreateNextStates2(m Matrix, s State) (r States) {
	if s.streak == 0 || s.streak == 10 {
		// on start (where streak is 0) and on streak 10
		// turn ONLY 90° on both sides
		r.Push(s.dy, -s.dx, 1, s.loc, m)
		r.Push(-s.dy, s.dx, 1, s.loc, m)
	} else if s.streak <= 3 {
		// continue ONLY in the same direction until streak 4
		r.Push(s.dx, s.dy, s.streak+1, s.loc, m)
	} else {
		// between streak 4 and 9
		// turn 90° in both sides and continue in the same direction
		r.Push(s.dy, -s.dx, 1, s.loc, m)
		r.Push(-s.dy, s.dx, 1, s.loc, m)
		r.Push(s.dx, s.dy, s.streak+1, s.loc, m)
	}
	return
}
