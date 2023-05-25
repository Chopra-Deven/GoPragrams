package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// a slice expression doesn't create a new backing array. The original and the returned slice are connected!
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3] //s3, s4 share the same backing array with s1

	s3[1] = 600             // modifying the backing array so s1, s3 and s4 are in fact modified!!
	fmt.Println("S1 :", s1) // -> [10 600 30 40 50]
	fmt.Println("S3 :", s3) // -> [10,600]
	fmt.Println("S4 :", s4) // -> [600 30]

	fmt.Printf("\ns1[0:5] : %p", &s1)
	fmt.Printf("\ns3[0:2] : %p", &s3)
	fmt.Printf("\ns4[1:3] : %p\n", &s4)

	// when a slice is created by slicing an array, that array becomes the backing array of the new slice.
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2                       // modifying the array
	fmt.Println(arr1, slice1, slice2) // -> [10 2 30 40] [10 2] [2 30]

	// append() function creates a complete new slice from an existing slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	var newCars []string

	// newCars doesn't share the same backing array with cars
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"                                  // only cars is modified
	fmt.Println("\ncars:", cars, "\nnewCars:", newCars) // => cars: [Nissan Honda Audi Range Rover] newCars: [Ford Honda

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("\narray's size in bytes : %d \n", unsafe.Sizeof(a))

	fmt.Printf("\nslice's size in bytes : %d \n", unsafe.Sizeof(s))

}
