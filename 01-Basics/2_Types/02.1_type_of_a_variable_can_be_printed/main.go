//The type of a variable can be printed using %T format specifier in Printf function
// Go has a unsafe package which has a Sizeof function which returns the size of the variable in bytes.

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a = 90
	b := 55

	fmt.Println("value of a is", a, "value of b", b)

	fmt.Printf("type of a is %T, size of a is %d bytes", a, unsafe.Sizeof(a))

	fmt.Printf("\ntype of b is %T, size of b is %d bytes", b,
		unsafe.Sizeof(b))

	fmt.Printf("\ntype of b is %T, value of b is %d, size of b is %d bytes", b, b, unsafe.Sizeof(b))

}

//For debug logging, this pattern is extremely common:
//fmt.Printf("var=%v type=%T size=%d\n", v, v, unsafe.Sizeof(v))

// The above program outputs the type and size of both variables a and b. %T is the format specifier to print the type and %d is used to print the size.

// We can infer from the above output that a and b are of type int and they have a size of 8 bytes(64 bits). The output will vary if you run the above program on a 32 bit system. In a 32 bit system, a and b occupy 4 bytes(32 bits)
