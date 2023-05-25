package main

import (
	"fmt"
)

type data map[string]string

func (data data) print() {

	for k, v := range data {
		fmt.Println(k, ":", v)
	}

}

func main() {

	student1 := data{
		"name":    "deven",
		"address": "abad",
		"age":     "22",
	}
	student1.print()
}
