package main

import "fmt"

func main() {

	a := true
	b := false

	fmt.Println("a:", a, " b:", b)

	c := a && b
	fmt.Println("c:", c)

	d := a || b
	fmt.Println("d:", d)
}

// && is a boolean operator which returns true when both of the operands are true.

// The || operator returns true when either a or b is true.
