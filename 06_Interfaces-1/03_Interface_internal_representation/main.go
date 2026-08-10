package main 

// An interface can be thought of as being represented internally by a tuple (type, value). type is the underlying concrete type that implements the interface and value holds the value of the concrete type.

import "fmt"

type Worker interface{
	Work()
}

type Person struct {
	name string 
}

func (p Person) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T and value %v\n", w, w)
}

func main() {
	p := Person{
		name: "Keerthi",
	}

	var w Worker = p

	describe(w)

	w.Work()
}

//Worker interface has one method Work() and Person struct type implements that interface. In line no. 27, we assign the variable p of type Person to w which is of type Worker. Now the concrete type of w is Person and it contains a Person with name field Naveen. The describe function in line no.19 prints the value and concrete type of the interface. This program outputs

// Interface type main.Person value {Naveen}
// Naveen is working