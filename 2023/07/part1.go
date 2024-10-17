package day07

import (
	"aoc/2023/common"
	"fmt"
)

// https://adventofcode.com/2023/day/7
func Part1() {
	content := parseFile("07/input.txt")
	hands := make([]Hand, len(content))
	for i := 0; i < len(content); i++ {
		hands[i] = Hand{
			content[i][0],
			common.ToInteger(content[i][1]),
			handStrength(content[i][0]),
		}
	}
	winnings := winings(hands, cardValues)
	fmt.Println("Part 1:", winnings)
}

// The hand strength is defined as the sum of the
// each card occurrences raised to the power of two
// Example: Three of a kind = 3 * 3 + 1 + 1 = 11
func handStrength(cards string) int {
	count := make(map[rune]int)
	for _, card := range cards {
		count[card]++
	}
	return strength(count)
}
