package day17

import (
	"aoc/2023/common"
	"bufio"
	"container/heap"
	"os"
)

type Matrix [][]int
type P struct{ x, y int }
type State struct {
	loc                  P
	dx, dy, streak, loss int
}
type States []State
type isEnd func(State, P) bool
type createStates func(Matrix, State) States

func parseFile(path string) (r Matrix) {
	file, err := os.Open(path)
	common.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		slice := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			slice[i] = int(line[i] - '0')
		}
		r = append(r, slice)
	}
	return
}

func solve(m Matrix, start, end P, fn1 isEnd, fn2 createStates) int {
	visited := make(map[State]struct{})
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{start, 1, 0, 0, 0})
	heap.Push(pq, State{start, 0, 1, 0, 0})

	for pq.Len() > 0 {
		// pop the state with the least accumulated heat loss
		state := heap.Pop(pq).(State)

		// if end location we're done
		if fn1(state, end) {
			return state.loss
		}

		for _, n := range fn2(m, state) {
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

func (r *States) Push(dx, dy, streak int, loc P, m Matrix) {
	x, y := loc.x+dx, loc.y+dy
	// if in bounds append new state to slice
	if x >= 0 && x < len(m) && y >= 0 && y < len(m[0]) {
		*r = append(*r, State{P{x, y}, dx, dy, streak, m[x][y]})
	}
}
