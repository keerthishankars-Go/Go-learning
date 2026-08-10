package main

import "fmt"

func main() {
	age := 27 //age is int

	age = "naveen" // error since we are trying to assign a string to a variable of type int

	fmt.Println("My age is", age)
}

//Since Go is strongly typed, variables declared as belonging to one type cannot be assigned a value of another type. The following program will print an error:  .\main.go:8:8: cannot use "naveen" (untyped string constant) as int value in assignment

//because age is declared as type int and we are trying to assign a string value to it.
