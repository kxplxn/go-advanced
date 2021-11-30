package _020305_atomicOperations_

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func Demo() {
	fmt.Println("\n020305 Goroutines: Atomic Operations")

	runtime.GOMAXPROCS(runtime.NumCPU())
	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}
