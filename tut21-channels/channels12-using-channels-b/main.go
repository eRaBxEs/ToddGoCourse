package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c) // will block in order to receive the channel as a parameter because it is not using a goroutine

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
