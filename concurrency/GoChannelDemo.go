package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	n := 10

	go func(c chan int, n int) {

		for i := 0; i < n; i++ {

			//fmt.Println("Printed :", i)
			time.Sleep(time.Second)
			c <- i
		}

	}(ch, n)

	go func(c chan int, n int) {

		for i := 0; i < n; i++ {

			//fmt.Println("Printed :", i)
			time.Sleep(time.Second)
			c <- i + 100
		}

	}(ch, n)

	for j := 0; j < n*2; j++ {
		fmt.Println("Index ", j, " : ", <-ch)
	}

	//fmt.Println(<-ch)
}
