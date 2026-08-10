package main

// Write a simple program to find whether an integer exists in an input list of integers.

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("Type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}

	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}

	fmt.Printf("\n")
}

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 76, 45, 88, 90)
	find(78, 38, 56, 98)
	find(89)
}

// In the above program, func find(num int, nums ...int) in line no.9, accepts variable number of arguments for the nums parameter. Inside the function find, the type of nums is []int i.e, an integer slice.

// The way variadic functions work is by converting the variable number of arguments to a slice of the variadic parameter’s type. For instance, in line no. 27 of the program above, the variable number of arguments to the find function are 89, 90, 95. The find function expects a variadic int argument. Hence these three arguments will be converted by the compiler to a slice of type int []int{89, 90, 95} and then it will be passed to the find function.

// In line no. 12, the for loop ranges over the nums slice and prints the index of num if it is present in the slice. If not, it prints that the number is not found.

// In line no. 30 of the above program, the find function call has only one argument. We have not passed any argument to the variadic nums ...int parameter. As discussed earlier, this is perfectly legal and in this case, nums will be a nil slice with length and capacity 0.
