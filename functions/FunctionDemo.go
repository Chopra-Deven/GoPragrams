package main

import (
	"fmt"
	"math"
)

var p = fmt.Println

// defining a function with no parameters
func f11() {
	fmt.Println("This is f1() function")
}

// defining a function with 2 parameters, a and b
func f22(a int, b int) {
	//a and b are local to the function
	fmt.Println("Sum:", a+b)
}

// defining a function using shorthand parameters notation
func f33(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// defining a function that have one parameter of type float64 and returns a value of type float64
func f44(a float64) float64 {
	return math.Pow(a, a)
	//any statements below the return statement are never executed

}

// defining a function that have two parameters of type int and returns two values of type int
func f55(a, b int) (int, int) {
	return a * b, a + b
}

// defining a function that have one parameter of type int and returns a "named parameter"
func sum(a, b int) (s int) {
	fmt.Println("s:", s) // -> s is a variable with the zero value inside the function
	s = a + b

	// it automatically return s
	return // This is known as a "naked" return.
}

func swap(a, b int) {

	p("before swaping from func : a = ", a, " , b = ", b)

	a = 10
	b = 20

	p("after swaping from func : a = ", a, " , b = ", b)

}

func main() {

	// calling functions
	f11() // => This is f1() function

	f33(3, 4, 5, 4., 5.5, "ss") // => 3 4 5 4 5.5 ss

	fmt.Println(f44(5.1))

	a, b := f55(6, 7)
	fmt.Printf("a:%d, b:%d\n", a, b) // => a:42, b:13

	ss := sum(4, 5)
	fmt.Println(ss) // -> 9

	// There are no default arguments in Go //

	number1 := 1
	number2 := 2

	p("before swaping from main : a = ", number1, " , b = ", number2)

	swap(number1, number2)

	p("after swaping from main : a = ", number1, " , b = ", number2)

}
