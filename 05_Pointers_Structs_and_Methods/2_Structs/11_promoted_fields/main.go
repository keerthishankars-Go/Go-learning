package main

// Fields that belong to an anonymous struct field are called promoted fields since they can be accessed as if they belong to the struct which holds the anonymous struct field.

// type Address struct {
// 	city string
// 	state string
// }
// type Person struct {
// 	name string
// 	age  int
// 	Address
// }

// In the above code snippet, the Person struct has an anonymous field Address which is a struct. Now the fields of the Address namely city and state are called promoted fields since they can be accessed as if they are directly declared in the Person struct itself.

import (
	"fmt"
)

type Address struct {
	city  string
	state string
}

type Person struct {
	name string
	age  int
	Address
}

func main() {
	p := Person{
		name: "Keerthi",
		age:  27,
		Address: Address{
			city:  "Hassan",
			state: "Karnataka",
		},
	}

	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("City: ", p.city)   // city is promoted field..
	fmt.Println("State: ", p.state) // state is promoted field..
}

// In line no. 44 and 45 of the program above, the promoted fields city and state are accessed as if they are declared in the struct p itself using the syntax p.city and p.state.
