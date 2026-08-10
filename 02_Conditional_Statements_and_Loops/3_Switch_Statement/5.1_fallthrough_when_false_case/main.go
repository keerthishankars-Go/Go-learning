// Fallthrough happens even when the case evaluates to false

// There is a subtlety to be considered when using fallthrough. Fallthrough will happen even when the case evaluates to false.

package main

import (
	"fmt"
)

func main() {
	switch num := 25; { 
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)		
	}

}

// In the above program, num is 25 which is less than 50 and hence the case in line no. 9 evaluates to true. A fallthrough is present in line no. 11. The next case case num > 100: in line no. 12 is false since num < 100. But fallthrough doesn’t consider this. Fallthrough will happen even though the case evaluates to false.

// The program above will print:

// 25 is lesser than 50
// 25 is greater than 100

// So be sure that you understand what you are doing when using fallthrough.

// One more thing is fallthrough cannot be used in the last case of a switch since there are no more cases to fallthrough. If fallthrough is present in the last case, it will result in the following compilation error cannot fallthrough final case in switch.

