package main  

//We have only 5 fingers in our hands. What will happen if we input an incorrect finger number? This is where the default case comes into the picture. The default case will be executed when none of the other cases match.

import (
	"fmt"
)

func main () {
	switch finger := 8; finger {
	case 1 :
		fmt.Println("thumb")
	case 2 :
		fmt.Println("index")
	case 3 :
		fmt.Println("middle")
	case 4 :
		fmt.Println("ring")
	case 5 :
		fmt.Println("pinky")
	default :
		fmt.Println("incorrect finger number")
	}
}

// In the above program finger is 8 and it does not match any of the cases and hence incorrect finger number in the default case is printed. It’s not necessary that default should be the last case in a switch statement. It can be present anywhere in the switch.

// notice a small change in the declaration of finger. It is declared in the switch itself. A switch can include an optional statement that is executed before the expression is evaluated. In line no. 8, finger is first declared and then used in the expression. The scope of finger in this case is limited to the switch block.