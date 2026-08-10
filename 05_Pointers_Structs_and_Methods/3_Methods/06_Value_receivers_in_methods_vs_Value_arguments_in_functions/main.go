package main

// When a function has a value argument, it will accept only a value argument.

//When a method has a value receiver, it will accept both pointer and value receivers.

import "fmt"

type rectangle struct {
	length int
	width int

}

func area(r rectangle) {
	fmt.Printf("Area function result: %d\n", (r.length * r.width))

}

func (r rectangle) area() {
		fmt.Printf("Area method result: %d", (r.length * r.width))
}

func main() {
		r := rectangle{
			length: 5,
			width: 4,

		}

		area(r)
		r.area()

		p := &r

		//area(p)
		p.area()

}

// function func area(r rectangle) in line no.12 accepts a value argument and method func (r rectangle) area() in line no. 16 accepts a value receiver.

// In line no. 25, we call the area function with a value argument area(r) and it will work. Similarly, we call the area method r.area() using a value receiver and this will work too.

// We create a pointer p to r in line no. 28. If we try to pass this pointer to the function area which accepts only a value, the compiler will complain. I have commented line no. 33 which does this. If you uncomment this line, then the compiler will throw error *compilation error, cannot use p (type rectangle) as type rectangle in argument to area. This works as expected.

// Now comes the tricky part, line no. 35 of the code p.area() calls the method area which accepts only a value receiver using the pointer receiver p. This is perfectly valid. The reason is that the line p.area(), for convenience will be interpreted by Go as (*p).area() since area has a value receiver.