// There is one more variant of if which includes an optional shorthand assignment statement that is executed before the condition is evaluated. Its syntax is

// if assignment-statement; condition {
// }

// In the above snippet, assignment-statement is first executed before the condition is evaluated.

package main  

import (
	"fmt"
)

func main () {
	ticketPrice := 0 

	if age := 10 ; age <= 5 {
		ticketPrice = 0 
	} else if age >= 15 && age <= 22 {
		ticketPrice = 15 
	} else {
		ticketPrice = 20 
	}

	fmt.Println("The ticket price is", ticketPrice)
}

// In the above program age is initialized in the if statement in line no.17. age can be accessed from only within the if construct. i.e. the scope of age is limited to the if, else if and else blocks. If we try to access age outside the if, else if or else blocks, the compiler will complain.
// This syntax often comes in handy when we declare a variable just for the purpose of if else construct. Using this syntax in such cases ensures that the scope of the variable is only within the if statement.