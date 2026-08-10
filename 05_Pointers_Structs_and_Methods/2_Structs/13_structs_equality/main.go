package main  

// Structs are value types and are comparable if each of their fields are comparable. Two struct variables are considered equal if their corresponding fields are equal.

import (
	"fmt"
)

type name struct {
	firstName string
	lastName  string

}

func main () {
	name1 := name{
		firstName : "Keerthi",
		lastName : "shankar",
	}

	name2 := name{
		firstName : "Keerthi",
		lastName : "shankar",
	}

	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name {
		firstName : "Keerthi",
		lastName : "Shankar",

	}

	name4 := name {
		firstName : "Keerthi",
		// lastName : "S",
	}

	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}