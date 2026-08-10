package main  

// The length of the array is found by passing the array as parameter to the built-in len function.

import "fmt"

func main () {
	a := [...]float64{66.8, 52.5, 99, 23.1}
	fmt.Println("length of a is ", len(a))
}