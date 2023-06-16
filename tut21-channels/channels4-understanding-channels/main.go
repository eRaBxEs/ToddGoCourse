package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) // buffered ; giving room for one value

	c <- 42
	c <- 43
	// deadlock will occur here because there is only room for one value to be sent to the channel as indicated by the buffer above

	fmt.Println(<-c)
}
