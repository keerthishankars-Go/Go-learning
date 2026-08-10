A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are just references to existing arrays.

A slice with elements of type T is represented by []T


Let’s look at another way to create a slice.

package main

import (
	"fmt"
)

func main() {
	c := []int{6, 7, 8} //creates an array and returns a slice reference
	fmt.Println(c)
}


In the above program, line no. 8 creates an array with 3 integers and returns a slice reference which is stored in the variable c.