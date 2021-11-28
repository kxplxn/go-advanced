package _020201_improvingPerformance_

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Running a billion "add"s sequentially vs concurrently.

func Demo() {
	fmt.Println("\n0202001 Concurrency: Improving Performance")
	vals := gen(1e9)

	var sum int64
	t := time.Now()
	var sqtime time.Duration
	var cntime time.Duration

	for i := 1; i < 6; i++ {
		fmt.Println("\nrun number:", i)

		t = time.Now()
		sum = sqadd(vals)
		sqtime = time.Since(t)
		fmt.Printf("Sequential operation time: %v. Total sum: %v.\n", sqtime, sum)

		t = time.Now()
		sum = cnadd(vals)
		cntime = time.Since(t)
		fmt.Printf("Concurrent operation time: %v. Total sum: %v.\n", cntime, sum)

		fmt.Printf("Concurrent operation took %v less time to compute sum!\n", sqtime-cntime)
		fmt.Printf("Concurrent operation took %.2f%% as much time to compute sum.\n", 100*float64(cntime)/float64(sqtime))
	}
}

// gen generates random values to work with
func gen(max int) []int {
	rand.Seed(time.Now().UnixNano())
	vals := make([]int, max)
	for i := 0; i < max; i++ {
		vals[i] = rand.Intn(10)
	}
	return vals
}

// sqadd adds values sequentially
func sqadd(vals []int) int64 {
	var sum int64
	for _, n := range vals {
		sum += int64(n)
	}
	return sum
}

// cnadd adds values concurrently
func cnadd(vals []int) int64 {
	var sum int64

	// determine the number of cores on this device
	cores := runtime.NumCPU()
	fmt.Println("Cores:", cores)
	// set the max no. of processors
	runtime.GOMAXPROCS(cores)

	// split the length of values to number of cores
	max := len(vals)
	splits := max / cores

	var wg sync.WaitGroup

	for i := 0; i < cores; i++ {
		// split the input up for concurrent operation
		start := i * splits
		end := start + splits
		split := vals[start:end]

		// use seperate goroutines to compute the sum
		wg.Add(1)
		go func(vals []int) {
			defer wg.Done()

			var splitSum int64

			// complete the sum for each split part of output
			for _, n := range vals {
				splitSum += int64(n)
			}

			atomic.AddInt64(&sum, splitSum)
		}(split)
	}

	wg.Wait()
	return sum
}
