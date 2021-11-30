package _020401_usingChannels_

import (
	"fmt"
	"reflect"
)

func Demo() {
	fmt.Println("\n020401 Channels: Goroutines and Channels")
	ch := make(chan int)
	mult := func(a, b int) {
		res := a * b
		ch <- res
	}
	go mult(10, 10)
	res, ok := <-ch
	fmt.Println("type of ch:", reflect.TypeOf(ch))
	fmt.Println("value of ch:", ch)
	fmt.Println("res:", res)
	fmt.Println("value of ok:", ok)
}
