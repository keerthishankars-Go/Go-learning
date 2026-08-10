// Variable is the name given to a memory location to store a value of a specific type. There are various syntaxes to declare variables in Go. 
// var name type is the syntax to declare a single variable.
package main

import "fmt"

func main() {

	var age int // variable declaration
	fmt.Println("my initial age is:", age)

}

//The statement var age int declares a variable named age of type int. We have not assigned any value to this variable. If a variable is not assigned any value, Go automatically initializes it with the zero value of the variable’s type. In this case, age is assigned the value 0, which is the zero value of int. 
