package main

import "fmt"

func rectProps(length, width float64) (area, perimeter float64) {
	//If a return value is named, it can be considered as being declared as a variable in the first line of the function.

	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value

}

func main() {
	area, perimeter := rectProps(23.12, 45.23)

	fmt.Printf("Area is %.2f and perimeter is %.2f", area, perimeter)

}

//area and perimeter are the named return values in the above function. Note that the return statement in the function does not explicitly return any value. Since area and perimeter are specified in the function declaration as return values, they are automatically returned from the function when a return statement is encountered.
