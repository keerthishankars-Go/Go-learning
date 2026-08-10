package main

// It is perfectly legal for a function to return a pointer of a local variable. The Go compiler is intelligent enough and it will allocate this variable on the heap.

import (
	"fmt"
)

func hello() *int {
	i := 5
	return &i
}

func main() {
	d := hello()
	fmt.Println("value of d is", *d)
}

// In line no. 11 of the program above, we return the address of the local variable i from the function hello. 
// The behaviour of this code is undefined in programming languages such as C and C++ as the variable i goes out of scope once the function hello returns. 
// But in the case of Go, the compiler does an escape analysis and allocates i on the heap as the address escapes the local scope. Hence this program will work and it will print,

// Value of d 5
