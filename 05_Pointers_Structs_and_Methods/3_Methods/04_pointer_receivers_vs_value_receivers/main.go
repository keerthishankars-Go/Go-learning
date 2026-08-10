package main

// It is possible to create methods with pointer receivers. The difference between value and pointer receiver is, changes made inside a method with a pointer receiver is visible to the caller whereas this is not the case in value receiver.

import (
	"fmt"
)

type Employee struct {
	name string
	age  int
}

// Method with value receiver..

func (e Employee) changeName(newName string) {
	e.name = newName
}

// Method with pointer receiver..

func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Keerthi",
		age:  27,
	}
	fmt.Printf("Employee name before change is %s\n", e.name)
	e.changeName("Keerthishankar")
	fmt.Printf("Employee name after change is %s\n", e.name)

	fmt.Printf("Employee age before change is %d\n", e.age)
	(&e).changeAge(28)
	//e.changeAge(29)
	fmt.Printf("Employee age after change is %d", e.age)
}

// In the program above, the changeName method has a value receiver (e Employee) whereas the changeAge method has a pointer receiver (e *Employee). Changes made to Employee struct’s name field inside changeName will not be visible to the caller and hence the program prints the same name before and after the method e.changeName("Keerthishankar") is called in line no. 32. Since changeAge method has a pointer receiver (e *Employee), changes made to age field after the method call (&e).changeAge(28) will be visible to the caller.

// In line no. 36 of the program above, we use (&e).changeAge(51) to call the changeAge method. Since changeAge has a pointer receiver, we have used (&e) to call the method. This is not needed and the language gives us the option to just use e.changeAge(29). e.changeAge(29) will be interpreted as (&e).changeAge(28) by the language.
