package main

// Go also provides a handy function new to create pointers. The new function takes a type as an argument and returns a pointer to a newly allocated zero value of the type passed as argument.

import (
	"fmt"
)

func main() {
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %vz\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
}

//  we use the new function to create a pointer of type int. This function will return a pointer to a newly allocated zero value of the type int. The zero value of type int is 0. Hence size will be of type *int and will point to 0 i.e *size will be 0.

// The above program will print

// Size value is 0, type is *int, address is 0x414020
// New size value is 85
