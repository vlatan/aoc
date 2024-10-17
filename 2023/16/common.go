package day16

import (
	"aoc/2023/utils"
	"bufio"
	"fmt"
	"os"
)

type Loc struct {
	x, y int
}

type Matrix [][]byte

func parseFile(path string) (r Matrix) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, []byte(scanner.Text()))
	}
	return
}

func solve(matrix Matrix, start Loc, comingFrom string) int {
	energized := make(map[Loc]struct{})
	visited := make(map[string]struct{})

	var recurse func(comingFrom string, curr Loc)
	recurse = func(comingFrom string, curr Loc) {

		// check the bounds
		if curr.x < 0 || curr.x > len(matrix)-1 ||
			curr.y < 0 || curr.y > len(matrix[0])-1 {
			return
		}

		// mark as energized
		energized[curr] = struct{}{}

		symbol := matrix[curr.x][curr.y]
		left, right := Loc{curr.x, curr.y - 1}, Loc{curr.x, curr.y + 1}
		up, down := Loc{curr.x - 1, curr.y}, Loc{curr.x + 1, curr.y}

		switch comingFrom {

		case "left":
			switch symbol {
			case '.', '-':
				key := fmt.Sprintf("h.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("left", right)
				}
			case '/':
				key := fmt.Sprintf("l.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("down", up)
				}
			case '\\':
				key := fmt.Sprintf("l.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("up", down)
				}
			case '|':
				key := fmt.Sprintf("l.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("up", down)
					recurse("down", up)
				}
			}

		case "up":
			switch symbol {
			case '.', '|':
				key := fmt.Sprintf("v.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("up", down)
				}
			case '/':
				key := fmt.Sprintf("l.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("right", left)
				}
			case '\\':
				key := fmt.Sprintf("r.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("left", right)
				}
			case '-':
				key := fmt.Sprintf("u.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("right", left)
					recurse("left", right)
				}
			}

		case "right":
			switch symbol {
			case '.', '-':
				key := fmt.Sprintf("h.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("right", left)
				}
			case '/':
				key := fmt.Sprintf("r.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("up", down)
				}
			case '\\':
				key := fmt.Sprintf("r.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("down", up)
				}
			case '|':
				key := fmt.Sprintf("r.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("down", up)
					recurse("up", down)
				}
			}

		case "down":
			switch symbol {
			case '.', '|':
				key := fmt.Sprintf("v.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("down", up)
				}
			case '/':
				key := fmt.Sprintf("r.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("left", right)
				}
			case '\\':
				key := fmt.Sprintf("l.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("right", left)
				}
			case '-':
				key := fmt.Sprintf("d.%v", curr)
				if _, ok := visited[key]; !ok {
					visited[key] = struct{}{}
					recurse("right", left)
					recurse("left", right)
				}
			}
		}
	}

	recurse(comingFrom, start)
	return len(energized)
}
