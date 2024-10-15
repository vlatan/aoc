package day18

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

type P struct{ x, y int }
type Polygon []P
type Box struct{ xMin, yMin, xMax, yMax int }
type processLine func([]string) (string, int)

func parseFile(path string, fn processLine) (Polygon, Box) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	vertex := P{0, 0}
	polygon, box := Polygon{vertex}, Box{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		direction, steps := fn(fields)
		x, y := vertex.x, vertex.y

		switch direction {
		case "U":
			vertex = P{x, y + steps}
			polygon = append(polygon, vertex)
			box.yMax = max(y+steps, box.yMax)
		case "D":
			vertex = P{x, y - steps}
			polygon = append(polygon, vertex)
			box.yMin = min(y-steps, box.yMin)
		case "L":
			vertex = P{x - steps, y}
			polygon = append(polygon, vertex)
			box.xMin = min(x-steps, box.xMin)
		case "R":
			vertex = P{x + steps, y}
			polygon = append(polygon, vertex)
			box.xMax = max(x+steps, box.xMax)
		}
	}
	return polygon[:len(polygon)-1], box
}

func (point P) castRay(polygon Polygon) (r int) {
	// fmt.Println(point)
	x, y := point.x, point.y
	for i, currentVertex := range polygon {
		xi, yi := currentVertex.x, currentVertex.y
		nextVertex := polygon[(i+1)%len(polygon)]
		xj, yj := nextVertex.x, nextVertex.y

		inXRange := (xi <= x && x <= xj) || (xi >= x && x >= xj)
		inYRange := (yi <= y && y <= yj) || (yi >= y && y >= yj)

		if yi == yj && inXRange {
			if y == yi {
				return 1
			}
		} else if xi == xj && inYRange {
			if x == xi {
				return 1
			}
		}

		if xi < xj {
			inXRange = (xi <= x && x <= xj-1) || (xi >= x && x >= xj-1)
		} else if xi > xj {
			inXRange = (xi-1 <= x && x <= xj) || (xi-1 >= x && x >= xj)
		}

		if yi == yj && y > yi && inXRange {
			r = 1 - r
		}

	}
	// if r == 1 {
	// 	fmt.Println(point, "INSIDE")
	// }
	return
}

// Yield a small bounding box (quadrant) to the consumer
func (b Box) Grid(size int) chan Box {
	numCols := (b.xMax - b.xMin) / size
	numRows := (b.yMax - b.yMin) / size

	// calculate remaining width/height if any
	remainingWidth := (b.xMax - b.xMin) % size
	remainingHeight := (b.yMax - b.yMin) % size

	ch := make(chan Box)
	produce := func() {
		defer close(ch)
		for i := range numRows {
			for j := range numCols {
				x := b.xMin + j*size
				y := b.yMin + i*size

				width := size
				if j == numCols-1 {
					width += remainingWidth
				}

				height := size
				if i == numRows-1 {
					height += remainingHeight
				}

				q := Box{x, y, x + width, y + height}
				if b.xMax != q.xMax {
					q.xMax--
				}
				if b.yMax != q.yMax {
					q.yMax--
				}
				ch <- q // yield quadrant to the consumer
			}
		}
	}
	go produce()
	return ch
}

// Recursivelly count quadrant perimeter points if they are in the polygon
func (b Box) Count(polygon Polygon) (r int) {

	// the bounding box is exausted
	if b.xMax < b.xMin || b.yMax < b.yMin {
		return
	}

	// one point left
	if b.xMax == b.xMin && b.yMax == b.yMin {
		r += (P{b.xMax, b.yMax}).castRay(polygon)
		return
	}

	// one column left
	if b.xMax == b.xMin && b.yMax > b.yMin {
		for y := b.yMin; y <= b.yMax; y++ {
			r += (P{b.xMax, y}).castRay(polygon)
		}
		return
	}

	// one row left
	if b.xMax > b.xMin && b.yMax == b.yMin {
		for x := b.xMin; x <= b.xMax; x++ {
			r += (P{x, b.yMax}).castRay(polygon)
		}
		return
	}

	// process perimeter
	for x := b.xMin + 1; x < b.xMax; x++ {
		r += (P{x, b.yMin}).castRay(polygon)
		r += (P{x, b.yMax}).castRay(polygon)
	}

	for y := b.yMin; y <= b.yMax; y++ {
		r += (P{b.xMin, y}).castRay(polygon)
		r += (P{b.xMax, y}).castRay(polygon)
	}

	if r == 0 {
		return
	}

	// size down the bounding box
	b = Box{b.xMin + 1, b.yMin + 1, b.xMax - 1, b.yMax - 1}
	return r + b.Count(polygon)
}
