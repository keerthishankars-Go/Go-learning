package main

// The range form of the for loop is used to iterate over all elements of a map.

import (
	"fmt"
)

func main() {
	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}
	for code, name := range currencyCode {
		fmt.Printf("Currency name for currency code %s is %s\n", code, name)
	}
}

// One important fact to note is the order of the retrieval of values from a map when using for range is not guaranteed to be the same for each execution of the program. It is also not the same as the order in which the elements were added to the map..
