package main  

// A rune is a builtin type in Go and it’s the alias of int32. Rune represents a Unicode code point in Go. It doesn’t matter how many bytes the code point occupies, it can be represented by a rune. Let’s modify the above program to print characters using a rune.

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Chars: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])

	}
}

func main () {
	name := "Keerthi shankar"
	fmt.Printf("String: %s\n",name)
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n\n")

	name = "Señor"
	fmt.Printf("String: %s", name)
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)

}

// the program above, the string is converted to a slice of runes. We then loop over it and display the characters.