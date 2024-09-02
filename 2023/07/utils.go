package camel

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"sort"
	"strings"
)

type Hand struct {
	cards    string
	bid      int
	strength int
}

var cardValues = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11,
	'T': 10, '9': 9, '8': 8, '7': 7, '6': 6,
	'5': 5, '4': 4, '3': 3, '2': 2,
}

func parseFile(path string) [][]string {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := [][]string{}
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		result = append(result, fields)
	}
	return result
}

func winings(hands []Hand, cardV map[rune]int) int {
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].strength == hands[j].strength {
			for k, card1 := range hands[i].cards {
				card2 := rune(hands[j].cards[k])
				if cardV[card1] != cardV[card2] {
					return cardV[card1] < cardV[card2]
				}
			}
		}
		return hands[i].strength < hands[j].strength
	})

	result := 0
	for i, hand := range hands {
		result += hand.bid * (i + 1)
	}
	return result
}

func strength(count map[rune]int) int {
	strength := 0
	for _, value := range count {
		strength += value * value
	}
	return strength
}
