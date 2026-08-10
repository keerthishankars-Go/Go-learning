package main

// The for loop can be used to iterate over the elements of an array.

import "fmt"

func main() {
	a := [...]float64{24.54, 42.66, 33, 82.76}

	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array..
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}

}

// The above program uses a for loop to iterate over the elements of the array starting from index 0 to length of the array - 1. This program works and will print:
