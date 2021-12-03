package _020406_channelsAndWaitGroups_

import (
	"fmt"
	"sync"
)

func Demo() {
	fmt.Println("\n020406 Channels: Channels and WaitGroups")

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(3)

	go func() {
		wg.Wait()
		close(ch)
	}()

	go func() {
		defer wg.Done()
		ch <- 1
	}()

	go func() {
		defer wg.Done()
		ch <- 2
	}()

	go func() {
		defer wg.Done()
		ch <- 3
	}()

	for val := range ch {
		fmt.Println("value received:", val)
	}
	val, ok := <-ch
	fmt.Println("val:", val, "; ok:", ok)
}
