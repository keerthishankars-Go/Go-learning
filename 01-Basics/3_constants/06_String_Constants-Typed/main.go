// const name string = "Hello World" : way to create typed constant
// name in the above code is a constant of type string.

// Go is a strongly typed language. Mixing types during the assignment is not allowed.

package main

import "fmt"

func main() {
	var defaultName = "Sam" // allowed
	type myString string
	var customName myString = "Sam" //allowed

	//customName = defaultName  //(make this uncomment)
	customName = myString(defaultName) //(make this comment)
	fmt.Println(customName)

}

// In the above code, we first create a variable defaultName and assign it to the constant Sam. The default type of the constant Sam is a string, so after the assignment defaultName is of type string.

// In the next line, we create a new type myString which is an alias of string.

// Then we create a variable customName of type myString and assign the constant Sam to it. Since the constant Sam is untyped, it can be assigned to any string variable. Hence this assignment is allowed and customName gets the type myString.

// Now we have a variable defaultName of type string and another variable customName of type myString. Even though we know that myString is an alias of string, Go’s strong typing policy disallows variables of one type to be assigned to another. Hence the assignment customName = defaultName is not allowed and the compiler throws the error .\main.go:15:15: cannot use defaultName (variable of type string) as myString value in assignment

//To make the above program work, defaultName must be converted to type myString.
