package _020402_bufferedChannels_

import (
	"fmt"
	"sync"
)

func Demo() {
	fmt.Println("\n020402 Channels: Buffered and Unbuffered — 02 Buffered Channels")

	ch := make(chan string, 3)
	var wg sync.WaitGroup

	fmt.Println("initial channel length:", len(ch))

	wg.Add(1)
	go func() {
		defer wg.Done()
		send(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		receive(ch)
	}()

	wg.Wait()
}

func send(ch chan string) {
	fmt.Println("sending...")
	for i := 1; i <= 3; i++ {
		msg := fmt.Sprintf("message%v", i)
		ch <- msg
		fmt.Println("sent:", msg)
		fmt.Println("channel length:", len(ch))
	}
}

func receive(ch chan string) {
	fmt.Println("receiving...")
	for i := 0; i < 3; i++ {
		fmt.Println("received:", <-ch)
		fmt.Println("channel length:", len(ch))
	}
}
