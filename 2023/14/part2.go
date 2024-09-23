package day14

import (
	"fmt"
	"strings"
)

// https://adventofcode.com/2023/day/14
func Part2() {
	matrix := parseFile("14/input.txt")
	for range minRange(matrix, 1_000_000_000) {
		matrix.Cycle()
	}
	fmt.Println(matrix.Count())
}

func minRange(m Matrix, bound int) int {
	m = m.Copy()
	var start, end int
	seen := map[string]int{}

	for i := 1; i <= bound; i++ {
		m.Cycle()
		key := m.String()
		if val, ok := seen[key]; ok {
			start, end = val, i
			break
		}
		seen[key] = i
	}
	return start + (bound-start)%(end-start)
}

func (m Matrix) Copy() Matrix {
	cp := make(Matrix, len(m))
	for i := range m {
		cp[i] = make([]byte, len(m[i]))
		copy(cp[i], m[i])
	}
	return cp
}

func (m Matrix) String() string {
	b := strings.Builder{}
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[x]); y++ {
			b.WriteByte(m[x][y])
		}
	}
	return b.String()
}

func (m Matrix) Cycle() {
	m.North()
	m.West()
	m.South()
	m.East()
}

func (m Matrix) West() {
	for move := true; move; {
		move = false
		for x := 0; x < len(m); x++ {
			for y := 1; y < len(m[0]); y++ {
				if m[x][y] == 'O' && m[x][y-1] == '.' {
					m[x][y-1], m[x][y] = 'O', '.'
					move = true
				}
			}
		}
	}
}

func (m Matrix) South() {
	for move := true; move; {
		move = false
		for x := 1; x < len(m); x++ {
			for y := 0; y < len(m[0]); y++ {
				if m[x-1][y] == 'O' && m[x][y] == '.' {
					m[x][y], m[x-1][y] = 'O', '.'
					move = true
				}
			}
		}
	}
}

func (m Matrix) East() {
	for move := true; move; {
		move = false
		for x := 0; x < len(m); x++ {
			for y := 1; y < len(m[0]); y++ {
				if m[x][y-1] == 'O' && m[x][y] == '.' {
					m[x][y], m[x][y-1] = 'O', '.'
					move = true
				}
			}
		}
	}
}
