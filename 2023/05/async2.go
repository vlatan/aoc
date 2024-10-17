package day05

import (
	"fmt"
	"sync"
)

// https://adventofcode.com/2023/day/5
// Finishes in close to a minute.
func SeedAsyncPartTwo() {

	seeds, maps := parseFile("05/input.txt")

	var wg sync.WaitGroup
	// buffered channel with length of seed/2 (num of goroutines)
	ch := make(chan int, len(seeds)/2)
	for s := 0; s < len(seeds); s += 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processRange(seeds[s], seeds[s]+seeds[s+1], maps, ch)
		}()
	}

	// Close the buffered channel when full, when all goroutines are finished
	// If you're not using a buffered channel meaning you don't know exactly how
	// many results will you send to, then you need
	// to wrap this into a goroutine which will wait in a separate
	// thread and close the channel when all needed goroutines are done
	// which will signal the channel loop below to terminate.
	wg.Wait()
	close(ch)

	minLoc := 1<<63 - 1
	for loc := range ch {
		minLoc = min(minLoc, loc)
	}
	fmt.Println(minLoc)
}

func processRange(start int, end int, maps Maps, ch chan int) {
	minLoc := 1<<63 - 1
	for s := start; s < end; s++ {
		loc := s
		for _, m := range maps {
			loc = m.Convert(loc)
		}
		minLoc = min(minLoc, loc)
	}
	ch <- minLoc
}
