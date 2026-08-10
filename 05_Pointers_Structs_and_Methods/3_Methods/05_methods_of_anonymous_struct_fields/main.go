package main

// Methods belonging to anonymous fields of a struct can be called as if they belong to the structure where the anonymous field is defined.

import (
	"fmt"
)

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address : %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

func main() {
	p := person{
		firstName: "Keerthi",
		lastName:  "shankar",
		address: address{
			city:  "Hassan",
			state: "Karnataka",
		},
	}
	p.fullAddress() //accessing fullAddress method of address struct
}

// In line no. 33 of the program above, we call the fullAddress() method of the address struct using p.fullAddress(). The explicit direction p.address.fullAddress() is not needed.
