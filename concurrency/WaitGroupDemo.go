package main

import (
	"fmt"
	"sync"
	"time"
)

func f11(wg *sync.WaitGroup) {

	fmt.Println("F1 started")

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("F1 :", i)
	}
	fmt.Println("F1 ended")

	wg.Done()
}

func f22() {

	fmt.Println("F2 started")

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("F2 :", i)
	}

	fmt.Println("F2 ended")

}

func main() {

	start := time.Now()

	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println(wg)

	fmt.Println("F1 started from main")

	go f11(&wg)

	fmt.Println("F1 ended from main")

	fmt.Println("F2 started from main")

	f22()

	fmt.Println("F2 ended from main")

	fmt.Println("Main completed")

	wg.Wait()

	end := time.Since(start)

	fmt.Println("\nTime Taken :", end)

}
