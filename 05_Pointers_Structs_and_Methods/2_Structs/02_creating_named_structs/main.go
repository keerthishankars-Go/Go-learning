package main

// Let’s declare a named struct Employee using the following simple program.

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	// creating struct specifying field names
	emp1 := Employee{
		firstName: "Keerthi",
		lastName:  "S",
		salary:    25,
		age:       27,
	}

	// creating struct without specifying field names

	emp2 := Employee{"Thomas", "paul", 23, 32}

	fmt.Println("Employee 1 is", emp1)
	fmt.Println("Employee 2 is", emp2)
}

// In line no.9 of the above program, we create a named struct type Employee. In line no.16 of the above program, the emp1 struct is defined by specifying the value for each field name. The order of the fields need not necessarily be the same as that of the order of the field names while declaring the struct type. In this case. we have changed the position of lastName and moved it to the end. This will work without any problems.

// In line 26. of the above program, emp2 is defined by omitting the field names. In this case, it is necessary to maintain the order of the fields to be the same as specified in the struct declaration. Please refrain from using this syntax since it makes it difficult to figure out which value is for which field. We specified this format here just to understand that this is also a valid syntax
