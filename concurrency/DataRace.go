package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100

	// declaring a WaitGroup to synchronize the goroutines with the main function.
	var wg sync.WaitGroup

	// adding 200 goroutines to the WaitGroup
	wg.Add(gr * 2)

	var mutex sync.Mutex

	// declaring a shared value
	var n = 0

	// starting 200 goroutines
	for i := 0; i < gr; i++ {

		// each goroutine is an anonymous function
		go func() {

			defer wg.Done()
			defer mutex.Unlock()

			// introducing some artificial time
			time.Sleep(time.Second / 10)

			mutex.Lock()
			// increment the shared value
			n++

		}()

		// goroutine that decrements the shared value
		go func() {
			defer wg.Done()
			defer mutex.Unlock()

			time.Sleep(time.Second / 10)

			mutex.Lock()
			n--
		}()

	}
	// waiting for the goroutines to terminate.
	wg.Wait()

	//  printing the final value of n
	fmt.Println(n) // it will have another value for each program execution -> DATA RACE
}
