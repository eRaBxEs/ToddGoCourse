package main

import (
	"fmt"
)

func main() {
	c := make(<-chan int, 2) // receive only channel

	c <- 42 // cannot send to a receive only channel
	c <- 43 // cannot send to a receive only channel

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("......")
	fmt.Printf("%T\n", c)
}
