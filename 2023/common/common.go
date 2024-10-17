package common

import (
	"bufio"
	"os"
	"strconv"
)

// Check error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Open file, return scanner
func ParseFile(path string) [][]byte {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()
	var result [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var currentSlice []byte
		for i := 0; i < len(line); i++ {
			currentSlice = append(currentSlice, line[i])
		}
		result = append(result, currentSlice)
	}
	return result
}

// Convert string to integer
func ToInteger(num string) int {
	result, err := strconv.Atoi(num)
	Check(err)
	return result
}

// Check if byte is digit
func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func Sum(lst []int) int {
	result := 0
	for _, num := range lst {
		result += num
	}
	return result
}
