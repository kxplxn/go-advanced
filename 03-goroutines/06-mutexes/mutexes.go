package _020306_mutexes_

import (
	"fmt"
	"sync"
)

func Demo() {
	fmt.Println("\n020306 Goroutines: Mutexes")

	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	inc := func() {
		mu.Lock()
		defer mu.Unlock()
		counter++
	}

	dec := func() {
		mu.Lock()
		defer mu.Unlock()
		counter--
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dec()
		}()
	}

	wg.Wait()
	fmt.Println("final counter:", counter)
}
