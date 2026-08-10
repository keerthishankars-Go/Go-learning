package main

//Type assertion is used to extract the underlying value of the interface.

// i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T.

import "fmt"

func assert(i interface{}) {
	s := i.(int)  //get the underlying int value from i
	fmt.Println(s)
}

func main() {
	var s interface{} = 46
	assert(s)
}

// The concrete type of s in line no. 12 is int. We use the syntax i.(int) in line no. 8 to fetch the underlying int value of i. This program prints 56.

//============================================================================//

// What will happen if the concrete type in the above program is not int? Well, let’s find out.

// package main

// import (
// 	"fmt"
// )

// func assert(i interface{}) {
// 	s := i.(int) 
// 	fmt.Println(s)
// }
// func main() {
// 	var s interface{} = "Steven Paul"
// 	assert(s)
// }


// In the program above we pass s of concrete type string to the assert function which tries to extract an int value from it. This program will panic with the message panic: interface conversion: interface {} is string, not int.

