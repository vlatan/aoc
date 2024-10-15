package day18

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

type P struct{ x, y int }
type Polygon []P
type processLine func([]string) (string, int)

func parseFile(path string, fn processLine) (Polygon, int) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	vertex := P{0, 0}
	polygon, perimeter := Polygon{vertex}, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		direction, steps := fn(fields)
		perimeter += steps
		x, y := vertex.x, vertex.y

		switch direction {
		case "U":
			vertex = P{x, y + steps}
		case "D":
			vertex = P{x, y - steps}
		case "L":
			vertex = P{x - steps, y}
		case "R":
			vertex = P{x + steps, y}
		}
		polygon = append(polygon, vertex)
	}
	return polygon[:len(polygon)-1], perimeter
}

// Shoelace formula is calculating the area of a polygon on a cartesian plane.
// But, only L/2 - 1 of the border is included in this area because the
// coordinates of the vertices are actually in the center of the squares that they
// reprsent. So the border line according to this formula is passing through the middle
// of the border squares. Therefore 1/2 of each square on the border is included
// in this area plus 1/4 or 3/4 of the squares on the corners depending on where they're facing.
// It turns out in a simple polygon the convex corners (those that add 1/4 to the area)
// always win by 4. All the other corners in pairs comprise 1/4 + 3/4 which is 1/2 + 1/2
// meaning they are treated the same as the other non-corner squares.
// Obviously, these four unpaired corners contribute one whole square 4 * 1/4 = 1.
// So the border amount included in this area is actually (L-4)/2 + 1 = L/2 - 1.
// Having said that, if you want to include the WHOLE border in the area of the polygon
// we need to add L/2 + 1 to the shoelace area.
// https://en.wikipedia.org/wiki/Shoelace_formula
// https://en.wikipedia.org/wiki/Pick%27s_theorem
func area(polygon Polygon, L int) (r int) {
	for i, vertex := range polygon {
		nextVertex := polygon[(i+1)%len(polygon)]
		r += vertex.x * nextVertex.y
		r -= vertex.y * nextVertex.x
	}
	shoelaceArea := abs(r) / 2
	return shoelaceArea + (L/2 + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
