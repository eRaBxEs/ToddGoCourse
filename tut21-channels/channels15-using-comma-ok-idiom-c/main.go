package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit) // At zero, exit the quit channel (send only)
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value in this channel is even:", v)
		case v := <-odd:
			fmt.Println("the value in this channel is odd:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("the value from ok:", i, ok)
			} else {
				fmt.Println("the value from ok:", i, ok)
			}
			// at zero, we exit the quit channel (recieve only)
			return
		}
	}
}
