package main

// len(s) is used to find the number of bytes in the string and it doesn’t return the string length. As we already discussed, some Unicode characters have code points that occupy more than 1 byte. Using len to find out the length of those strings will return the incorrect string length.

// The RuneCountInString(s string) (n int) function of the utf8 package can be used to find the length of the string. This method takes a string as an argument and returns the number of runes in it.

import (
	"fmt"
	"unicode/utf8"
)

func main () {
	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))

	fmt.Printf("\n\n")
	word2 := "Keerthi"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2))

}

// The above output confirms that len(s) and RuneCountInString(s) return different values.