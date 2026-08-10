// To fix the error, both a and b must be of the same type. Let’s convert b to int. T(v) is the syntax to convert a value v to type T .

package main

import (
	"fmt"
)

func main() {
	a := 23
	b := 77.2

	sum := a + int(b)

	//sum := float64(a) + b

	fmt.Println("Sum of a and b is", sum)
}

// Since b is converted from float to int, its floating point will be truncated and hence we see 100 as the output.
