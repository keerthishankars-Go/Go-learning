package main  

import "fmt"

type person struct {
	firstName string 
	lastName string 
}

func (p person) fullName() {
	fmt.Printf("%s, %s", p.firstName, p.lastName)

}

func main() {
	p := person{
		firstName: "Keerthi",
		lastName: "shankar",
	}

	defer p.fullName()

	fmt.Printf("Welcome ")
}

// Defer is not restricted only to functions. It is perfectly legal to defer a method call too.

// In the above program we have deferred a method call in line no. 21. The rest of the program is self explanatory. This program outputs,

// Welcome Keerthi, shankar