package main

import "fmt"

var a int

func into2(a *int) int {

	*a = *a * 2

	return *a
}

func sum2(a *int) int {
	defer into2(a)
	*a += *a

	return *a
}

func main() {

	a = 10

	sum2(&a)

	fmt.Println("a :", a)

}
