package day05

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

type Maps []Map

// one map is a list of map lines
type Map []Line

type Line struct {
	dest, src, len int
}

func parseFile(path string) ([]int, Maps) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// scan the first line of the file and extract seeds
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	seeds := make([]int, len(fields[1:]))
	for i, v := range fields[1:] {
		seeds[i] = utils.ToInteger(v)
	}

	// get the maps in a slice
	maps, i := make(Maps, 7), -1
	for scanner.Scan() {
		line := scanner.Text()

		// skip empty row
		if line == "" {
			continue
		}

		// there's a map title, the map begins on the next line
		if strings.Contains(line, "map:") {
			i++
			continue
		}

		// scoop a line, a map instruction
		fields := strings.Fields(line)
		dest := utils.ToInteger(fields[0])
		src := utils.ToInteger(fields[1])
		len := utils.ToInteger(fields[2])
		maps[i] = append(maps[i], Line{dest, src, len})
	}

	return seeds, maps
}
