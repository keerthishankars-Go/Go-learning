// If else statement:

// The if statement has an optional else construct which will be executed if the condition in the if statement evaluates to false.

package main

// if condition {
// } else {
// }

// In the above snippet, if condition evaluates to false, then the block of code between else { and } will be executed.

import (
	"fmt"
)

func main() {
	num := 11
	if num%2 == 2 {
		fmt.Println("The given number", num, "is even")
	} else {
		fmt.Println("The number", num, "is odd")
	}
}

// In the above code, instead of returning if the condition is true as we did in the previous section, we create an else statement that will be executed if the condition is false. In this case, since 11 is odd, the if condition is false and the lines of code within the else statement is executed.
