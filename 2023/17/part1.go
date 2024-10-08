package day17

import (
	"container/heap"
	"fmt"
)

// https://adventofcode.com/2023/day/17
func Part1() {
	m := parseFile("17/input.txt")
	result := solve(m, P{0, 0}, P{len(m) - 1, len(m[0]) - 1})
	fmt.Println(result)
}

type States []State
type P struct{ x, y int }
type Direction uint8

const (
	N Direction = iota
	E
	S
	W
)

func solve(m Matrix, start, end P) int {
	visited := make(map[State]struct{})
	pq := &PriorityQueue{{start, W, 0, 0}, {start, N, 0, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		// pop the state with the least accumulated heat loss
		state := heap.Pop(pq).(State)

		// if end state
		if state.loc == end {
			return state.loss
		}

		// get legit States
		nextStates := getNextStates(m, state)
		for _, n := range nextStates {
			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
				n.loss = state.loss + m[n.loc.x][n.loc.y]
				heap.Push(pq, n)
			}
		}
	}
	return 0
}

func getNextStates(m Matrix, c State) (r States) {

	switch c.comingFrom {
	case N, S:
		r.Add(c.loc.x, c.loc.y-1, m, E, 0)
		r.Add(c.loc.x, c.loc.y+1, m, W, 0)
	case E, W:
		r.Add(c.loc.x-1, c.loc.y, m, S, 0)
		r.Add(c.loc.x+1, c.loc.y, m, N, 0)
	}

	switch c.comingFrom {
	case N:
		r.Add(c.loc.x+1, c.loc.y, m, N, c.streak)
	case E:
		r.Add(c.loc.x, c.loc.y-1, m, E, c.streak)
	case S:
		r.Add(c.loc.x-1, c.loc.y, m, S, c.streak)
	case W:
		r.Add(c.loc.x, c.loc.y+1, m, W, c.streak)
	}
	return
}

func (slice *States) Add(x, y int, m Matrix, d Direction, streak int) {
	if inBounds(m, P{x, y}) && streak < 3 {
		crucible := State{P{x, y}, d, streak + 1, m[x][y]}
		*slice = append(*slice, crucible)
	}
}

func inBounds(m Matrix, p P) bool {
	return p.x >= 0 && p.x < len(m) &&
		p.y >= 0 && p.y < len(m[0])
}
