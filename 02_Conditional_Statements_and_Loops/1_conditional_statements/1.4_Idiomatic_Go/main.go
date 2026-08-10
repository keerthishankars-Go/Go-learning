package main

// In Go’s philosophy, it is better to avoid unnecessary branches and indentation of code. It is also considered better to return as early as possible. I have provided the program from the previous section below..

// import(
// 	"fmt"
// )

// func main () {
// 	if num := 10 ; num%2 == 0 {
// 		fmt.Println(num, "is even")
// 	} else {
// 		fmt.Println(num, "is odd")
// 	}
// }

//====================================================//

//The idiomatic way of writing the above program in Go’s philosophy is to avoid the else and return from the if if the condition is true.

import (
	"fmt"
)

func main() {
	num := 10

	if num%2 == 0 { // checks if number is even
		fmt.Println(num, "is even")
		return
	}
	fmt.Println(num, "is odd")

}

// In the above program, as soon as we find out the number is even, we return immediately. This avoids the unnecessary else code branch. This is the way things are done in Go.
