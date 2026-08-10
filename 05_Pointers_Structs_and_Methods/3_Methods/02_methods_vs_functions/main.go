package main

// The above program can be rewritten using only functions and without methods.

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
 displaySalary() method converted to function with Employee as parameter
*/

func displaySalary(e Employee) {
	fmt.Printf("The salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee{
		name:     "Keerthishankar S",
		salary:   50000,
		currency: "$",
	}
	displaySalary(emp1)
}

// In the program above, the displaySalary method is converted to a function and the Employee struct is passed as a parameter to it. This program also produces the exact same output..

// Methods with the same name can be defined on different types whereas functions with the same names are not allowed.