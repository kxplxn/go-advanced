package _020303_returningFunctions_

import (
	"fmt"
	"reflect"
)

func Demo() {
	fmt.Println("\n02030203 Goroutines: Closures — Returning Functions")
	myfunc := retfunc()
	fmt.Println(myfunc())
	fmt.Println("myfunc is of type:", reflect.TypeOf(myfunc))
}

func retfunc() func() bool {
	return func() bool {
		return true
	}
}
