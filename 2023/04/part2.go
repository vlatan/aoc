package day04

import (
	"aoc/2023/common"
	"bufio"
	"fmt"
	"os"
	"slices"
)

// https://adventofcode.com/2023/day/4
func Part2() {
	file, err := os.Open("04/input.txt")
	common.Check(err)
	defer file.Close()
	cache := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		winning, scratched := getSets(line)
		count := 0
		for _, num := range scratched {
			if slices.Contains(winning, num) {
				count += 1
			}
		}
		cache[i] = count
	}
	sum := 0
	for gameID := range cache {
		sum += processCard(gameID, cache)
	}
	fmt.Println("Part 2:", sum)
}

// Recursive function to count total number of cards
// When comes to the card with ZERO points returns 1
// Otherwise sums up results from the cards down the recursion
// and returns the sum plus one for the card itself
func processCard(gameID int, cache map[int]int) int {
	points, sum := cache[gameID], 0
	if points == 0 {
		return 1
	}
	for i := range points {
		sum += processCard(gameID+i+1, cache)
	}
	return sum + 1
}
