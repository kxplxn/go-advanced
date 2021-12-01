package _1_unbufferedChannels

import (
	"fmt"
	"sync"
)

func Demo() {
	fmt.Println("\n020402 Channels: Buffered and Unbuffered — 01 Unbuffered Channels")

	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		send(ch)
	}()

	fmt.Println("send goroutine is blocked")
	fmt.Println("channel length:", len(ch))

	wg.Add(1)
	go func() {
		defer wg.Done()
		receive(ch)
	}()

	wg.Wait()
}

func send(ch chan string) {
	ch <- "message"
}

func receive(ch chan string) {
	fmt.Println("channel length:", len(ch))
	fmt.Println("received:", <-ch)
	fmt.Println("send goroutine is unblocked")
}
