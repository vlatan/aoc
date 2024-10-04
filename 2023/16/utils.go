package day16

import (
	"aoc/2023/utils"
	"bufio"
	"fmt"
	"os"
)

type Matrix [][]byte

func parseFile(path string) (r Matrix) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for x := 0; scanner.Scan(); x++ {
		r = append(r, []byte(scanner.Text()))
	}
	return
}

func solve(matrix Matrix, start Loc, comingFrom string) int {
	energized := make(map[Loc]struct{})
	visited := make(map[string]struct{})

	var recurse func(comingFrom string, curr Loc)
	recurse = func(comingFrom string, curr Loc) {

		if curr.x < 0 || curr.x > len(matrix)-1 ||
			curr.y < 0 || curr.y > len(matrix[0])-1 {
			return
		}

		energized[curr] = struct{}{}

		switch comingFrom {

		case "left":
			switch matrix[curr.x][curr.y] {
			case '.', '-':
				if _, ok := visited[fmt.Sprintf("h.%v", curr)]; !ok {
					visited[fmt.Sprintf("h.%v", curr)] = struct{}{}
					recurse("left", Loc{curr.x, curr.y + 1})
				}
			case '/':
				if _, ok := visited[fmt.Sprintf("l.%v", curr)]; !ok {
					visited[fmt.Sprintf("l.%v", curr)] = struct{}{}
					recurse("down", Loc{curr.x - 1, curr.y})
				}
			case '\\':
				if _, ok := visited[fmt.Sprintf("l.%v", curr)]; !ok {
					visited[fmt.Sprintf("l.%v", curr)] = struct{}{}
					recurse("up", Loc{curr.x + 1, curr.y})
				}
			case '|':
				if _, ok := visited[fmt.Sprintf("l.%v", curr)]; !ok {
					visited[fmt.Sprintf("l.%v", curr)] = struct{}{}
					recurse("up", Loc{curr.x + 1, curr.y})
					recurse("down", Loc{curr.x - 1, curr.y})
				}
			}

		case "up":
			switch matrix[curr.x][curr.y] {
			case '.', '|':
				if _, ok := visited[fmt.Sprintf("v.%v", curr)]; !ok {
					visited[fmt.Sprintf("v.%v", curr)] = struct{}{}
					recurse("up", Loc{curr.x + 1, curr.y})
				}
			case '/':
				if _, ok := visited[fmt.Sprintf("l.%v", curr)]; !ok {
					visited[fmt.Sprintf("l.%v", curr)] = struct{}{}
					recurse("right", Loc{curr.x, curr.y - 1})
				}
			case '\\':
				if _, ok := visited[fmt.Sprintf("r.%v", curr)]; !ok {
					visited[fmt.Sprintf("r.%v", curr)] = struct{}{}
					recurse("left", Loc{curr.x, curr.y + 1})
				}
			case '-':
				if _, ok := visited[fmt.Sprintf("u.%v", curr)]; !ok {
					visited[fmt.Sprintf("u.%v", curr)] = struct{}{}
					recurse("right", Loc{curr.x, curr.y - 1})
					recurse("left", Loc{curr.x, curr.y + 1})
				}
			}

		case "right":
			switch matrix[curr.x][curr.y] {
			case '.', '-':
				if _, ok := visited[fmt.Sprintf("h.%v", curr)]; !ok {
					visited[fmt.Sprintf("h.%v", curr)] = struct{}{}
					recurse("right", Loc{curr.x, curr.y - 1})
				}
			case '/':
				if _, ok := visited[fmt.Sprintf("r.%v", curr)]; !ok {
					visited[fmt.Sprintf("r.%v", curr)] = struct{}{}
					recurse("up", Loc{curr.x + 1, curr.y})
				}
			case '\\':
				if _, ok := visited[fmt.Sprintf("r.%v", curr)]; !ok {
					visited[fmt.Sprintf("r.%v", curr)] = struct{}{}
					recurse("down", Loc{curr.x - 1, curr.y})
				}
			case '|':
				if _, ok := visited[fmt.Sprintf("r.%v", curr)]; !ok {
					visited[fmt.Sprintf("r.%v", curr)] = struct{}{}
					recurse("down", Loc{curr.x - 1, curr.y})
					recurse("up", Loc{curr.x + 1, curr.y})
				}
			}

		case "down":
			switch matrix[curr.x][curr.y] {
			case '.', '|':
				if _, ok := visited[fmt.Sprintf("v.%v", curr)]; !ok {
					visited[fmt.Sprintf("v.%v", curr)] = struct{}{}
					recurse("down", Loc{curr.x - 1, curr.y})
				}
			case '/':
				if _, ok := visited[fmt.Sprintf("r.%v", curr)]; !ok {
					visited[fmt.Sprintf("r.%v", curr)] = struct{}{}
					recurse("left", Loc{curr.x, curr.y + 1})
				}
			case '\\':
				if _, ok := visited[fmt.Sprintf("l.%v", curr)]; !ok {
					visited[fmt.Sprintf("l.%v", curr)] = struct{}{}
					recurse("right", Loc{curr.x, curr.y - 1})
				}
			case '-':
				if _, ok := visited[fmt.Sprintf("d.%v", curr)]; !ok {
					visited[fmt.Sprintf("d.%v", curr)] = struct{}{}
					recurse("right", Loc{curr.x, curr.y - 1})
					recurse("left", Loc{curr.x, curr.y + 1})
				}
			}
		}
	}

	recurse(comingFrom, start)
	return len(energized)
}
