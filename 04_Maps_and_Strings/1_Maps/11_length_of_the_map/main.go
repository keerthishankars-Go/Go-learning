package main

// Length of the map can be determined using the len function.

import (
	"fmt"
)

func main() {
	currencyCode := map[string]string{
		"INR": "Indian Rupee",
		"USD": "US Dollar",
		"EUR": "Euro",
	}
	fmt.Println("The length of the map is", len(currencyCode))
}

// len(currencyCode) in the above program returns the length of the map. The above program prints:

// length is 3