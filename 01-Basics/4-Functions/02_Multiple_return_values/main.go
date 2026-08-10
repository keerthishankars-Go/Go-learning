// It is possible to return multiple values from a function.
// Let’s write a function rectProps which takes the length and width of a rectangle and returns both the area and perimeter of the rectangle. The area of the rectangle is the product of length and width and the perimeter is twice the sum of the length and width.

package main

import "fmt"

func rectProps(length, width float64) (float64, float64) {

	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	//Calling a multi-return function:
	area, perimeter := rectProps(16.4, 23.47)
	fmt.Printf("Area %f, perimeter %f\n", area, perimeter)

	// Better readable version:
	fmt.Printf("Area %.2f and Perimeter %.2f", area, perimeter)

}

// rectProps runs

// Returns (areaValue, perimeterValue)

// First value → assigned to area

// Second value → assigned to perimeter

// 📌 Rule

// Left side variables must match number and order of returned values.

// ========================================
// ========================================

// Ignoring return values (VERY IMPORTANT)

// If you don’t want one value:

// area, _ := rectProps(10.8, 5.6)

// _ (blank identifier)

// Tells Go: “I don’t care about this value”

// Prevents unused variable error

// 📌 Rule

// Use _ to ignore unwanted return values.