package main 

import "fmt"

func displayValue(a int) {
	fmt.Println("Value of a in differed function", a)
}

func main() {
	a := 5

	defer displayValue(a)

	a = 10

	fmt.Println("Value of a before deferred function call", a)
}

// The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual function call is done.

// In the program above a initially has a value of 5 in line no. 11. When the defer statement is executed in line no. 12, the value of a is 5 and hence this will be the argument to the displayValue function which is deferred. We change the value of a to 10 in line no. 13. The next line prints the value of a. This program outputs,

// value of a before deferred function call 10
// value of a in deferred function 5

// From the above output it can be understood that although the value of a changes to 10 after the defer statement is executed, the actual deferred function call displayValue(a) still prints 5.