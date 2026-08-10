package main  

// It is possible to declare and operate on multiple variables in a for loop.

import (
	"fmt"
)

func main () {
	for no, i := 10, 1; i <= 10 && no <= 19; i , no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)

	}
}

// In the above program no and i are declared and initialized to 10 and 1 respectively. They are incremented by 1 at the end of each iteration. The boolean operator && is used in the condition to ensure that i is less than or equal to 10 and also no is less than or equal to 19.