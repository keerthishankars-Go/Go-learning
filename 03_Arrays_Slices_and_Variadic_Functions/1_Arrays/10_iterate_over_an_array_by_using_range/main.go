package main

// Go provides a better and concise way to iterate over an array by using the range form of the for loop. range returns both the index and the value at that index. We will also find the sum of all elements of the array.

// Syntax of for range :

// for index, value := range collection {
// }

import "fmt"

func main() {
	a := [...]float64{56.23, 44.87, 87.21, 21}
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d th element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a is ", sum)
}

//the above program uses the range form of the for loop. It will return both the index and the value at that index

// In case you want only the value and want to ignore the index, you can do this by replacing the index with the _ blank identifier.

// for _, v := range a { //ignores index
// }

// The above for loop ignores the index. Similarly, the value can also be ignored.
