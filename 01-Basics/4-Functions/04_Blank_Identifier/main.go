// _ is known as the blank identifier in Go. It can be used in place of any value of any type.
package main

import "fmt"

// The rectProps function returns the area and perimeter of the rectangle. What if we only need the area and want to discard the perimeter. This is where _ is of use.

// The program below uses only the area returned from the rectProps function.

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2

	return area, perimeter

}

func main() {
	area, _ := rectProps(33.65, 29.56) // perimeter is discarded

	fmt.Printf("Area is %.2f", area)
}

//we use only the area and the _ identifier is used to discard the perimeter.
