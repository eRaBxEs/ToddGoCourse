package main

import (
	"fmt"
	"reflect"
)

func main() {
	c := make(chan<- int, 2) // now making the channel a send only channel, so you can only send values to it

	c <- 42
	c <- 43

	fmt.Println(<-c) // fails because one cannot receive from a send only channel
	fmt.Println(<-c) // fails because one cannot receive from a send only channel
	fmt.Println("........")
	fmt.Printf("%T\n", c)
	fmt.Println("Type :", reflect.TypeOf(c))
}
