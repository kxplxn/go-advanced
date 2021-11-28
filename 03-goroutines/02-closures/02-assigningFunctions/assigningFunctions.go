package _02030202_functionAssignments

import (
	"fmt"
	"reflect"
)

func Demo() {
	defer fmt.Println("baban")
	fmt.Println("\n02030203 Goroutines: Closures — Function Variables")
	anon := func(msg string) bool {
		fmt.Println(msg)
		return true
	}
	res := anon("Hello, anonymous function!")
	fmt.Println(res)
	fmt.Println("anon is of type", reflect.TypeOf(anon))
}
