package _020403_rangeAndClose_

import (
	"fmt"
	"sync"
)

func Demo() {
	fmt.Println("\n020403 Channels: `range` and `close()`")

	ch := make(chan string, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(ch chan string) {
		defer func() {
			wg.Done()
			close(ch)
		}()

		for i := 1; i <= 5; i++ {
			msg := fmt.Sprintf("message%v", i)
			ch <- msg
			fmt.Println("SEND goroutine:", msg)
		}
	}(ch)

	wg.Add(1)
	go func(ch chan string) {
		defer wg.Done()
		for val := range ch {
			fmt.Println("RECV goroutine:", val)
		}
	}(ch)

	fmt.Println("MAIN goroutine: waiting...")
	wg.Wait()
	val, ok := <-ch
	fmt.Println("val:", val, "; ok:", ok)
	fmt.Println("MAIN goroutine: done!")
}
