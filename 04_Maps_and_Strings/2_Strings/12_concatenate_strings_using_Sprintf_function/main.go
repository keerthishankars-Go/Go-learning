package main  

import (
	"fmt"
)

// The second way to concatenate strings is using the Sprintf function of the fmt package.

// The Sprintf function formats a string according to the input format specifier and returns the resulting string. Let’s rewrite the above program using Sprintf function.

func main () {
	string1 := "Go"
	string2 := "is awesome"
	result := fmt.Sprintf("%s %s ", string1, string2)
	fmt.Println(result)
}

// In the program above, %s %s is the format specifier input for Sprintf. This format specifier takes two strings as input and has a space in between. This will concatenate the two strings with a space in the middle. The resulting string is stored in result. This program also prints,

// Go is awesome