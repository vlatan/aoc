package seed

import "fmt"

// https://adventofcode.com/2023/day/5
func SeedPartTwo() {
	seeds, maps := parseFile("05/input.txt")
	for loc := 0; ; loc++ {
		x := loc
		for m := len(maps) - 1; m >= 0; m-- {
			x = maps[m].InverseConvert(x)
		}
		for s := 0; s < len(seeds); s += 2 {
			if x >= seeds[s] && x < seeds[s]+seeds[s+1] {
				fmt.Println(loc)
				return
			}
		}
	}
}

func (m Map) InverseConvert(num int) int {
	for _, line := range m {
		if line.dest <= num && num < line.dest+line.len {
			return line.src + (num - line.dest)
		}
	}
	return num
}
