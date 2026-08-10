package main 

import (
	"fmt"
)

func subtractOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func main () {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtractOne(nos) //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside
}

// First: What a slice REALLY is (mentally)

// A slice is NOT the data itself.

// Internally, a slice looks like this (conceptually):

// type slice struct {
//     ptr *T   // points to underlying array
//     len int
//     cap int
// }


// So a slice is just:

// a pointer to an array

// plus length

// plus capacity

// 👉 The actual elements live in an array somewhere else

//===================================================================//

// Slices can be thought of as being represented internally by a structure type. This is how it looks,

// type slice struct {
//     Length        int
//     Capacity      int
//     ZerothElement *byte
// }
// A slice contains the length, capacity and a pointer to the zeroth element of the array. When a slice is passed to a function, even though it’s passed by value, the pointer variable will refer to the same underlying array. Hence when a slice is passed to a function as parameter, changes made inside the function are visible outside the function too.

// The function call in line number 17 of the above program decrements each element of the slice by 2. When the slice is printed after the function call, these changes are visible. If you can recall, this is different from an array where the changes made to an array inside a function are not visible outside the function.