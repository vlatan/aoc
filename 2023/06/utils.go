package wait

import (
	"aoc/2023/utils"
	"bufio"
	"os"
	"strings"
)

func parseFile(path string) ([]string, []string) {
	file, err := os.Open(path)
	utils.Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	times := strings.Fields(scanner.Text())[1:]
	scanner.Scan()
	distances := strings.Fields(scanner.Text())[1:]
	return times, distances
}
