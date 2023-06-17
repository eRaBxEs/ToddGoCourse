package main

import (
	"fmt"
)

func main() {
	c := make(chan int)    // bi-directional channel
	cr := make(<-chan int) // receive only channel
	cs := make(chan<- int) // send only channel

	fmt.Println("-------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// specific to general doesn't assign
	c = cr
	c = cs
}
