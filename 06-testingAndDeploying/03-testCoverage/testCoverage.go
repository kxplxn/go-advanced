package _020502_unitTesting_

import "fmt"

func Demo() {
	fmt.Println("\n020603 Testing and Deploying: Test Coverage")
	fmt.Println("[run `go test -cover` in 03-testCoverage dir for demo")
	fmt.Println("also, check out the table-driven tests in testCoverage_test]")
}

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}
