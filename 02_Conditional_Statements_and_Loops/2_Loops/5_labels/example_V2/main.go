package main

// What if we stop printing when i and j are equal.

// To do this we need to break from the outer for loop. Adding a break in the inner for loop when i and j are equal will only break from the inner for loop.

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}
}

//In the program above, I have added a break inside the inner for loop when i and j are equal in line no. 15. This will break only from the inner for loop and the outer loop will continue.
