package _020304_waitGroups_

import (
	"fmt"
	"sync"
)

func Demo() {
	fmt.Println("\n020304 Goroutines: WaitGroups")
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go myFunc(&wg, i)
	}
	wg.Wait()
	fmt.Println("each goroutine has run to completion")
}

func myFunc(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("finished executing iteration", i)
}
