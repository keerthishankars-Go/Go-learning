package main

// It is possible to include multiple expressions in a case by separating them with comma.

import (
	"fmt"
)

func main() {
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Printf("%s is an vowel", letter)
	default:
		fmt.Printf("%s is not an vowel", letter)

	}
}

// The above program finds whether letter is a vowel or not. The code case "a", "e", "i", "o", "u": in line no. 12 matches any of the vowels. Since i is a vowel, this program prints:

// i is a vowel
