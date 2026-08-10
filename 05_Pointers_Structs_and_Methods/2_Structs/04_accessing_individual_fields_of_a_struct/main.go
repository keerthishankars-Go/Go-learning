package main

// The dot . operator is used to access the individual fields of a struct.

import (
	"fmt"


)

type Employee struct {
	firstName string
	lastName string
	age		int
	salary	int
}

func main () {
	emp6 := Employee{
		firstName: "Keerthi",
		lastName: "S",
		age: 27,
		salary: 25,
	}
	fmt.Println("The first name is", emp6.firstName)
	fmt.Println("last name is", emp6.lastName)
	fmt.Println("The age is", emp6.age)
	fmt.Printf("The salary: $%d\n", emp6.salary)
	emp6.salary = 33
	fmt.Printf("The revised salary: $%d\n", emp6.salary)
	fmt.Printf("age of %s is %d", emp6.firstName, emp6.age)

}

// emp6.firstName in the above program is used to access the firstName field of the emp6 struct. In line no. 25 we modify the salary of the employee.