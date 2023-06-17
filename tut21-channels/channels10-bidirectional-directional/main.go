package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive only channel
	cs := make(chan<- int) // send only channel

	fmt.Println("-------")
	fmt.Printf("c\t%T", c)
	fmt.Printf("cr\t%T", cr)
	fmt.Printf("cs\t%T", cs)

	// general to specific converts
	fmt.Print("------")
	fmt.Printf("c\t%T", (<-chan int)(c)) // convert to a receive only
	fmt.Printf("c\t%T", (chan<- int)(c)) // convert to a send only

}
