// Short hand syntax can only be used when at least one of the variables on the left side of := is newly declared.
package main

import "fmt"

func main() {

	a, b := 10, 5
	fmt.Println("a is ", a, "b is ", b)

	b, c := 15, 23
	fmt.Println("new b is ", b, "c is ", c)

	b, c := 12, 55
	fmt.Println("Changed b is ", b, "c is ", c)
}

//It will print error .\main.go:14:7: no new variables on left side of := This is because both the variables a and b have already been declared and there are no new variables in the left side of :=
