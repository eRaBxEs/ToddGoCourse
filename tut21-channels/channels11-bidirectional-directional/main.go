package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Println("-------")
	fmt.Printf("c\t%T", c)
	fmt.Printf("cr\t%T", cr)
	fmt.Printf("cs\t%T", cs)

	// conversion from specific to general doesn't convert
	fmt.Println("-------")
	fmt.Printf("cr\t%T", (chan int)(cr))
	fmt.Printf("cr\t%T", (chan int)(cs))

}
