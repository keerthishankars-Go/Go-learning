//Multiple variables can be declared using a single statement.

// " var name1, name2 type = initialvalue1, initialvalue2 " is the syntax for multiple variable declaration.

package main

import "fmt"

func main() {

	var price, quantity int = 2000, 50 //declaring multiple variables..

	fmt.Println("The price is", price, "for quantities", quantity)
}

//The type can be removed if the variables have an initial value. Since the above program has initial values for variables, the int type can be removed.

func withoutInitialValue () {

	var price, quantity int 

	fmt.Println("Price is", price , "quantity is", quantity)
}