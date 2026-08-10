package main

// When the switch case is inside a for loop, there might be a need to terminate the for loop early. This can be done by labeling the for loop and breaking the for loop using that label inside the switch statement.

// Let’s write a program to generate a random even number.

// We will create an infinite for loop and use a switch case to determine whether the generated random number is even. If it is even, the generated number is printed and the for loop is terminated using its label. The Intn function of the rand package is used to generate non-negative pseudo-random numbers.

import (
	"fmt"
	"math/rand"
)

func main() {
randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Printf("Generated even number %d", i)
			break randloop
		}
	}
}

// In the program above, the for loop is labeled randloop in line no. 9. A random number is generated between 0 and 99 (100 is not included) using the Intn function in line no. 17. If the generated number is even, the loop is broken in line no. 20 using the label.

// This program prints,

// Generated even number 18

// Please note that if the break statement is used without the label, the switch statement will only be broken and the loop will continue running.
// So labeling the loop and using it in the break statement inside the switch is necessary to break the outer for loop.
