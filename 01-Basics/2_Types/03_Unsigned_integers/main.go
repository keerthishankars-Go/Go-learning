// Unsigned integer types can only be used to store positive integers
// uint8, uint16, uint32, uint64, uint are the unsigned integer data types available in Go.
// Unsigned integers are used in places where negative values are not applicable.
package main

import "fmt"

func main() {
	var a uint = 60
	var b uint = 43

	c := a * b
	// var a uint = 60
	// var b int = 43

	// c := a * uint(b)
	fmt.Println("c =", c)
	fmt.Printf("Data type of variable c is %T ", c)
}

