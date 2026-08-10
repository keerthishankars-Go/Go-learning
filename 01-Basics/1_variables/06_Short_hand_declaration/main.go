// Another concise way to declare variables is short hand declaration..
// It uses := operator
package main

import "fmt"

func main() {

	// "name := initialvalue" is the short hand syntax to declare a variable.
	count := 10 //declare a variable count initialized to 10

	fmt.Println("Count =", count)

}

//Go will automatically infer that count is of type int since it has been initialized with the integer value 10.
