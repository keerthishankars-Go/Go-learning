package main

// When a struct is defined and if it is not explicitly initialized with any value, the fields of the struct are assigned their zero values by default.

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main () {
	var emp4 Employee // zero valued struct
	fmt.Println("The first name is", emp4.firstName)
	fmt.Println("The last name is", emp4.lastName)
	fmt.Println("Age is: ", emp4.age)
	fmt.Println("Salary is: ", emp4.salary)
}

// The above program defines emp4 but it is not initialized with any value. Hence firstName and lastName are assigned the zero values of string which is an empty string "" and age, salary are assigned the zero values of int which is 0. This program prints,

// First Name: 
// Last Name: 
// Age: 0
// Salary: 0