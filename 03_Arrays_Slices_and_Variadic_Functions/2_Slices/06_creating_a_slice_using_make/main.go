package main  

// func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity. The capacity parameter is optional and defaults to the length. The make function creates an array and returns a slice reference to it..

import (
	"fmt"
)

func main () {
	i := make([]int, 5)
	fmt.Println(i)
}

// The values are zeroed by default when a slice is created using make. The above program will print [0 0 0 0 0].