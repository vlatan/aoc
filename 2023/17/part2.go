package day17

import (
	"container/heap"
	"fmt"
)

// https://adventofcode.com/2023/day/17
func Part2() {
	m := parseFile("17/input.txt")
	start, end := P{0, 0}, P{len(m) - 1, len(m[0]) - 1}
	result := solve2(m, start, end, CreateNextStates2)
	fmt.Println(result)
}

func solve2(m Matrix, start, end P, fn createStates) int {
	visited := make(map[State]struct{})
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{start, 1, 0, 0, 0})
	heap.Push(pq, State{start, 0, 1, 0, 0})

	for pq.Len() > 0 {
		// pop the state with the least accumulated heat loss
		state := heap.Pop(pq).(State)

		// if end location we're done
		if state.loc == end && !(state.streak >= 1 && state.streak <= 3) {
			return state.loss
		}

		states := fn(m, state)
		// fmt.Println(states)
		for _, n := range states {
			// if state not processed
			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
				// compound the heat loss for this state
				n.loss += state.loss
				// push the state to heap to be processed
				heap.Push(pq, n)
			}
		}
	}
	return 0
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
