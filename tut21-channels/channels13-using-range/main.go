package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// send with goroutine
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // closing the channel to avoid the range blocking
		// N.B : You can only close a send only channel
	}()

	// receive with no goroutine
	for v := range c {
		fmt.Println(v)
	}
}
