package day08

import "fmt"

// https://adventofcode.com/2023/day/8
func Part2() {
	commands, graph := parseFile("08/input.txt")

	currentNodes := []string{}
	for node := range graph {
		if node[len(node)-1] == 'A' {
			currentNodes = append(currentNodes, node)
		}
	}

	idx := map[rune]int{'L': 0, 'R': 1}
	count, steps := 0, make([]int, len(currentNodes))
	for {
		for _, cmd := range commands {
			for i, node := range currentNodes {
				currentNodes[i] = graph[node][idx[cmd]]
				if node[len(node)-1] == 'Z' && steps[i] == 0 {
					steps[i] = count
				}
			}
			if allDone(steps) {
				fmt.Println("Part 2:", lcmm(steps))
				return
			}
			count++
		}
	}

}

// GCD is the Greatest Common Divisor
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM is the Least Common Multiple which is the smallest positive integer
// that is divisible by all the nums in the  xs []int.
func lcmm(xs []int) int {
	// a * b / gcd(a, b)  =  formula for LCM of two numbers
	lcm := func(a, b int) int { return a * b / gcd(a, b) }

	// calculate the LCM for all nums in xs []int
	// by recalculating the LCM of all the previous numbers and the current one
	result := 1
	for _, n := range xs {
		result = lcm(result, n)
	}
	return result
}

func allDone(lst []int) bool {
	for _, num := range lst {
		if num == 0 {
			return false
		}
	}
	return true
}
