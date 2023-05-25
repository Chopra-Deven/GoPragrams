package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

type shape interface {
	area() float64
	perimeter() float64
}

// declaring a method for circle type
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func (c circle) area() float64 {
	return c.radius * c.radius * 3.14
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func printData(s shape) {
	fmt.Println()
	fmt.Printf("\nShape : %#v", s)
	fmt.Printf("\nArea : %v", s.area())
	fmt.Printf("\nPerimeter : %v", s.perimeter())
}

func main() {

	ball := circle{
		radius: 10,
	}

	block := rectangle{width: 5, height: 10}

	printData(ball)
	printData(block)

	var s shape = circle{radius: 2.5}

	fmt.Printf("%T\n", s) //interface dynamic type is circle

	// no direct access to interface's dynamic values
	// s.volume() -> error

	// there is access only to the methods that are defined inside the interface
	fmt.Printf("\n\nCircle Area:%v\n", s.area())

	// an interface value hides its dynamic value.
	// use type assertion to extract and return the dynamic value of the interface value.
	fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())

	// checking if the assertion succeded or not
	ball, ok := s.(circle)
	if ok == true {
		fmt.Printf("Ball Volume:%v\n", ball.volume())
	}

	//** TYPE SWITCHES **//

	// it permits several type assertions in series
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)

	}

}
