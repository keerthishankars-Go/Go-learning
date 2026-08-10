package main

// A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type.

// The syntax of a method declaration is provided below.

// func (t Type) methodName(parameter list) {
// }
// The above snippet creates a method named methodName with receiver type Type. t is called as the receiver and it can be accessed within the method.

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
 displaySalary() method has Employee as the receiver type
*/

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)

}

func main() {
	emp1 := Employee{
		name:     "Keerthishankar S",
		salary:   50000,
		currency: "$",
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type
}

// In line no. 25 of the program above, we have created a method displaySalary on Employee struct type. The displaySalary() method has access to the receiver e inside it. In line no. 26, we are using the receiver e and printing the name, currency and salary of the employee.

// In line no. 36 we have called the method using syntax emp1.displaySalary().
