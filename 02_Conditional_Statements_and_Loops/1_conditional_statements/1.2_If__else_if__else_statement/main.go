// The if statement also has optional else if and else components. The syntax for the same is provided below:

// if condition1 {
// ...
// } else if condition2 {
// ...
// } else {
// ...
// }

package main 

// Let’s write a bus ticket pricing program using else if. The program must satisfy the following requirements.

// If the age of the passenger is less than 5 years, the ticket is free.
// If the age of the passenger is between 5 and 22 years, then the ticket is $10.
// If the age of the passenger is above 22 years, then the ticket price is $15.

import (
	"fmt"
)

func main () {
	age := 10
	ticketPrice := 0

	if age < 5 {
		ticketPrice = 0
	} else if age >= 5 && age <= 22 {
		ticketPrice = 10 
	} else {
		ticketPrice = 15
	}
	fmt.Printf("Ticket price is $%d", ticketPrice)
}

// In the above program, the age of the passenger is set to 10. The condition in line no. 29 is true and hence the program will print

//Ticket price is $10

