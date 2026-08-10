package main

// It is possible to create structs with fields that contain only a type without the field name. These kinds of fields are called anonymous fields.

// The snippet below creates a struct Person which has two anonymous fields string and int

import (
	"fmt"
)

// type Person struct {
// 	string
// 	int
// }

// Even though anonymous fields do not have an explicit name, by default the name of an anonymous field is the name of its type. For example in the case of the Person struct above, although the fields are anonymous, by default they take the name of the type of the fields. So Person struct has 2 fields with name string and int.

type Personal struct {
	string
	int
}

func main() {
	p1 := Personal{
		string: "Keerthi",
		int:    27,
	}

	fmt.Println(p1.string)
	fmt.Println(p1.int)
}

// In line no. 25 and 26 of the above program, we access the anonymous fields of the Person struct using their types as field name which is string and int respectively.
