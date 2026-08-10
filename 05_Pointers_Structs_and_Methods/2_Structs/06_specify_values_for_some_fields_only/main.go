package main

// It is also possible to specify values for some fields and ignore the rest. In this case, the ignored fields are assigned zero values.

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
	emp5 := Employee{
		firstName: "Keerthi",
		lastName:  "S",
	}

	fmt.Println("The first name is", emp5.firstName)
	fmt.Println("The last name is", emp5.lastName)
	fmt.Println("Age: ", emp5.age)
	fmt.Println("salary: \n", emp5.salary)
	emp5.salary = 56
	fmt.Printf("The new salary is: $%d", emp5.salary)

}

// In the above program in line. no 18 and 19, firstName and lastName are initialized whereas age and salary are not. Hence age and salary are assigned their zero values.
