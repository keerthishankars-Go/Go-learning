package main  

import (
	"fmt"
)

func compareStrings(str1 string, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}

func main() {
	string1 := "Go"
	string2 := "Go"
	compareStrings(string1, string2)

	string3 := "hello"
	string4 := "world"
	compareStrings(string3, string4)

}

// In the compareStrings function above, line no. 8 compares whether the two strings str1 and str2 are equal using the == operator. If they are equal, it prints a corresponding message and the function returns.

// The above program prints,

// Go and Go are equal
// hello and world are not equal