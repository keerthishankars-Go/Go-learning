// Go has various value types including strings, integers, floats, booleans, etc.
package main

import "fmt"

func main() {
	fmt.Println("go" + "lang") // Strings, which can be added together with +.

	// Integers and floats.

	fmt.Println("1+1=", 1+1)

	fmt.Println("7.0/3.0=", 7.0/3.0)

	// Booleans, with boolean operators

	fmt.Println(true && false)

	fmt.Println(true || false)

	fmt.Println(!true)

}
