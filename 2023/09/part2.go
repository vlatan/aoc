package day05

import (
	"fmt"
)

func Part2() {
	data := parseFile("09/input.txt")
	result := 0
	for _, history := range data {
		lastNums := []int{history[0]}
		for {
			current := make([]int, len(history)-1)
			for i := 0; i < len(history)-1; i++ {
				current[i] = history[i+1] - history[i]
			}
			lastNums = append(lastNums, current[0])
			if allZeroes(current) {
				result += substractBackwards(lastNums)
				break
			}
			history = current
		}
	}
	fmt.Println(result)
}

func substractBackwards(lst []int) int {
	result := 0
	for i := len(lst) - 1; i >= 0; i-- {
		result = lst[i] - result
	}
	return result
}
