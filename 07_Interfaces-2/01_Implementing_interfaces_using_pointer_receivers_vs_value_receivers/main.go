package main

// All the interfaces we discussed in part 1 were implemented using value receivers. It is also possible to implement interfaces using pointer receivers.

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("%s is in %s", a.state, a.country)

}

func main() {
	var d1 Describer
	p1 := Person{"Keerthi", 27}
	d1 = p1
	d1.Describe()

	p2 := Person{"sunil", 30}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Karnataka", "India"}
	//d2 = a

	/*compilation error if the following line is
	uncommented
	./prog.go:45:7: cannot use a (variable of type Address)
	as Describer value in assignment: Address does
	not implement Describer
	(method Describe has pointer receiver)
	*/

	d2 = &a //This works since Describer interface is implemented by Address pointer in line 22

	d2.Describe()
}

// In the program above, the Person struct implements the Describer interface using value receiver in line no. 13.

// As we have already learnt during our discussion about methods, methods with value receivers accept both pointer and value receivers. It is legal to call a value method on anything which is a value or whose value can be dereferenced.

// p1 is a value of type Person and it is assigned to d1 in line no. 29. Person implements the Describer interface and hence line no. 30 will print Sam is 25 years old.

// Similarly &p2 is assigned to d1 in line no. 32 and hence line no. 33 will print James is 32 years old. Awesome :).

// The Address struct implements the Describer interface using pointer receiver in line no. 22.

// If line. no 45 of the program above is uncommented, we will get the compilation error ./prog.go:45:7: cannot use a (variable of type Address) as Describer value in assignment: Address does not implement Describer (method Describe has pointer receiver). This is because, the Describer interface is implemented using a pointer receiver in line 22 and we are trying to assign a which is a value type and it has not implemented the Describer interface. This will definitely surprise you since we learnt earlier that methods with pointer receivers will accept both pointer and value receivers. Then why isn’t the code in line no. 45 working.

// The reason is that it is legal to call a pointer-valued method on anything that is already a pointer or whose address can be taken. The concrete value stored in an interface is not addressable and hence it is not possible for the compiler to automatically take the address of a in line no. 45 and hence this code fails.

// Line no. 47 works because we are assigning the address of a &a to d2.

// The rest of the program is self explanatory. This program will print,

// Sam is 25 years old
// James is 32 years old
// Washington is in USA