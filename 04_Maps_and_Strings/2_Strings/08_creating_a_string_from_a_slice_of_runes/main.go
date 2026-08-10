package main

import (
	"fmt"
)

func main() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)
}

// In the above program runeSlice contains the Unicode code points of the string Señor in hexadecimal. The program outputs:

// Señor