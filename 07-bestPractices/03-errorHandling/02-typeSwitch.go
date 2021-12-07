package _020703_errorHandling_

import "fmt"

type MessageLogError struct{}

// implement the `error` interface
func (e *MessageLogError) Error() string {
	return "message is received but not logged"
}

type TCPError struct{}

func (e *TCPError) Error() string {
	return "network connection is refused"
}

func SaveMessageError(val int) error {
	if val == 1 {
		return &TCPError{}
	} else if val == 2 {
		return &MessageLogError{}
	} else {
		return nil
	}
}

func DemoTypeSwitch() {
	fmt.Println("\n020703 Best Practices: Error Handling — 02 Type Switch")

	switch err := SaveMessageError(1); err.(type) {
	case nil:
		fmt.Println("Message received and logged succesfuly.")
	case *TCPError:
		fmt.Println("Network Error:", err)
	case *MessageLogError:
		fmt.Println("Message Log Error:", err)
	}
}
