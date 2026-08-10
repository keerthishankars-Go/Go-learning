package main

import (
	"fmt"
)

func main() {
	byteSlice := []byte{0x4B, 0x65, 0x65, 0x72, 0x74, 0x68, 0x69}
	// byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
}

// byteSlice in line no. 9 of the program above contains the UTF-8 Encoded hex bytes of the string Café. The program prints

// What if we have the decimal equivalent of hex values. Will the above program work? Let’s check it out.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	byteSlice := []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
// 	str := string(byteSlice)
// 	fmt.Println(str)
// }

// Decimal values also work and the above program will also print Café.