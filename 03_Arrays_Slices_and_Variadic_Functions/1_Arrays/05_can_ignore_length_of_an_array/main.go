package main  

// we can even ignore the length of the array in the declaration and replace it with ... and let the compiler find the length for us. This is done in the following program.

import (
	"fmt"
)

func main () {
	a := [...]int{23, 43, 88} // ... makes the compiler determine the length
	fmt.Println(a)
}

// The size of the array is a part of the type. Hence [5]int and [25]int are distinct types. Because of this, arrays cannot be resized. Don’t worry about this restriction since slices exist to overcome this.