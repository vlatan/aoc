package day19

import (
	"aoc/2023/common"
	"fmt"
	"strings"
)

// https://adventofcode.com/2023/day/19
func Part1() {
	workflows, ratings := parseFile("19/input.txt")
	result := 0
	for _, rating := range ratings {
		result += solve("in", workflows, rating)
	}
	fmt.Println("Part 1:", result)
}

func solve(start string, workflows Workflows, rating Rating) int {
	result, w := 0, workflows[start]

	switch endCase := w.NextWorkflow(rating); endCase {
	case "A":
		for _, r := range rating {
			result += r
		}
		return result
	case "R":
		return 0
	default:
		return solve(endCase, workflows, rating)
	}
}

func (workflow Workflow) NextWorkflow(rating Rating) string {
	for _, w := range workflow {

		parts := strings.Split(w, ":")
		if len(parts) == 1 {
			return parts[0]
		}

		if comparison := strings.Split(parts[0], ">"); len(comparison) == 2 {
			if rating[comparison[0]] > common.ToInteger(comparison[1]) {
				return parts[1]
			}
			continue
		}

		comparison := strings.Split(parts[0], "<")
		if rating[comparison[0]] < common.ToInteger(comparison[1]) {
			return parts[1]
		}
	}
	return "A"
}
