package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched() // yields the processor, allowing other goroutines to run.
			v++
			counter = v
			wg.Done() // to show that a particular goroutine is done
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait() // don't exit my program till every goroutine is done.
	// Till the counter inside waitgroup gets to zero
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count:", counter)
	// Race condition will occur here because we have multiple goroutine accessing a shared variable

}
