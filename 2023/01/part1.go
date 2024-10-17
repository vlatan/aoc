package day01

import (
	"aoc/2023/utils"
	"bufio"
	"fmt"
	"os"
)

// '0' -> 48, '9' -> 57
// https://adventofcode.com/2023/day/1
func Part1() {
	file, err := os.Open("01/input.txt")
	utils.Check(err)
	defer file.Close()

	sum, scanner := 0, bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += firstDigit(line)*10 + lastDigit(line)
	}
	fmt.Println("Part 1:", sum)
}

func firstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		if utils.IsDigit(s[i]) {
			return int(s[i] - '0')
		}
	}
	return 0
}

func lastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if utils.IsDigit(s[i]) {
			return int(s[i] - '0')
		}
	}
	return 0
}
