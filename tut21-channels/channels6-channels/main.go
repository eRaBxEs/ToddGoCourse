package main

import (
	"fmt"
)

func main() {
	c := make(chan int)    // bi-directionl channel
	cr := make(<-chan int) // recieve only channel
	cs := make(chan<- int) // send only channel

	fmt.Println("------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

}
