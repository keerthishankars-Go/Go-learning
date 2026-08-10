package main 

// A slice does not own any data of its own. It is just a representation of the underlying array. Any modifications done to the slice will be reflected in the underlying array and vice versa.

import (
	"fmt"
)

func main () {
	darr := [...]int{56, 59, 43, 76, 141, 32, 87, 83, 12}
	dslice := darr[2:5]
	fmt.Println("array before", darr)

	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}

// In the above program, we create dslice from indexes 2, 3, 4 of the array. The for loop increments the value in these indexes by one. When we print the array after the for loop, we can see that the changes to the slice are reflected in the array.