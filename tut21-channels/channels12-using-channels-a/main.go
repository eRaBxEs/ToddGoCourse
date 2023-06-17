package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	go bar(c) // will not block in order to receive the channel as a parameter therefore might not receive anything

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42 // putting in the value
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c) // pulling the value off
}
