package day18

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// https://adventofcode.com/2023/day/18
func Part2() {
	graph, pb := parseFile("18/input.txt", processLine2)
	result := 0
	fmt.Println("Polygon done")
	start := time.Now()
	grid := pb.Grid(1000)
	elapsed := time.Since(start)
	log.Printf("Grid made in %s.", elapsed)
	for b := range grid {
		start = time.Now()
		result += b.Count(&pb, graph)
		elapsed = time.Since(start)
		log.Printf("Grid counted in %s", elapsed)

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
