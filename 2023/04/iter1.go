package day04

import (
	"aoc/2023/common"
	"bufio"
	"fmt"
	"math"
	"os"
)

// One iteration per line solution of D4P1
// https://adventofcode.com/2023/day/4
func IterPart1() {
	file, err := os.Open("04/input.txt")
	common.Check(err)
	defer file.Close()

	result, scanner := 0, bufio.NewScanner(file)
	for scanner.Scan() {
		line, winning := scanner.Text(), make(map[string]bool)
		for main_cursor := 0; main_cursor < len(line); main_cursor++ {
			if line[main_cursor] == ':' {
				for w := main_cursor + 1; w < len(line); w++ {
					if common.IsDigit(line[w]) {
						new_cursor, num := constructNum(w, line)
						w, winning[num] = new_cursor, true
					} else if line[w] == '|' {
						count := 0.0
						for s := w + 1; s < len(line); s++ {
							if common.IsDigit(line[s]) {
								new_cursor, num := constructNum(s, line)
								s = new_cursor
								if _, ok := winning[num]; ok {
									count += 1
								}
							}
							w = s
						}
						result += int(math.Pow(2, count-1))
					}
					main_cursor = w
				}
			}
		}
	}
	fmt.Println("Iter Part 1:", result)
}

// Advance through string and return num and new index when no more digits
func constructNum(cursor int, line string) (new_cursor int, num string) {
	new_cursor, num = cursor, string(line[cursor])
	for i := cursor + 1; i < len(line); i++ {
		if !common.IsDigit(line[i]) {
			break
		}
		num += string(line[i])
		new_cursor = i
	}
	return
}
