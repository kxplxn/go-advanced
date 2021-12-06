package _020601_debugging_

import "fmt"

func Demo() {
	fmt.Println("\n020601 Testing and Deploying: Debugging")

	x := 1
	y := 2
	fmt.Println("adding", x, "and", y)
	fmt.Println("result from operation:", Add(x, y))
}

func Add(x, y int) int {
	res := x + y
	//res += 1
	return res
}
