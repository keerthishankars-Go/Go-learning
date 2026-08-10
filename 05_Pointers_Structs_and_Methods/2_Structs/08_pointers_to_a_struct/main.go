package main

// It is also possible to create pointers to a struct.

import(
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main () {
	emp8 := &Employee{
		firstName : "Keerthi",
		lastName : "S",
		age : 27,
		salary : 28,
	}

	fmt.Println("The first name is", (*emp8).firstName)
	fmt.Println("The last name is ", (*emp8).lastName)
	
}

// emp8 in the above program is a pointer to the Employee struct. (*emp8).firstName is the syntax to access the firstName field of the emp8 struct.

// The Go language gives us the option to use emp8.firstName instead of the explicit dereference (*emp8).firstName to access the firstName field.

// func main() {
// 	emp8 := &Employee{
// 		firstName: "Sam",
// 		lastName:  "Anderson",
// 		age:       55,
// 		salary:    6000,
// 	}
// 	fmt.Println("First Name:", emp8.firstName)
// 	fmt.Println("Age:", emp8.age)
// }
