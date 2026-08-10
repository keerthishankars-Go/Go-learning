package main

import (
	"errors"
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{
			err:    "radius is negative",
			radius: radius,
		}

	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0

	area, err := circleArea(radius)
	if err != nil {
		var areaError *areaError
		if errors.As(err, &areaError) {
			fmt.Printf("Area calculation failed, radius %0.2f is less than zero", areaError.radius)
			return
		}
		fmt.Println(err)
		return

	}
	fmt.Printf("Area of rectangle %0.2f", area)
}

// In the program above, circleArea in line no. 18 is used to calculate the area of the circle. This function first checks if the radius is less than zero, if so it creates a value of type areaError using the radius responsible for the error and the corresponding error message and then returns the address of it in line no. 20 along with 0 as area. Thus we have provided more information about the error, in this case the radius which caused the error using the fields of a custom error struct.

// If the radius is not negative, this function calculates and returns the area along with a nil error in line no. 25.

// In line no. 30 of the main function, we are trying to find the area of a circle with radius -20. Since the radius is less than zero, an error will be returned.

// We check whether the error is not nil in line no. 31 and in line no. 33 line we try to convert it to type *areaError. If the error is of type *areaError, we get the radius which caused the error in line no. 34 using areaError.radius, print a custom error message and return from the program.

// If the error is not of type *areaError, we simply print the error in line no. 37 and return. If there is no error, the area will be printed in line no.40.

// The program will print,

// Area calculation failed, radius -20.00 is less than zero
