package main

// Do not pass a pointer to an array as an argument to a function. Use slice instead.

import (
	"fmt"
)

// func modify(arr *[3]int) {
// 	// (*arr)[0] = 90 or
// 	arr[0] = 90
// }

// func main() {
//     a := [3]int{89, 90, 91}
//     modify(&a)
//     fmt.Println(a)
// }

// In line no. 16 of the above program, we are passing the address of the array a to the modify function. In line no.10 in the modify function we are dereferencing arr and assigning 90 to the first element of the array. This program outputs [90 90 91]

// a[x] is shorthand for (*a)[x]. So (*arr)[0] in the above program can be replaced by arr[0].

// Although this way of passing a pointer to an array as an argument to a function and making modification to it works, it is not the idiomatic way of achieving this in Go. We have slices for this.

func modify(sls []int) {
	sls[0] = 90

}

func main() {
	a := [3]int{89, 90, 91}
	modify(a[:])
	fmt.Println(a)

}

// In line no.33 of the program above, we pass a slice to the modify function. The first element of the slice is changed to 90 inside the modify function. This program also outputs [90 90 91]. So forget about passing pointers to arrays around and use slices instead
