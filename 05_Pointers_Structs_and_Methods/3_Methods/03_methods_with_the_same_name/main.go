package main  

// Methods with the same name can be defined on different types whereas functions with the same names are not allowed. 
// 
// Let’s assume that we have a Square and Circle structure. It’s possible to define a method named Area on both Square and Circle. This is done in the program below.

// It’s possible to define a method named Area on both Square and Circle.

import (
	"fmt"
	"math"
)

type Rectangle struct {
	length  int 
	width int 
}

type Circle struct {
	radius float64 
}

func (r Rectangle) Area() int {
	return r.length * r.width 

}

func (c Circle) Area() float64 {
	return math.Pi * c.radius *c.radius
}

func main () {
	r := Rectangle{
		length : 4,
		width :6,
	}
	fmt.Printf("Area of rectangle is %d\n", r.Area())

	c := Circle {
		radius : 11.2,
	}
	fmt.Printf("The area of circle is %.2f", c.Area())
}

// The above property of methods is used to implement interfaces.