// if statement has a condition and it executes a block of code if that condition evaluates to true. It executes an alternate else block if the condition evaluates to false.

// If statement syntax:

// if condition {
// }

package main

import (
	"fmt"
)

func main() {
	num := 10

	if num%2 == 0 { //checks if number is even
		fmt.Println("The number", num, "is even")
		return
	}
	fmt.Println("The number", num, "is odd")
}

// In the above program, the condition num%2 in line no. 17 finds whether the remainder of dividing num by 2 is zero or not. Since it is 0 in this case, the text The number 10 is even is printed and the program exits.