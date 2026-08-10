package main

// The break statement is used to terminate the for loop abruptly before it finishes its normal execution and move the control to the line of code just after the for loop.

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		if i > 5 { // loop is terminated if i > 5
			break
		}
		fmt.Printf("%d  ", i)
	}
	fmt.Println("\nloop terminated")
}

//In the above program, the value of i is checked during each iteration. If i is greater than 5 then break executes and the loop is terminated. The print statement just after the for loop is then executed. The above program will output,

// 1 2 3 4 5 
// loop ended