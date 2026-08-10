package main

// Let’s start with a simple example which will take a finger number as input and outputs the name of that finger :) . For example, 1 is thumb, 2 is index, and so on.

import (
	"fmt"
)

func main() {
	finger := 4
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("index")
	case 3:
		fmt.Println("middle")
	case 4:
		fmt.Println("ring")
	case 5:
		fmt.Println("pinky")
	}

}

// In the above program switch finger in line no. 12, compares the value of finger with each of the case statements. The cases are evaluated from top to bottom and the first case which matches the expression is executed. In this case, finger has a value of 4 and hence:

// Finger 4 is Ring

// Duplicate cases are not allowed..

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	finger := 4
// 	fmt.Printf("Finger %d is ", finger)
// 	switch finger {
// 	case 1:
// 		fmt.Println("Thumb")
// 	case 2:
// 		fmt.Println("Index")
// 	case 3:
// 		fmt.Println("Middle")
// 	case 4:
// 		fmt.Println("Ring")
// 	case 4: //duplicate case
// 		fmt.Println("Another Ring")
// 	case 5:
// 		fmt.Println("Pinky")

// 	}
// }

