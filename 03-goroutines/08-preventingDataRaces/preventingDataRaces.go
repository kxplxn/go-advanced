package _020308_preventingDataRaces_

import (
	"fmt"
	"sync"
)

func Demo() {
	fmt.Println("\n020308 Goroutines: Preventing Data Races")

	var counter int64
	var wg sync.WaitGroup
	var mu sync.Mutex

	inc := func() {
		mu.Lock()
		defer mu.Unlock()
		for i := 0; i < 1000; i++ {
			counter++
		}
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc()
		}()
	}

	wg.Wait()
	fmt.Println("final counter:", counter)
}
