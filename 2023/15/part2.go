package day15

import "fmt"

type Label struct {
	name  string
	focus int
}

// https://adventofcode.com/2023/day/15
func Part2() {
	data := parseFile("15/input.txt")
	var hashmap [256][]Label

	for _, item := range data {

		switch item[len(item)-1] {
		case '-':
			name := item[:len(item)-1]
			box := hash(name)
			for i, l := range hashmap[box] {
				if l.name == name {
					hashmap[box] = append(hashmap[box][:i], hashmap[box][i+1:]...)
					break
				}
			}
		default:
			box := hash(item[:len(item)-2])
			label := Label{item[:len(item)-2], int(item[len(item)-1] - '0')}
			replaced := false
			for i, l := range hashmap[box] {
				if l.name == label.name {
					hashmap[box][i] = label
					replaced = true
					break
				}
			}
			if !replaced {
				hashmap[box] = append(hashmap[box], label)
			}

		}
	}

	result := 0
	for i, box := range hashmap {
		for j, label := range box {
			result += (i + 1) * (j + 1) * label.focus
		}
	}

	fmt.Println(result)
}
