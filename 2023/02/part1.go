package cube

import (
	"aoc/2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var maxCubes = map[string]int{"red": 12, "green": 13, "blue": 14}

// https://adventofcode.com/2023/day/2
func CubeConundrumPartOne() {
	file, err := os.Open("02/input.txt")
	utils.Check(err)
	defer file.Close()

	result, scanner := 0, bufio.NewScanner(file)
	for scanner.Scan() {
		line, validGame := scanner.Text(), true
		game := strings.Split(line, ": ")
		sets := strings.Split(game[1], "; ")
		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				values := strings.Split(cube, " ")
				if utils.ToInteger(values[0]) > maxCubes[values[1]] {
					validGame = false
					break
				}
			}
			if !validGame {
				break
			}
		}
		if validGame {
			game = strings.Split(game[0], " ")
			result += utils.ToInteger(game[1])
		}
	}
	fmt.Println(result)
}
