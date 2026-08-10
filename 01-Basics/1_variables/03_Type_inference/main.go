// If a variable has an initial value, Go will automatically be able to infer the type of that variable using that initial value. Hence if a variable has an initial value, the type in the variable declaration can be removed.
package main

import "fmt"

func main() {

	// If the variable is declared using the following syntax:
	// var name = initialvalue
	var age = 27 //type will be inferred by go automatically..

	fmt.Println("My age currently is", age)
}

// Go will automatically infer the type of that variable from the initial value.

//In this example, we can see that the type int of the variable age has been removed in line no. 6. Since the variable has an initial value 29, Go can infer that it is of type int.