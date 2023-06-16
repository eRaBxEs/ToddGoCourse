package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// here the go routine makes it happen,  the send and receive happen at the same time
	go func() {
		c <- 42 // this one blocks right there in this goroutine and the other go routine which is the func main is able to receive it
	}()

	fmt.Println(<-c) // for the func main; this one blocks until it takesthe values off
}
