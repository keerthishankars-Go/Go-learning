// float32, float64 are float data types used in Go .
package main

import (
	"fmt"
)

func main() {
	a, b := 8.23, 6.78
	fmt.Printf("type of a is %T, b is %T\n", a, b)

	sum := a + b  // We add a and b and assign it to a variable sum.
	diff := a - b // We subtract b from a and assign it to diff.

	fmt.Printf("sum of %f and %f is %f and difference is %f\n", a, b, sum, diff)

	no1, no2 := 56, 67
	fmt.Printf("type of no1 %T no2 %T\n", no1, no2)

	fmt.Printf("sum of %d and %d is %d, diff is %d", no1, no2, no1+no2, no1-no2)
}

//The type of a and b is inferred from the value assigned to them. In this case a and b are of type float64. float64 is the default type for floating point values.
