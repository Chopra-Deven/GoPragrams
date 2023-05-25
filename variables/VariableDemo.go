package main

import (
	"fmt"
)

func main() {

	var student = "deven"

	fmt.Println("\n" + student)

	var surname = "Jain"

	_ = surname

	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	//fmt.Println("\n" + student)

	car, year := "Honda", 2022

	fmt.Println(car, year)

	//car, year := "CIty", 2022 	// here both the var are already declared that's why it is showing error

	car, hello := "City", 2021 // there should be atleast one var unique

	fmt.Println(car, hello)

	var (
		a         int
		firstName string
		gender    bool
	)
	_, _, _ = a, firstName, gender

	var i, j int

	i, j = 5, 10

	fmt.Print("\nbefore swap : ")
	fmt.Print(i, j)

	i, j = j, i
	fmt.Print("\nafter swap : ")
	fmt.Print(i, j)

	const pi = 3.14

	const (
		min1 = 100
		min2
	)
}
