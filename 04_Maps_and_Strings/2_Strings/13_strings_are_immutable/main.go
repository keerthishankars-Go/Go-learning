package main

// Strings are immutable in Go. Once a string is created it’s not possible to change it.

import (
	"fmt"
)

// func mutate(s string)string {
// 	s[0] = 'a' //any valid unicode character within single quote is a rune
// 	return s

// }

// func main () {
// 	h := "hello"
// 	fmt.Println(mutate(h))
// }

// we try to change the first character of the string to 'a'. Any valid Unicode character within a single quote is a rune. We try to assign the rune a to the zeroth position of the slice. This is not allowed since the string is immutable and hence the program fails to compile with error ./prog.go:8:2: cannot assign to s[0] (neither addressable nor a map index expression)

// so...

// To workaround this string immutability, strings are converted to a slice of runes. Then that slice is mutated with whatever changes are needed and converted back to a new string.

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	h := "hello"
	fmt.Println("Before mutation: ", h)
	fmt.Println(mutate([]rune(h)))
}

// In line no.26 of the above program, the mutate function accepts a rune slice as an argument. It then changes the first element of the slice to 'a', converts the rune back to string and returns it. This method is called from line no. 34 of the program. h is converted to a slice of runes and passed to mutate in line no. 34. This program outputs aello
