// Short hand declaration requires initial values for all variables on the left-hand side of the assignment.

package main

import "fmt"

func main () {

	name, age := "keerthi" //error

	//it will print an error assignment mismatch: 2 variables but 1 value. This is because age has not been assigned a value.

	fmt.Println("My name is", name)
	fmt.Println("age is", age)
}
