package main 

// Since a string is a slice of bytes, it’s possible to access each byte of a string.

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i:=0 ; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func main () {
	name := "Hello world"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
}

// %s is the format specifier to print a string. In line no. 16, the input string is printed. In line no. 9 of the program above, len(s) returns the number of bytes in the string and we use a for loop to print those bytes in hexadecimal notation. %x is the format specifier for hexadecimal. The above program outputs

// String: Hello World

// Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64  
// These are the Unicode UT8-encoded values of Hello World. A basic understanding of Unicode and UTF-8 is needed to understand strings better.