//Numeric constants include integers, floats and complex constants.

package main

import (
	"fmt"
)

func main() {
	const c = 5
	var intVar int = c
	var int32Var int32 = c
	var float64Var float64 = c
	var complex64Var complex64 = c
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}

// In the program above, the value of c is 5 and the syntax of c is generic. It can represent a float, integer or even a complex number with no imaginary part. Hence it is possible to be assigned to any compatible type. The default type of these kinds of constants can be thought of as being generated based on the context where they are used. var intVar int = c requires c to be int so it becomes an int constant. var complex64Var complex64 = c requires c to be a complex number and hence it becomes a complex constant.
