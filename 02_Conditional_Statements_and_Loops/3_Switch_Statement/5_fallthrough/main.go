package main

// In Go, the control comes out of the switch statement immediately after a case is executed. A fallthrough statement is used to transfer control to the first statement of the case that is present immediately after the case which has been executed.

// Our program will check whether the input number is less than 50, 100, or 200. For instance, if we input 75, the program will print that 75 is less than both 100 and 200. We will achieve this using fallthrough.

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {
	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)

	}
}

// Switch and case expressions need not be only constants. They can be evaluated at runtime too. In the program above num is initialized to the return value of the function number() in line no. 17.
// The control comes inside the switch and the cases are evaluated. case num < 100: in line no. 21 is true and the program prints 75 is lesser than 100.

//  The next statement is fallthrough. When fallthrough is encountered the control moves to the first statement of the next case and also prints 75 is lesser than 200. The output of the program is

// 75 is lesser than 100
// 75 is lesser than 200

// fallthrough should be the last statement in a case. If it is present somewhere in the middle, the compiler will complain that fallthrough statement out of place.
