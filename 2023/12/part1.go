package day12

import (
	"fmt"
	"slices"
	"strings"
)

// Every '?' can be '.' or '#'.
// Determine the indexes of the '?' in the string
// Produce the permutations of them being populated with '.' or/and '#'
// The number of the permutations should be 2^n, whene n is the number of '?'
// Try every permutation and check if the string satisfies the contiguous group criteria.
// If so count the permutation as a valid arrangement.
// Take the sum of all posible arangements.

// https://adventofcode.com/2023/day/12
func Part1() {
	lines, groups := parseFile("12/input.txt")

	result := 0
	for i := 0; i < len(lines); i++ {

		// prepare group
		gs := []string{}
		for _, num := range groups[i] {
			group := ""
			for i := 0; i < num; i++ {
				group += "#"
			}
			gs = append(gs, group)
		}

		result += solve(lines[i], gs)
	}
	fmt.Println(result)
}

func solve(s string, gs []string) int {
	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			continue
		}
		s1 := s[:i] + "." + s[i+1:]
		s2 := s[:i] + "#" + s[i+1:]
		return solve(s1, gs) + solve(s2, gs)
	}
	if valid(s, gs) {
		return 1
	}
	return 0
}

func valid(s string, gs []string) bool {
	bss := []string{}
	for _, item := range strings.Split(s, ".") {
		if strings.ContainsRune(item, '#') {
			bss = append(bss, item)
		}
	}
	return slices.Equal(bss, gs)
}
