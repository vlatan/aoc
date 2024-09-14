package day12

import (
	"fmt"
	"slices"
	"strings"
)

// Every '?' can be '.' or '#'.
// Determine the indexes of the '?' in the string
// Produce the permutations of them being populated with '.' or/and '#'
// The number of the permutations should be 2^n, whene n is the number of '?'
// Try every permutation and check if the string satisfies the contiguous group criteria.
// If so count the permutation as a valid arrangement.
// Take the sum of all posible arangements.

// https://adventofcode.com/2023/day/12
func Part1() {
	data := parseFile("12/input.txt")

	result := 0
	for str, groups := range data {
		// prepare groups
		gs := []string{}
		for _, num := range groups {
			group := ""
			for i := 0; i < num; i++ {
				group += "#"
			}
			gs = append(gs, group)
		}

		// see where all the '?' are at
		indexes := []int{}
		for i, b := range str {
			if b == '?' {
				indexes = append(indexes, i)
			}
		}

		// generate permutations of '.' and '#' with length len(indexes)
		values := []byte{'.', '#'}
		ch := genPermutations(values, len(indexes))

		for permutation := range ch {
			bs := []byte(str)
			// construct a slice according to the permutation
			for i, index := range indexes {
				bs[index] = permutation[i]
			}

			bss := []string{}
			for _, item := range strings.Split(string(bs), ".") {
				if item != "" {
					bss = append(bss, item)
				}
			}

			if slices.Equal(bss, gs) {
				// fmt.Println(string(bs), gs)
				result++
			}
		}
	}
	fmt.Println(result)
}

// Generate permutations of values in a slice of length
// Feed each permutation in a channel
func genPermutations(values []byte, length int) chan []byte {
	ch, k := make(chan []byte), len(values)

	go func(ch chan []byte) {
		defer close(ch)
		p, pn := make([]byte, length), make([]int, length)
		for {
			// generate permutaton
			for i, x := range pn {
				p[i] = values[x]
			}

			// push permutation to channel
			ch <- slices.Clone(p)

			// increment permutation number
			for i := 0; ; {
				pn[i]++
				if pn[i] < k {
					break
				}
				pn[i] = 0
				i++
				if i == length {
					return // all permutations generated
				}
			}
		}
	}(ch)

	return ch
}
