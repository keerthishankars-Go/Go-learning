package main 

// Let’s modify the above program a little bit to print the characters of the string.

import (
	"fmt"
)

func printBytes (s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Chars: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func main () {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
}

// %c format specifier is used to print the characters of the string in the printChars method. 

// Although the above program looks like a legitimate way to access the individual characters of a string, this has a serious bug.