package day01

import (
	"aoc/2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// https://adventofcode.com/2023/day/1
func Part2() {
	file, err := os.Open("01/input.txt")
	utils.Check(err)
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += firstNumber(line)*10 + lastNumber(line)
	}

	fmt.Println("Part 2:", sum)
}

func firstNumber(s string) int {
	for i := 0; i < len(s); i++ {
		if utils.IsDigit(s[i]) {
			return int(s[i] - '0')
		}

		for j, d := range digitWords {
			if strings.HasSuffix(s[:i+1], d) {
				return j + 1
			}
		}
	}
	return 0
}

func lastNumber(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if utils.IsDigit(s[i]) {
			return int(s[i] - '0')
		}

		for j, d := range words {
			if strings.HasPrefix(s[i:], d) {
				return j + 1
			}
		}
	}
	return 0
}
