package _020405_nonblockingSelect_

import (
	"fmt"
	"time"
)

func Demo() {
	fmt.Println("\n020405 Channels: Non-Blocking Select")

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(3 * time.Millisecond)
			ch1 <- i
			time.Sleep(1 * time.Millisecond)
			ch2 <- i
		}
	}()

	for i := 1; i < 5; i++ {
		select {
		case val := <-ch1:
			fmt.Println("value received from ch1:", val)
		case val := <-ch2:
			fmt.Println("value received from ch2:", val)
		default:
			fmt.Println("no data received... performing some other operation...")
		}
		time.Sleep(2 * time.Millisecond)
	}
}
