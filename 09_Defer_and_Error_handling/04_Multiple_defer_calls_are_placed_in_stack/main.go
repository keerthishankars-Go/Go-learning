package main  

import "fmt"

func main() {
	str := "Gopher"
	fmt.Printf("Original string: %s\n", string(str))
	fmt.Printf("Reversed string: ")

	for _, v := range str {
		defer fmt.Printf("%c", v)
	}
}

// When a function has multiple defer calls, they are pushed to a stack and executed in Last In First Out (LIFO) order.

// We will write a small program which prints a string in reverse using a stack of defers.

// In the program above, the for range loop in line no. 11, iterates the string and calls defer fmt.Printf("%c", v) in line no. 12. These deferred calls will be added to a stack.

//  The stack is a last in first out datastructure. The defer call that is pushed to the stack last will be popped out and executed first. In this case defer fmt.Printf("%c", 'n') will be executed first and hence the string will be printed in reverse order.

// This program will print

// Original String: Gopher
// Reversed String: rehpoG

//