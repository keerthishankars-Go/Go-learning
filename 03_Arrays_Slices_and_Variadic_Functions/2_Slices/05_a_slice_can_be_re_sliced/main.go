package main 

// A slice can be re-sliced upto its capacity. Anything beyond that will cause the program to throw a run time error.

import (
	"fmt"
)

func main () {
	fruitarray := [...]string{"orange", "watermelon", "grape", "mango", "butterfruit", "papaya"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("The length of fruitslice is %d and capacity is %d", len(fruitslice), cap(fruitslice))
	fruitslice = fruitslice[:cap(fruitslice)] //re-slicing fruitslice till its capacity
	fmt.Printf("\nThe length of the fruitslice after re-slicing is %d and capacity is %d", len(fruitslice), cap(fruitslice))
}