package main  

// A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are just references to existing arrays.

//A slice with elements of type T is represented by []T

import (
	"fmt"
)

func main () {
	a := [5]int{23, 33, 46, 67, 21}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)
}

// The syntax a[start:end] creates a slice from array a starting from index start to index end - 1. So in line no. 13 of the above program a[1:4] creates a slice representation of the array a starting from indexes 1 through 3. Hence the slice b has values [77 78 79].