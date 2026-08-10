// The value of a constant should be known at compile time. Hence it cannot be assigned to a value returned by a function call since the function call takes place at run time.

package main   

import (
	"fmt"
	"math"
)

func main () {
	var a = math.Sqrt(4) // allowed
	fmt.Println(a)

	const b = math.Sqrt(4)  // not allowed
	fmt.Println(b)
}

// In the above program, a is a variable and hence it can be assigned to the result of the function math.Sqrt(4) 

// b is a constant and the value of b needs to be known at compile time. The function math.Sqrt(4) will be evaluated only during run time and hence const b = math.Sqrt(4) fails to compile with error
