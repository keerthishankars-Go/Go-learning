package main

// Let’s write one more program where we change the value in b using the pointer.

import (
	"fmt"
)

func main() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b", b)
}

// In line no. 14 of the above program, we increment the value pointed by a by 1 which changes the value of b since a points to b. Hence the value of b becomes 256. The output of the program is

// address of b is 0x1040a124
// value of b is 255
// new value of b is 256