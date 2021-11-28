package _02030201_anonymousFunctions

import "fmt"

func Demo() {
	fmt.Println("\n02030201 Goroutines: Closures — Anonymous Functions")
	func() {
		fmt.Println("hello from anonymous function")
	}()
}
