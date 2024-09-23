package day15

import "fmt"

func Part1() {
	data := parseFile("15/input.txt")
	sum := 0
	for _, s := range data {
		sum += hash(s)
	}
	fmt.Println(sum)
}
