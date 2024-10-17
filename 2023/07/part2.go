package day07

import (
	"aoc/2023/utils"
	"fmt"
	"maps"
)

// https://adventofcode.com/2023/day/7
func CamelCardsPartTwo() {
	content := parseFile("07/input.txt")
	hands := make([]Hand, len(content))
	for i := 0; i < len(content); i++ {
		hands[i] = Hand{
			content[i][0],
			utils.ToInteger(content[i][1]),
			jHandStrength(content[i][0]),
		}
	}
	cv := maps.Clone(cardValues)
	cv['J'] = 1
	winnings := winings(hands, cv)
	fmt.Println(winnings)
}

// The hand strength is defined as the sum of the
// each card occurrences raised to the power of two,
// except for the Joker occurences which are
// counted towards the card with most occurrences
func jHandStrength(cards string) int {
	count := make(map[rune]int)
	jokerCount, maxCount, bestCard := 0, 0, 'A'
	for _, card := range cards {
		if card == 'J' {
			jokerCount++
			continue
		}
		count[card]++
		if count[card] > maxCount {
			maxCount = count[card]
			bestCard = card
		}
	}
	// add joker count to the best card count
	count[bestCard] += jokerCount
	return strength(count)
}
