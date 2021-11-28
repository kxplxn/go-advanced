package _020301_goroutines_

import (
	"fmt"
	"time"
)

func Demo() {
	fmt.Println("\n020301 Goroutines: Goroutines")

	fmt.Println("running in main routine")
	g1()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("exiting from main routine")
}

func g1() {
	fmt.Println("running in g1 routine")
}
