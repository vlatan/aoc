package day09

import (
	"aoc/2023/common"
	"fmt"
)

func Part1() {
	data := parseFile("09/input.txt")
	result := 0
	for _, history := range data {
		lastNums := []int{history[len(history)-1]}
		for {
			current := make([]int, len(history)-1)
			for i := 0; i < len(history)-1; i++ {
				current[i] = history[i+1] - history[i]
			}
			lastNums = append(lastNums, current[len(current)-1])
			if allZeroes(current) {
				result += common.Sum(lastNums)
				break
			}
			history = current
		}
	}
	fmt.Println(result)
}
