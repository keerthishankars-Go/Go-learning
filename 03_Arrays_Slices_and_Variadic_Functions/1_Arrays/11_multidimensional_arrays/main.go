package main

import (
	"fmt"
)

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
func main() {

	a := [3][2]string{
		{"lion", "cat"},
		{"tiger", "rabbit"},
		{"pigeon", "peacock"},
	}
	printarray(a)
	
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "redmi"
	b[2][0] = "dell"
	b[2][1] = "google"
	fmt.Printf("\n")
	printarray(b)
}

// In the above program in line no. 17, a two dimensional string array a has been declared and defined using short hand syntax. The comma at the end of line no. 20 is necessary. This is because of the fact that the lexer automatically inserts semicolons according to simple rules.

// The printarray function in line no. 7 uses two for range loops to print the contents of 2d arrays.

// Although arrays seem to be flexible enough, they come with the restriction that they are of fixed length. It is not possible to increase the length of an array. This is where slices come into the picture. In fact in Go, slices are more commonly used than arrays.
