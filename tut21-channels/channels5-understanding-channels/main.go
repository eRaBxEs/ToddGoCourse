package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2) // buffered channel that provides room for 2 values

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

}
