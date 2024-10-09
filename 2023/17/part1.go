package day17

import (
	"container/heap"
	"fmt"
)

type P struct{ x, y int }
type State struct {
	loc                  P
	dx, dy, streak, loss int
}
type States []State

// https://adventofcode.com/2023/day/17
func Part1() {
	m := parseFile("17/input.txt")
	start, end := P{0, 0}, P{len(m) - 1, len(m[0]) - 1}
	result := solve(m, start, end)
	fmt.Println(result)
}

func solve(m Matrix, start, end P) int {
	visited := make(map[State]struct{})
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{start, 1, 0, 0, 0})
	heap.Push(pq, State{start, 0, 1, 0, 0})

	for pq.Len() > 0 {
		// pop the state with the least accumulated heat loss
		state := heap.Pop(pq).(State)

		// if end state
		if state.loc == end {
			return state.loss
		}

		// create the next states
		for _, n := range state.CreateNewStates(m) {
			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
				n.loss = state.loss + m[n.loc.x][n.loc.y]
				heap.Push(pq, n)
			}
		}
	}
	return 0
}

func (s State) CreateNewStates(m Matrix) (r States) {
	r.Push(s.dy, -s.dx, 1, s.loc, m)
	r.Push(-s.dy, s.dx, 1, s.loc, m)
	if s.streak < 3 {
		r.Push(s.dx, s.dy, s.streak+1, s.loc, m)
	}

	return
}

func inBounds(xMax, yMax, x, y int) bool {
	return x >= 0 && x <= xMax &&
		y >= 0 && y <= yMax
}

func (r *States) Push(dx, dy, streak int, loc P, m Matrix) {
	xMax, yMax := len(m)-1, len(m[0])-1
	x, y := loc.x+dx, loc.y+dy
	if inBounds(xMax, yMax, x, y) {
		*r = append(*r, State{P{x, y}, dx, dy, streak, m[x][y]})
	}
}
