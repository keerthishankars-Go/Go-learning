package main

// If a struct type starts with a capital letter, then it is an exported type and it can be accessed from other packages. Similarly, if the fields of a struct start with caps, they can be accessed from other packages.

import (
	"fmt"
	"go-learning/computer"
)

func main() {
	spec := computer.Spec{
		Maker: "Apple",
		Price: 55000,
		//model: "Mac Mini",
	}
	fmt.Println("Maker: ", spec.Maker)
	fmt.Println("Price: ", spec.Price)
}

// In line no. 7 of the program above, we import the computer package. In line no. 12 and 13, we access the two exported fields Maker and Price of the struct Spec.

// If we try to access unexported struct field, compiler will complain like: in line number 14 : .\main.go:14:3: cannot refer to unexported field model in struct literal of type computer.Spec

// Since model field is unexported, it cannot be accessed from other packages.

