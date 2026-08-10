package main

// As we already know arrays are restricted to fixed length and their length cannot be increased. Slices are dynamic and new elements can be appended to the slice using append function.
// The definition of append function is func append(s []T, x ...T) []T.

// x …T in the function definition means that the function accepts variable number of arguments for the parameter x. These type of functions are called variadic functions.

import (
	"fmt"
)

func main() {
	cars := []string{"Honda", "Tata", "BMW"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity:", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // //capacity of cars is doubled to 6

}

// In the above program, the capacity of cars is 3 initially. We append a new element to cars and assign the slice returned by append(cars, "Toyota") to cars again. Now the capacity of cars is doubled and becomes 6.
