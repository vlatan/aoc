package scratchcards

import (
	"aoc/2023/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

// https://adventofcode.com/2023/day/4
func ScratchcardsPartOne() {
	file, err := os.Open("04/input.txt")
	utils.Check(err)
	defer file.Close()
	result, scanner := 0, bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		winning, scratched := getSets(line)
		count := 0.0
		for _, num := range scratched {
			if slices.Contains(winning, num) {
				count += 1
			}
		}
		result += int(math.Pow(2, count-1))
	}
	fmt.Println(result)
}
