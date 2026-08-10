package main

// delete(map, key) is the syntax to delete key from a map. The delete function does not return any value.

import (
	"fmt"
)

func main() {
	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}
	fmt.Println("Map before deletion", currencyCode)
	delete(currencyCode, "EUR")
	// delete(currencyCode, "INR")
	fmt.Println("Map after deletion", currencyCode)
}

// Even if we try to delete a key that is not present in the map, there will be no runtime error.
