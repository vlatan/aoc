package day12

import (
	"fmt"
	"slices"
)

type Cache map[string]int

// https://adventofcode.com/2023/day/12
func Part1() {
	lines, groups := parseFile("12/input.txt")
	result := 0
	for i := 0; i < len(lines); i++ {
		result += solve(lines[i], groups[i], Cache{})
	}
	fmt.Println(result)
}

// Recursive function to branch out on "?"
// and test two paths ("." and "#")
// only if those paths are worth considering.
func solve(s string, gs []int, cache Cache) int {
	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			continue
		}
		result := 0
		ps := s[:i] + "."
		if gg := getGroups(ps); viable(gg, gs) {
			key := s[i+1:] + fmt.Sprintf("%v", gg)
			if val, ok := cache[key]; ok {
				result += val
			} else {
				result += solve(ps+s[i+1:], gs, cache)
				cache[key] = result
			}
		}
		ps = s[:i] + "#"
		if gg := getGroups(ps); viable(gg, gs) {
			key := s[i+1:] + fmt.Sprintf("%v", gg)
			if val, ok := cache[key]; ok {
				result += val
			} else {
				result += solve(ps+s[i+1:], gs, cache)
			}
		}
		return result
	}
	if slices.Equal(getGroups(s), gs) {
		return 1
	}
	return 0
}

// Calculate groups of '#' from a string
func getGroups(s string) []int {
	groupFound, groups := false, []int{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '#':
			if groupFound {
				groups[len(groups)-1]++
				continue
			}
			groupFound = true
			groups = append(groups, 1)
		case '.':
			groupFound = false
		}
	}
	return groups
}

// Check if a partial string is viable to continue.
// Has the correct groups so far.
func viable(g1 []int, g2 []int) bool {
	if len(g1) > len(g2) {
		return false
	}
	for i := 0; i < len(g1); i++ {
		if i == len(g1)-1 && g1[i] <= g2[i] {
			return true
		}
		if g1[i] != g2[i] {
			return false
		}
	}
	return true
}
