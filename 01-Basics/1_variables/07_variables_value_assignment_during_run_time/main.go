// Variables can also be assigned values which are computed during run time
package main

import (
	"fmt"
	"math"
)

func main() {

	a, b := 125.3, 465.2

	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)

}

//In the above program math is a package and Min is a function in that package.
//the value of c is calculated at run time and it’s the minimum of a and b.
