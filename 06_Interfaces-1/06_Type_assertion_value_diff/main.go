package main

// To solve the above problem, we can use the syntax

// v, ok := i.(T)
// If the concrete type of i is T, then v will have the underlying value of i and ok will be true.

// If the concrete type of i is not T, then ok will be false and v will have the zero value of type T and the program will not panic.

import "fmt"

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func main() {
	var s interface{} = 56
	assert(s)
	var j interface{} = "keerthi"
	assert(j)
}

// When keerthi is passed to the assert function, ok will be false since the concrete type of i is not int and v will have the value 0 which is the zero value of int. This program will print,

// 56 true
// 0 false
