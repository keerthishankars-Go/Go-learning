// Strings are a collection of bytes in Go, assume a string to be a collection of characters.

package main

import "fmt"

func main() {

	first := "keerthi"
	last := "shankar"

	name := first + " " + last

	fmt.Println("My name is", name)
}

// In the above program, first is assigned the string keerthi, last is assigned the string shankar. Strings can be concatenated using the + operator.

// Name is assigned the value of first concatenated by a space followed by last.
