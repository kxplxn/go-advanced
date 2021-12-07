package _020701_concurrency_

import "fmt"

func Demo() {
	fmt.Println("\n\n020705 Best Practices: Concurrency")

	msg := make(chan int, 10)
	done := make(chan bool)

	go func() {
		for {
			i, work := <-msg
			fmt.Println("i:", i, ", work:", work)
			if work {
				fmt.Printf("Message %d processed!\n", i)
			} else {
				fmt.Println("All messages processed!")
				done <- true // tell the main goroutine to shut down
				return
			}
		}
	}()

	for i := 1; i <= 5; i++ {
		msg <- i
		fmt.Printf("Message %d sent to be processed.\n", i)
	}

	close(msg)
	fmt.Println("Processing completed.")

	// block until something is sent over the done channel
	<-done
}
