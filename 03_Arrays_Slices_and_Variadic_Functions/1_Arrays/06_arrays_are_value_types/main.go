package main

// Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable. If changes are made to the new array, it will not be reflected in the original array.

import (
	"fmt"
)

func main() {
	a := [...]string{"India", "Germany", "China", "USA", "France"}
	b := a
	b[0] = "singapore"

	fmt.Println("a is", a)
	fmt.Println("b is", b)
}

// In the above program in line no. 11, a copy of a is assigned to b. In line no. 12, the first element of b is changed to Singapore. This will not reflect in the original array a. The program will print,
