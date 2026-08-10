package main 

// for initialisation; condition; post {
// }

// Let’s write a program that uses for loop to print the numbers 1 to 10.

import (
	"fmt"
)

func main () {

	for i := 1 ; i <= 10 ; i++ {
		fmt.Printf("%d  ", i) //give space after %d to have a spaced numbers in return
	}
}


//In the above program, i is initialised to 1. The conditional statement will check whether i <= 10. If the condition is true, the value of i is printed, else the loop is terminated. The post statement increments i by 1 at the end of each iteration. Once i is greater than 10, the loop terminates.

// The above program will print 1 2 3 4 5 6 7 8 9 10

// The variables declared in a for loop are only available within the scope of the loop. Hence i cannot be accessed outside the body of the for loop.