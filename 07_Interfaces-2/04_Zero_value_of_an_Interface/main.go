package main

//The zero value of a interface is nil. A nil interface has both its underlying value and as well as concrete type as nil.

import "fmt"

type Describer interface {
	Describe()
}

func main() {
	var d1 Describer
	if d1 == nil {
		fmt.Printf("d1 is nil and it's type is %T and value is %v", d1, d1)
	}
}

//d1 in the above program is nil and this program will output

// d1 is nil and has type <nil> value <nil>

//============================================================================//

//If we try to call a method on the nil interface, the program will panic since the nil interface neither has a underlying value nor a concrete type.

// package main

// type Describer interface {
// 	Describe()
// }

// func main() {
// 	var d1 Describer
// 	d1.Describe()
// }

// Since d1 in the program above is nil, this program will panic with runtime error

// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4664b0]

// goroutine 1 [running]:
// main.main()
// 	/tmp/sandbox2797051632/prog.go:9 +0x10
