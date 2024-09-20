package day13

import (
	"fmt"
)

// https://adventofcode.com/2023/day/13
func Part2() {
	result := 0
	data, transposed := parseFile("13/input.txt")
	for i := 0; i < len(data); i++ {
		ignore, curr := "p", data[i].Process(100, -1)/100
		if curr == 0 {
			curr = transposed[i].Process(1, -1)
			ignore = "t"
		}
		r := fixSmudge(data[i], transposed[i], curr, ignore)
		result += r
	}
	fmt.Println(result)
}

func fixSmudge(p, t Pattern, curr int, ignore string) int {
	for x := 0; x < len(p); x++ {
		for y := 0; y < len(p[0]); y++ {
			switch p[x][y] {
			case '.':
				p[x] = replace(p[x], '#', y)
				t[y] = replace(t[y], '#', x)

				if r := checkPattern(p, t, curr, ignore); r != 0 {
					return r
				}

				p[x] = replace(p[x], '.', y)
				t[y] = replace(t[y], '.', x)
			case '#':
				p[x] = replace(p[x], '.', y)
				t[y] = replace(t[y], '.', x)

				if r := checkPattern(p, t, curr, ignore); r != 0 {
					return r
				}

				p[x] = replace(p[x], '#', y)
				t[y] = replace(t[y], '#', x)
			}

		}
	}
	return 0
}

func replace(s string, b byte, index int) string {
	out := []byte(s)
	out[index] = b
	return string(out)
}

func checkPattern(p, t Pattern, curr int, ignore string) int {
	switch ignore {
	case "p":
		if r := p.Process(100, curr); r != 0 {
			return r
		}
		return t.Process(1, -1)
	case "t":
		if r := p.Process(100, -1); r != 0 {
			return r
		}
		return t.Process(1, curr)
	case "":
		if r := p.Process(100, -1); r != 0 {
			return r
		}
		return t.Process(1, -1)

	}
	return 0
}
