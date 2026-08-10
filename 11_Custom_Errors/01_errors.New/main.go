package main  

import (
	"errors"
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil 
}

func main() {
	radius := -3.15
	area, err := circleArea(radius)

	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Printf("Area of circle is %0.2f", area)
}

// In the program above, we check whether the radius is less than zero in line no. 10. If so we return zero for the area along with the corresponding error message. If the radius is greater than 0, then the area is calculated and nil is returned as the error in line no. 13.

// In the main function, we check whether the error is not nil in line no. 19. If it’s not nil, we print the error and return, else the area of the circle is printed.

// In this program the radius is less than zero and hence it will print,

// Area calculation failed, radius is less than zero