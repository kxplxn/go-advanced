package _020703_errorHandling_

import (
	"fmt"
	"log"
)

func startPanic() {
	panic("panic occured in startPanic()")
}

func DemoPanic() {
	fmt.Println("\n020703 Best Practices: Error Handling — 01 Panic")

	role := "user"

	defer func(string) {
		if r := recover(); r != nil {
			switch role {
			case "admin":
				log.Println("recovered from panic:", r)
			default:
				log.Println("operation could not be completed, contact admin.")
			}
		}
	}(role)

	startPanic()

	fmt.Println("this is never reached")
}
