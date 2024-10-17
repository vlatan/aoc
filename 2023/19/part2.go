package day19

import (
	"aoc/2023/common"
	"fmt"
	"maps"
	"strings"
)

type Range struct{ from, to int }
type RatingState map[string]Range
type WorkflowState struct {
	workflow string
	state    RatingState
}

// https://adventofcode.com/2023/day/19
func Part2() {
	workflows, _ := parseFile("19/input.txt")
	sr := Range{1, 4000}
	state := RatingState{"x": sr, "m": sr, "a": sr, "s": sr}
	result := solve2("in", workflows, state)
	fmt.Println(result)
}

func solve2(start string, workflows Workflows, state RatingState) int {
	result, w := 0, workflows[start]
	workflowStates := w.NextWorkflows(state)
	for _, ws := range workflowStates {
		switch ws.workflow {
		case "A":
			combo := 1
			for _, r := range ws.state {
				combo *= r.to - r.from + 1
			}
			result += combo
		case "R":
			continue
		default:
			result += solve2(ws.workflow, workflows, ws.state)
		}
	}
	return result
}

func (workflow Workflow) NextWorkflows(state RatingState) (r []WorkflowState) {
	for _, w := range workflow {

		parts := strings.Split(w, ":")
		if len(parts) == 1 {
			r = append(r, WorkflowState{parts[0], state})
			return
		}

		if comparison := strings.Split(parts[0], ">"); len(comparison) == 2 {
			stateRange := state[comparison[0]]
			value := common.ToInteger(comparison[1])
			if stateRange.from <= value && value <= stateRange.to {
				currentState := maps.Clone(state)
				currentState[comparison[0]] = Range{value + 1, stateRange.to}
				r = append(r, WorkflowState{parts[1], currentState})
				state[comparison[0]] = Range{stateRange.from, value}
			}
			continue
		}

		if comparison := strings.Split(parts[0], "<"); len(comparison) == 2 {
			stateRange := state[comparison[0]]
			value := common.ToInteger(comparison[1])
			if stateRange.from <= value && value <= stateRange.to {
				currentState := maps.Clone(state)
				currentState[comparison[0]] = Range{stateRange.from, value - 1}
				r = append(r, WorkflowState{parts[1], currentState})
				state[comparison[0]] = Range{value, stateRange.to}
			}
		}
	}
	return
}
