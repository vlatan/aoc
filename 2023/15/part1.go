package day15

import "fmt"

func Part1() {
	data := parseFile("15/input.txt")
	result := 0
	for _, code := range data {
		current := 0
		for i := 0; i < len(code); i++ {
			current = (current + int(code[i])) * 17 % 256
		}
		result += current
	}
	fmt.Println(result)
}
