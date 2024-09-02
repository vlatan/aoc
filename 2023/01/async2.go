package trebuchet

import (
	"aoc/2023/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var digitWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type safeSum struct {
	m sync.Mutex
	v int
}

func (s *safeSum) add(num int) {
	s.m.Lock()
	defer s.m.Unlock()
	s.v += num
}

func (s *safeSum) String() string {
	return fmt.Sprintf("%d", s.v)
}

// Implementation with goroutines, sync.Mutex and sync.WaitGroup
// https://adventofcode.com/2023/day/1
func TrebuchetASyncPartTwo() {
	file, err := os.Open("01/input.txt")
	utils.Check(err)
	defer file.Close()

	var sum safeSum
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go firstNum(line, &sum, &wg)
		wg.Add(1)
		go lastNum(line, &sum, &wg)
	}

	wg.Wait()
	fmt.Println(&sum)
}

func firstNum(s string, sum *safeSum, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(s); i++ {
		if utils.IsDigit(s[i]) {
			sum.add(int(s[i]-'0') * 10)
			return
		}

		for j, d := range digitWords {
			if strings.HasSuffix(s[:i+1], d) {
				sum.add((j + 1) * 10)
				return
			}
		}
	}
}

func lastNum(s string, sum *safeSum, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := len(s) - 1; i >= 0; i-- {
		if utils.IsDigit(s[i]) {
			sum.add(int(s[i] - '0'))
			return
		}

		for j, d := range digitWords {
			if strings.HasPrefix(s[i:], d) {
				sum.add(j + 1)
				return
			}
		}
	}
}
