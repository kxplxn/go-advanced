package _020307_raceConditions_

import (
	"fmt"
	"runtime"
	"sync"
)

func Demo() {
	fmt.Println("\n020307 Goroutines: Race Conditions")

	runtime.GOMAXPROCS(runtime.NumCPU())

	var counter int
	var wg sync.WaitGroup

	inc := func() {
		counter++
	}

	dec := func() {
		counter--
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dec()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc()
		}()
	}

	wg.Wait()
	fmt.Println("expected final counter: 0")
	fmt.Println("actual final counter:", counter)
}
