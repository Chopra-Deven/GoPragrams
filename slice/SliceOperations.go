package main

import (
	"fmt"
	"reflect"
)

func main() {

	var cities = []string{"Abad", "Surat", "Banglore", "Pune", "Mumbai", "Delhi"}

	fmt.Println("\nCities : ", reflect.ValueOf(cities), "\nSize : ", len(cities))

	fmt.Println("\nCities are : ", cities)

	var updatedCities = removeElement(cities, 2)

	fmt.Print("Updated Cities are : ", updatedCities, "\n")

	a := []int{}

	fmt.Println(a == nil)

	n1 := []int{2, 3}
	n2 := n1

	_, _ = n1, n2
	//fmt.Println(n1 == n2)

}

func removeElement(slice []string, index int) []string {

	return append(slice[:index], slice[index+1:]...)

}
