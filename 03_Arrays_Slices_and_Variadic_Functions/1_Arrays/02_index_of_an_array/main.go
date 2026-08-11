package main

// The index of an array starts from 0 and ends at length - 1.

import (
	"fmt"
)

func main() {
	var a [3]int //int array with length 3

	a[0] = 12 // array index starts at 0
	a[1] = 56
	a[2] = 98

	fmt.Println(a)
}

// a[0] assigns value to the first element of the array. The program will print:

// [12 56 98]
