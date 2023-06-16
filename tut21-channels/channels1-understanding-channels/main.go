package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	c <- 42 // pass 42 into channel c
	/**
	when you send an receive on a channel, it's like
	a relay race in a track that they have to pass the baton hand to hand
	and the transaction cannot occur until both send and receive happen at the same time
	and if they can't happen at the same time, it blocks, the send and receive blocks until
	until the receive is ready to pull it off.
	**/

	fmt.Println(<-c)
}
