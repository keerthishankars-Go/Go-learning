package main

//A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements. It is similar to switch case. The only difference being the cases specify types and not values as in normal switch.

// The syntax for type switch is similar to Type assertion. In the syntax i.(T) for Type assertion, the type T should be replaced by the keyword type for type switch

import "fmt"

// func findType(i interface{}) {
// 	switch i.(type) {
// 	case string:
// 		fmt.Printf("I am a string and my value is %s\n", i.(string))

// 	case int:
// 		fmt.Printf("I am an int and my value is %d\n", i.(int))

// 	default:
// 		fmt.Printf("unknown type")
// 	}

// }

func findType(i interface{}) {

	switch value := i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", value)

	case int:
		fmt.Printf("I am an int and may value is %d\n", value)

	default:
		fmt.Printf("unknown type")
	}
}

func main() {
	findType("Keerthi")
	findType(56)
	findType(98.06)
}

// Now look at this line:

// switch value := i.(type)

// This does two things:

// Checks the type
// Creates a variable with the actual type