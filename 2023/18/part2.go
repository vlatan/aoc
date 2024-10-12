package day18

import (
	"fmt"
	"strconv"
)

// https://adventofcode.com/2023/day/18
func Part2() {
	graph, b := parseFile("18/input.txt", processLine1)
	grid := b.Grid(7)

	result := 0
	for b := range grid {
		result += b.Count(b, graph)

		// Traverse the perimeter of b.
		// If all points are not in the polygon skip the entire b.
		// If you encounter a point that is in the polygon,
		// continue processing the perimeter and count points that are inside the polygon.
		// When done shrink the perimeter and repeat the procedure.
		// Probably a nice oportunity for recursion

	}

	fmt.Println(result)
}

func processLine2(fields []string) (string, uint64) {
	steps, _ := strconv.ParseUint(fields[2][2:7], 16, 32)
	direction := ""
	switch fields[2][7] {
	case '0':
		direction = "R"
	case '1':
		direction = "D"
	case '2':
		direction = "L"
	case '3':
		direction = "U"
	}
	return direction, steps
}

// Yield a small bounding box (quadrant) to the consumer
func (b BoundingBox) Grid(size int) chan BoundingBox {
	numCols := (b.xMax - b.xMin) / size
	numRows := (b.yMax - b.yMin) / size

	// calculate remaining width/height if any
	remainingWidth := (b.xMax - b.xMin) % size
	remainingHeight := (b.yMax - b.yMin) % size

	ch := make(chan BoundingBox)
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

				q := BoundingBox{x, y, x + width, y + height}
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

func (b BoundingBox) Count(p BoundingBox, graph Graph) (r int) {

	if b.xMax <= b.xMin && b.yMax <= b.yMin {
		return
	}

	for x := b.xMin + 1; x < b.xMax; x++ {
		r += (P{x, b.yMin}).castRay(graph, p)
		r += (P{x, b.yMax}).castRay(graph, p)
	}
	for y := b.yMin; y <= b.yMax; y++ {
		r += (P{b.xMin, y}).castRay(graph, p)
		r += (P{b.xMax, y}).castRay(graph, p)
	}

	if r == 0 {
		return
	}

	b = BoundingBox{b.xMin + 1, b.yMin + 1, b.xMax - 1, b.yMax - 1}
	return r + b.Count(p, graph)
}
