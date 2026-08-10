package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55

}

func main() {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is", a)
}

// In the above program, in line no. 16 we are passing the pointer variable b which holds the address of a to the function change. Inside change function, the value of a is changed using dereference in line no 8. This program outputs,

// value of a before function call is 58
// value of a after function call is 55
