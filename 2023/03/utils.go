package gear

type pos struct {
	x, y int
}

type number struct {
	value string
	start pos
	end   pos
}

// Check if the number is valid and gahter stars
func inspectAroundNumber(num number, matrix [][]byte) (map[pos]string, bool) {

	// map to gather stars
	stars, valid := make(map[pos]string), false

	// left and right index around the number
	x, leftY, rightY := num.start.x, num.start.y-1, num.end.y+1
	for _, y := range []int{leftY, rightY} {
		if y < 0 || y >= len(matrix[x]) {
			continue
		}
		if matrix[x][y] != '.' {
			valid = true
		}
		if matrix[x][y] == '*' {
			stars[pos{x, y}] = num.value
		}
	}

	// top and bottom indices around the number
	prevX, nextX := x-1, x+1
	for y := leftY; y <= rightY; y++ {
		if y < 0 || y >= len(matrix[x]) {
			continue
		}
		for _, cx := range []int{prevX, nextX} {
			if cx < 0 || cx >= len(matrix) {
				continue
			}
			if matrix[cx][y] != '.' {
				valid = true
			}
			if matrix[cx][y] == '*' {
				stars[pos{cx, y}] = num.value
			}
		}
	}
	return stars, valid
}
