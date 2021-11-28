package _020303_goroutinesAsClosures_

import (
	"fmt"
	"time"
)

func Demo() {
	fmt.Println("\n020302 Goroutines: Goroutines As Closures")

	fmt.Println("running goroutine in a for loop WITHOUT parameters")
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("i =", i)
		}()
	}
	time.Sleep(10 * time.Millisecond)

	fmt.Println("running goroutine in a for loop WITH parameters")
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println("i =", i)
		}(i)
	}
	time.Sleep(10 * time.Millisecond)
}
