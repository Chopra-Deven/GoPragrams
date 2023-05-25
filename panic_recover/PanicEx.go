package main

import (
	"fmt"
	"runtime/debug"
)

func method3() {
	panic("message---3")
}

func method2() {
	defer method3()
	panic("message---2")
}

func method1() {
	defer method2()
	panic("message---1")
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	method1()
}
