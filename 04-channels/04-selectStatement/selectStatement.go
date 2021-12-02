package _020404_selectStatement_

import (
	"fmt"
	"time"
)

func Demo() {
	fmt.Println("\n020404 Channels: Select Statement")

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- 2
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch3 <- 3
	}()

	for i := 0; i < 3; i++ {
		select {
		case val1 := <-ch1:
			fmt.Println("value received from ch1:", val1)
		case val2 := <-ch2:
			fmt.Println("value received from ch2:", val2)
		case val3 := <-ch3:
			fmt.Println("value received from ch3:", val3)
		case to := <-time.After(100 * time.Millisecond):
			fmt.Println("timed out after 1.25 seconds at", to.Format(time.ANSIC))
			return
		}
	}
}
