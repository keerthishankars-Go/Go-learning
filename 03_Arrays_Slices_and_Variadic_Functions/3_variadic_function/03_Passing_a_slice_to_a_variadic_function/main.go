package main

import (
	"fmt"
)

// Let’s pass a slice to a variadic function and find out what happens from the below example.

func find(num int, nums ...int) {
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
	nums := []int{89, 90, 89, 95}
	//find(89, nums) // This wont compile
	find(89, nums...)

}

// In line no. 25, we are passing a slice to a function that expects a variable number of arguments.

// This will not work. The above program will fail with compilation error ./prog.go:23:10: cannot use nums (type []int) as type int in argument to find

// Why does this don’t work? Well, it’s pretty straight forward. The signature of the find function is provided below,

// func find(num int, nums ...int)
// According to the definition of a variadic function, nums ...int means that it will accept a variable number of arguments of type int.

// In line no. 25 of the program above, nums which is []int slice is passed to the find function which is expecting a variadic int argument. As we already discussed, these variadic arguments will be converted to a slice of type int since find expects variadic int arguments. In this case, nums is already a []int slice and the compiler tries to create a new []int i.e the compiler tries to do

// find(89, []int{nums})

// which will fail since nums is a []int and not a int.
