package main //Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

import "fmt"

func sum(nums ...int) { //Here’s a function that will take an arbitrary number of ints as arguments.
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums { //Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.
		total += num

	}
	fmt.Println(total)
}

func main() {
	sum(1, 2) //Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4} //If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	sum(nums...)
}
