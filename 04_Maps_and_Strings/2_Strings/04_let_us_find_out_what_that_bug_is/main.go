package main

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
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n\n")

	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
}

// The output of the above program is

// String: Hello World
// Characters: H e l l o   W o r l d 
// Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64 

// String: Señor
// Characters: S e Ã ± o r 
// Bytes: 53 65 c3 b1 6f 72 

// In line no. 30 of the program above, we are trying to print the characters of Señor and it outputs S e Ã ± o r which is wrong. 
// Why does this program break for Señor when it works perfectly fine for Hello World . 
// The reason is that the Unicode code point of ñ is U+00F1 and its UTF-8 encoding occupies 2 bytes c3 and b1. We are trying to print the characters assuming that each code point will be one byte long which is wrong. 
// In UTF-8 encoding a code point can occupy more than 1 byte. 
// So how do we solve this? This is where rune saves us.