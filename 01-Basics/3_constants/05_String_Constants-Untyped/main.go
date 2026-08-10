// Any value enclosed between double quotes is a string constant in Go. For example, strings like "Hello World", "Sam" are all constants in Go.
// String constant belongs to untyped..
// Ex: A string constant like “Hello World” does not have any type..
// const hello = "Hello World"

package main

import "fmt"

func main() {
	const n = "sam"
	var name = n

	fmt.Printf("type %T value %v", name, name)
}

// Go is a strongly typed language. All variables require an explicit type. How does the above program which assigns a variable name to an untyped constant n work?

// answer: untyped constants have a default type associated with them and they supply it if and only if a line of code demands it.

// In the statement var name = n in line no. 12, name needs a type and it gets it from the default type of the string constant n which is a string.
