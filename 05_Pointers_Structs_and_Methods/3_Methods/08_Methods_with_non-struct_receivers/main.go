// package main  

//  func (a int) add(b int) {

//  }

// func main() {

// }

// So far we have defined methods only on struct types. It is also possible to define methods on non-struct types, but there is a catch. To define a method on a type, the definition of the receiver type and the definition of the method should be present in the same package. So far, all the structs and the methods on structs we defined were all located in the same main package and hence they worked.

// In the program above, in line no. 3 we are trying to add a method named add on the built-in type int. This is not allowed since the definition of the method add and the definition of type int is not in the same package. This program will throw compilation error cannot define new methods on non-local type int

// The way to get this working is to create a type alias for the built-in type int and then create a method with this type alias as the receiver.


//==========================================================================//

// The way to get this working is to create a type alias for the built-in type int and then create a method with this type alias as the receiver.

package main  

import "fmt"

type myInt int 

func (a myInt) add(b myInt) myInt {
	return a + b

}

func main() {
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum:", sum)
}