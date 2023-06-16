package main

import (
	"fmt"
)

func main() {
	// make it a buffered channel; which allows something to sit in there regardless of whether or not something is ready to pull it off
	c := make(chan int, 1) // succesful buffer

	c <- 42

	fmt.Println(<-c)
}
