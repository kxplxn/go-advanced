package _030309_usingRaceDetector_

import (
	"fmt"
	"sync"
)

func Demo() {
	fmt.Println("\n020309 Goroutines: Using Race Detector")

	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Println("final counter:", counter)
}
