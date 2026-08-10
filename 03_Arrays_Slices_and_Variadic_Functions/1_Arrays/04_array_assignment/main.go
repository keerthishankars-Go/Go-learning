package main  

// It is not necessary that all elements in an array have to be assigned a value during short hand declaration.

import (
	"fmt"
)

func main () {
	a := [3]int{23}
	fmt.Println(a)
}

// Line no. 10 of the above program declares an array of length 3 but adds one element 23 to it. The remaining 2 elements are assigned 0 automatically. 

// This program will print:

// [12 0 0]