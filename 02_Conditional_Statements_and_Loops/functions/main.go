// Functions are central in Go.
package main

import "fmt"

func plus(a int, b int) int { //Here’s a function that takes two ints and returns their sum as an int.
	return a + b //Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
}

func plusPlus(a, b, c int) int { //When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
	return a + b + c
}

func main() {
	res := plus(1, 2) //Call a function just as you’d expect, with name(args).
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}

//There are several other features to Go functions. One is multiple return values
