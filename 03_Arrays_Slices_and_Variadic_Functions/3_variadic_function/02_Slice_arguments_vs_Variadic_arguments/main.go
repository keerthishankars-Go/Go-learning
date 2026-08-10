package main

// we learned that the variadic arguments to a function are in fact converted a slice. Then why do we even need variadic functions when we can achieve the same functionality using slices? I have rewritten the program above using slices below.

import (
	"fmt"
)

func find(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func main() {
	find(89, []int{89, 90, 95})
	find(45, []int{56, 67, 45, 90, 109})
	find(78, []int{38, 56, 98})
	find(87, []int{})
}

// The following of the advantages of using variadic arguments instead of slices.

// There is no need to create a slice during each function call. If you look at the program above, we have created new slices during each function call in line nos. 24, 25, 26 and 27. This additional slice creation can be avoided when using variadic functions
// In line no.27 of the program above, we are creating an empty slice just to satisfy the signature of the find function. This is totally not needed in the case of variadic functions. This line can just be find(87) when variadic function is used.
// I personally feel that the program with variadic functions is more readable than the once with slices :)
