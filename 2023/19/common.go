package day19

import (
	"aoc/2023/common"
	"bufio"
	"os"
	"strings"
)

type Workflow []string
type Workflows map[string]Workflow
type Rating map[string]int

func parseFile(path string) (Workflows, []Rating) {
	file, err := os.Open(path)
	common.Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	workflows, ratings := Workflows{}, []Rating{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if line[0] == '{' {
			rating := Rating{}
			line = line[1 : len(line)-1]
			lineParts := strings.Split(line, ",")
			for _, item := range lineParts {
				itemParts := strings.Split(item, "=")
				rating[itemParts[0]] = common.ToInteger(itemParts[1])
			}
			ratings = append(ratings, rating)
		} else {
			lineParts := strings.Split(line, "{")
			lineParts[1] = lineParts[1][:len(lineParts[1])-1]
			workflows[lineParts[0]] = strings.Split(lineParts[1], ",")
		}
	}
	return workflows, ratings

}
