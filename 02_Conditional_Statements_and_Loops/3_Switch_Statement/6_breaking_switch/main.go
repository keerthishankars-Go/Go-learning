package main

// The break statement can be used to terminate a switch early before it completes.

// Let’s add a condition that if num is less than 0 then the switch should terminate.

import (
	"fmt"
)

func main() {
	switch num := -5; {
	case num < 50:
		if num < 0 {
			break
		}
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)

	}
}

// In the above program num is -5. When the control reaches the if statement in line no. 10, the condition is satisfied since num < 0. The break statement terminates the switch before it completes and the program doesn’t print anything..
