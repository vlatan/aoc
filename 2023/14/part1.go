package day14

import (
	"fmt"
)

func Part1() {
	matrix := parseFile("14/input.txt")
	matrix.tilt()
	fmt.Println(matrix.count())

}

func (m Matrix) tilt() (r bool) {
	for {
		for x := 1; x < len(m); x++ {
			for y := 0; y < len(m[0]); y++ {
				if m[x][y] == 'O' && m[x-1][y] == '.' {
					m[x-1][y], m[x][y] = 'O', '.'
					r = true
				}
			}
		}
		if !r {
			return
		}
		r = false
	}
}

func (m Matrix) count() (r int) {
	for x := 0; x < len(m); x++ {
		count := 0
		for y := 0; y < len(m[0]); y++ {
			if m[x][y] == 'O' {
				count++
			}
		}
		r += count * (len(m) - x)
	}
	return
}
