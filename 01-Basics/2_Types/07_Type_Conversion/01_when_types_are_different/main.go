// Go is very strict about explicit typing. There is no automatic type promotion or conversion.

package main

import "fmt"

func main() {
	a := 70   //int
	b := 44.2 // float64

	sum := a + b //int + float are not allowed

	fmt.Println("Sum of a and b is", sum)
}

// The above code is perfectly legal in C language, but in Go this program won’t compile. a is of type int and b is float64. We are trying to add 2 numbers of different types which is not allowed.

// When you run the program, you will get the following compilation error

// .\main.go:11:9: invalid operation: a + b (mismatched types int and float64)
