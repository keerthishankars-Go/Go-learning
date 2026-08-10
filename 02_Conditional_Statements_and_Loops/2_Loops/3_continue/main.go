package main  

// The continue statement is used to skip the current iteration of the for loop. All code present in a for loop after the continue statement will not be executed for the current iteration. The loop will move on to the next iteration.

// Let’s write a program to print all odd numbers from 1 to 10 using continue.

import (
	"fmt"
)

func main () {

	for i := 1; i <= 10 ; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

// In the above program line no. 9 checks if the remainder of dividing i by 2 is 0. If it is zero, then the number is even and continue statement is executed and the control moves to the next iteration of the loop. Hence the print statement after the continue will not be called and the loop proceeds to the next iteration. The output of the above program is 1 3 5 7 9