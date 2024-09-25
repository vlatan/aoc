package day15

import "fmt"

type Label struct {
	name  string
	focus int
}

type Box []Label

// https://adventofcode.com/2023/day/15
func Part2() {
	data := parseFile("15/input.txt")
	var hashmap [256]Box

	for _, item := range data {
		switch item[len(item)-1] {
		case '-':
			name := item[:len(item)-1]
			hashmap[hash(name)].Remove(name)
		default:
			name := item[:len(item)-2]
			focus := int(item[len(item)-1] - '0')
			label := Label{name, focus}
			hashmap[hash(name)].Replace(label)
		}
	}

	fmt.Println(focusPower(hashmap))
}

func (b *Box) Remove(name string) {
	for i, label := range *b {
		if label.name == name {
			*b = append((*b)[:i], (*b)[i+1:]...)
			return
		}
	}
}

func (b *Box) Replace(label Label) {
	for i, l := range *b {
		if l.name == label.name {
			(*b)[i] = label
			return
		}
	}
	*b = append(*b, label)
}

func focusPower(hashmap [256]Box) (r int) {
	for i, box := range hashmap {
		for j, label := range box {
			r += (i + 1) * (j + 1) * label.focus
		}
	}
	return
}
