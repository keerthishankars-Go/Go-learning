package main

import (
	"fmt"
)

func main() {
	hour := 16

	switch {
	case hour >= 6 && hour < 12:
		fmt.Println("It is a morning shift")
	case hour >= 12 && hour < 17:
		fmt.Println("It is an afternoon shift")
	case hour >= 17 && hour < 21:
		fmt.Println("It is an evening shift")
	case (hour >= 21 && hour <= 12) || (hour >= 0 && hour <= 6):
		fmt.Println("It is a night shift")
	default:
		fmt.Println("It is an invalid hour")
	}
}

// In the above program, the expression is absent in switch and hence it is considered as true and each of the cases is evaluated. hour >= 12 && hour < 17 in line no. 13 is true and the program prints

// It's the afternoon shift.
// This type of switch can be considered as an alternative to multiple if else clauses.

// Why this is better than if-else if
// Equivalent if-else version

// if hour >= 6 && hour < 12 {
// 	fmt.Println("It's the morning shift.")
// } else if hour >= 12 && hour < 17 {
// 	fmt.Println("It's the afternoon shift.")
// } else if hour >= 17 && hour < 21 {
// 	fmt.Println("It's the evening shift.")
// } else if (hour >= 21 && hour <= 24) || (hour >= 0 && hour < 6) {
// 	fmt.Println("It's the night shift.")
// } else {
// 	fmt.Println("Invalid hour.")
// }
