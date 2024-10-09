package day17

import (
	"container/heap"
	"fmt"
)

type P struct{ x, y int }
type State struct{ x, y, dx, dy, streak, loss int }

// https://adventofcode.com/2023/day/17
func Part1() {
	m := parseFile("17/input.txt")
	result := solve(m, P{0, 0}, P{len(m) - 1, len(m[0]) - 1})
	fmt.Println(result)
}

func solve(m Matrix, start, end P) int {
	visited := make(map[State]struct{})
	x, y := start.x, start.y
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{x, y, 1, 0, 0, 0})
	heap.Push(pq, State{x, y, 0, 1, 0, 0})

	for pq.Len() > 0 {
		// pop the state with the least accumulated heat loss
		state := heap.Pop(pq).(State)

		// if end state
		if (P{state.x, state.y}) == end {
			return state.loss
		}

		// create the next states
		for _, n := range getNextStates(m, state) {
			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
				n.loss = state.loss + m[n.x][n.y]
				heap.Push(pq, n)
			}
		}
	}
	return 0
}

func getNextStates(m Matrix, s State) (r []State) {
	xMax, yMax := len(m)-1, len(m[0])-1

	dx, dy := s.dy, -s.dx
	x, y := s.x+dx, s.y+dy
	if inBounds(xMax, yMax, x, y) {
		r = append(r, State{x, y, dx, dy, 1, m[x][y]})
	}

	dx, dy = -s.dy, s.dx
	x, y = s.x+dx, s.y+dy
	if inBounds(xMax, yMax, x, y) {
		r = append(r, State{x, y, dx, dy, 1, m[x][y]})
	}

	dx, dy = s.dx, s.dy
	x, y = s.x+dx, s.y+dy
	if inBounds(xMax, yMax, x, y) && s.streak < 3 {
		r = append(r, State{x, y, dx, dy, s.streak + 1, m[x][y]})
	}

	return
}

func inBounds(xMax, yMax, x, y int) bool {
	return x >= 0 && x <= xMax &&
		y >= 0 && y <= yMax
}
