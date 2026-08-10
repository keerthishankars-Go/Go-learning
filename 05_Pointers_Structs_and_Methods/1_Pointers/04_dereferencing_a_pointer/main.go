package main

// Dereferencing a pointer means accessing the value of the variable to which the pointer points. *a is the syntax to dereference a.

import (
	"fmt"
)

func main () {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
}


// In line no 10 of the above program, we deference a and print the value of it. As expected it prints the value of b. The output of the program is

// address of b is 0x1040a124
// value of b is 255