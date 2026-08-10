package main  

// A pointer stores the memory address of another variable.

// In this example, So instead of storing 255, it stores where 255 lives in memory.

// Pointer type syntax: *T

// *T
// means:

// * → pointer

// T → type it points to

// Examples:

// *int → pointer to an int

// *string → pointer to a string

import (
	"fmt"
)

func main () {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a) 	
}

// The & operator is used to get the address of a variable. In line no. 28 of the above program we are assigning the address of b to a whose type is *int. Now a is said to point to b. When we print the value in a, the address of b will be printed. This program outputs

// Type of a is *int
// address of b is 0x1040a124
// You might get a different address for b since the location of b can be anywhere in memory.