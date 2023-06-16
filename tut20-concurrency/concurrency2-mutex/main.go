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

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // locked so no body can access the variable
			v := counter
			runtime.Gosched() // yields the processor, allowing other goroutines to run.
			v++
			counter = v
			fmt.Println("Counter:\t", counter)
			mu.Unlock() // to unlock it
			wg.Done()   // to show that a particular goroutine is done
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
