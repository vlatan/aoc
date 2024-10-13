package day18

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

type P struct{ x, y int }
type Polygon map[P]struct{}
type BoundingBox struct{ xMin, yMin, xMax, yMax int }
type processLine func([]string) (string, uint64)

func parseFile(path string, fn processLine) (Polygon, BoundingBox) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	graph, b := make(Polygon), BoundingBox{}
	current := P{0, 0}
	graph[current] = struct{}{}

	// TODO: Work with just the edges of the polygon, not every point
	// It will change the way how func castRay works

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		direction, steps := fn(fields)
		x, y := current.x, current.y

		switch direction {
		case "U":
			for range steps {
				x++
				current = P{x, y}
				graph[current] = struct{}{}
			}
			b.xMax = max(x, b.xMax)
		case "D":
			for range steps {
				x--
				current = P{x, y}
				graph[current] = struct{}{}
			}
			b.xMin = min(x, b.xMin)
		case "L":
			for range steps {
				y--
				current = P{x, y}
				graph[current] = struct{}{}
			}
			b.yMin = min(y, b.yMin)
		case "R":
			for range steps {
				y++
				current = P{x, y}
				graph[current] = struct{}{}
			}
			b.yMax = max(y, b.yMax)
		}
	}
	return graph, b
}

func (p P) castRay(graph Polygon, b BoundingBox) int {
	// the point is on the polygon
	if _, ok := graph[p]; ok {
		return 1
	}

	count := 0
	// count diagonal ray interections with the polygon
	for i := 1; p.x+i <= b.xMax && p.y+i <= b.yMax; i++ {
		xi, yi := p.x+i, p.y+i

		// check if it's NOT an intersection with the polygon
		if _, ok := graph[P{xi, yi}]; !ok {
			continue
		}

		// Check left/down or up/right neighbours to see
		// if the ray is just grazing a corner on the outside.
		// If so, do not count as an intersection.
		_, left := graph[P{xi, yi - 1}]
		_, down := graph[P{xi + 1, yi}]
		_, right := graph[P{xi, yi + 1}]
		_, up := graph[P{xi - 1, yi}]
		if (left && down) || (up && right) {
			continue
		}

		count++
	}
	// Odd number of interesections means
	// the point is inside inside the polygon
	if count%2 != 0 {
		return 1
	}

	return 0
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

// Recursivelly count quadrant perimeter points if they are in the polygon
func (b BoundingBox) Count(pb *BoundingBox, graph Polygon) (r int) {

	// the bounding box is exausted
	if b.xMax < b.xMin || b.yMax < b.yMin {
		return
	}

	// one point left
	if b.xMax == b.xMin && b.yMax == b.yMin {
		r += (P{b.xMax, b.yMax}).castRay(graph, *pb)
		return
	}

	// one column left
	if b.xMax == b.xMin && b.yMax > b.yMin {
		for y := b.yMin; y <= b.yMax; y++ {
			r += (P{b.xMax, y}).castRay(graph, *pb)
		}
		return
	}

	// one row left
	if b.xMax > b.xMin && b.yMax == b.yMin {
		for x := b.xMin; x <= b.xMax; x++ {
			r += (P{x, b.yMax}).castRay(graph, *pb)
		}
		return
	}

	// process perimeter
	for x := b.xMin + 1; x < b.xMax; x++ {
		r += (P{x, b.yMin}).castRay(graph, *pb)
		r += (P{x, b.yMax}).castRay(graph, *pb)
	}

	for y := b.yMin; y <= b.yMax; y++ {
		r += (P{b.xMin, y}).castRay(graph, *pb)
		r += (P{b.xMax, y}).castRay(graph, *pb)
	}

	if r == 0 {
		return
	}

	// size down the bounding box
	b = BoundingBox{b.xMin + 1, b.yMin + 1, b.xMax - 1, b.yMax - 1}
	return r + b.Count(pb, graph)
}
