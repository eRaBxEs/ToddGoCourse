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

	fmt.Println("About to exit")
}

// send
func send(even, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok", i, ok)
			} else {
				fmt.Println("from comma ok", i, ok)
			}
			return // N.B without this return, you will never get out of the "from comma ok 0 false" loop
		}
	}
}
