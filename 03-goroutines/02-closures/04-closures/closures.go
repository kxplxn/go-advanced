package _02030404_closures_

import (
	"fmt"
	"reflect"
)

func Demo() {
	fmt.Println("\n02030204 Goroutines: Closures — Closures")

	x := inc()
	fmt.Println("x is of type", reflect.TypeOf(x))

	for i := 0; i < 4; i++ {
		fmt.Println("x =", x())
	}

	y := inc()
	fmt.Println("y =", y())
}

func inc() func() int {
	j := 0
	return func() int {
		j++
		return j
	}
}
