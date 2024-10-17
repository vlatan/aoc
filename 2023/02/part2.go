package day02

import (
	"aoc/2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2023/day/2
func Part2() {
	file, err := os.Open("02/input.txt")
	utils.Check(err)
	defer file.Close()

	result, scanner := 0, bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")
		gameSets := strings.Split(game[1], "; ")
		minCubes := map[string]int{"red": 0, "green": 0, "blue": 0}
		setResult := 1
		for _, set := range gameSets {
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				cubeInfo := strings.Split(cube, " ")
				numCubes := utils.ToInteger(cubeInfo[0])
				cubeColor := cubeInfo[1]
				if minCubes[cubeColor] < numCubes {
					minCubes[cubeColor] = numCubes
				}
			}
		}
		for _, numCubes := range minCubes {
			setResult *= numCubes
		}
		result += setResult
	}
	fmt.Println("Part 2:", result)
}
