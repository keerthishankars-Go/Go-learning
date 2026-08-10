package main

// A for loop which has another for loop inside it is called a nested for loop.

// Let’s understand nested for loops by writing a program that prints the sequence below.

import (
	"fmt"
)

func main() {
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}

// The program above uses nested for loops to print the sequence. The variable n in line no. 8 stores the number of lines in the sequence. In our case it’s 5. The outer for loop iterates i from 0 to 4 and the inner for loop iterates j from 0 to the current value of i. The inner loop prints * for each iteration and the outer loop prints a new line at the end of each iteration.
