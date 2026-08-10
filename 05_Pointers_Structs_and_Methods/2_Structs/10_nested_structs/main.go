package main

// It is possible for a struct to have a field which in turn is a struct type. These kinds of structs are called nested structs.

import (
	"fmt"
)

type Address struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	address Address
}

func main () {	
	p := Person {
		name : "Keerthi",
		age : 27,
		address: Address{
			city : "Hassan",
			state: "Karnataka",
		},

	}
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ",p.age)
	fmt.Println("City: ", p.address.city)
	fmt.Println("State: ", p.address.state)

}

// The Person struct in the above program has a field address which in turn is a struct.