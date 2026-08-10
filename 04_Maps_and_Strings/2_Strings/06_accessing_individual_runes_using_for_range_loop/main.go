package main  

import (
	"fmt"
)

func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func main () {
	name := "Señor"
	charsAndBytePosition(name)
}

// From the above output, it’s clear that ñ occupies 2 bytes since the next character o starts at byte 4 instead of byte 3 