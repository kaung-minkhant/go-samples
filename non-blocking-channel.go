package main

import "fmt"

func main() {
	messages := make(chan string)
	sig := make(chan bool)

	go func() {
		messages <- "hello"
	}()

	fmt.Println("Message Received", <-messages)

	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent Message", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	case signals := <-sig:
		fmt.Println("Received signals", signals)
	default:
		fmt.Println("No activity")
	}
}
