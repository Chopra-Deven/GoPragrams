package main

import "fmt"

func main() {

	var s = make([]int, 2, 5)

	s[0] = 1

	s[1] = 2

	//s[2] = 3

	fmt.Println("Slice : ", s)

	var s2 = []int{2, 4, 5, 6, 7, 8}

	fmt.Println("Appended : ", append(s, s2...))

}
