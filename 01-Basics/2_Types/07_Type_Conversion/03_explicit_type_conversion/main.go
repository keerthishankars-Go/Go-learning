// Rule: If types are different → you MUST convert explicitly
package main

import "fmt"

func main() {

	i := 10

	var j float64 = float64(i)

	// var j float64 = i

	// .\main.go:10:18: cannot use i (variable of type int) as float64 value in variable declaration

	fmt.Println("j =", j)

}

//Explicit type conversion is required to assign a variable of one type to another.
