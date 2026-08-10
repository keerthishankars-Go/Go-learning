// Numeric constants are free to be mixed and matched in expressions and a type is needed only when they are assigned to variables or used in any place in code which demands a type.

package main

import "fmt"

func main() {

	var a = 5.9 / 8

	fmt.Printf("a's type is %T and value is %v", a, a)
}

// In the program above, 5.9 is a float by syntax and 8 is an integer by syntax. Still, 5.9/8 is allowed as both are numeric constants. The result of the division is 0.7375 is a float and hence variable a is of type float .
