package day18

import (
	"fmt"
	"strconv"
)

// https://adventofcode.com/2023/day/18
func Part2() {
	graph, bBox := parseFile("18/input.txt", processLine1)
	grid := bBox.MakeGrid(7)

	result := 0
	for b := range grid {
		vertices := []P{
			{b.xMin, b.yMin}, {b.xMin, b.yMax},
			{b.xMax, b.yMin}, {b.xMax, b.yMax},
		}

		// Traverse the perimeter of b.
		// If all points are not in the polygon skip the entire b.
		// If you encounter a point that is in the polygon,
		// continue processing the perimeter and count points that are inside the polygon.
		// When done shrink the perimeter and repeat the procedure.
		// Probably a nice oportunity for recursion

		clean := true
		for _, vertex := range vertices {
			if vertex.castRay(graph, bBox) {
				clean = false
				break
			}
		}

		if clean {
			continue
		}

		// cast rays for all the points in this quadrant
		for x := b.xMin; x <= b.xMax; x++ {
			for y := b.yMin; y <= b.yMax; y++ {
				if (P{x, y}).castRay(graph, bBox) {
					result++
				}
			}
		}
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
func (bBox BoundingBox) MakeGrid(cellSize int) chan BoundingBox {
	numCols := (bBox.xMax - bBox.xMin) / cellSize
	numRows := (bBox.yMax - bBox.yMin) / cellSize

	// calculate remaining width/height if any
	remainingWidth := (bBox.xMax - bBox.xMin) % cellSize
	remainingHeight := (bBox.yMax - bBox.yMin) % cellSize

	ch := make(chan BoundingBox)
	produce := func() {
		defer close(ch)
		for i := range numRows {
			for j := range numCols {
				x := bBox.xMin + j*cellSize
				y := bBox.yMin + i*cellSize

				width := cellSize
				if j == numCols-1 {
					width += remainingWidth
				}

				height := cellSize
				if i == numRows-1 {
					height += remainingHeight
				}

				b := BoundingBox{x, y, x + width, y + height}
				if bBox.xMax != b.xMax {
					b.xMax--
				}
				if bBox.yMax != b.yMax {
					b.yMax--
				}
				ch <- b // yield quadrant to the consumer
			}
		}
	}
	go produce()
	return ch
}
