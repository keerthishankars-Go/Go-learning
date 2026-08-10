package main

// What will happen if an element is not present? The map will return the zero value of the type of that element. In the case of currencyCode map, if we try to access an item which is not present, the zero value of string, “"(the empty string) is returned.

import (
	"fmt"
)

func main() {
	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}
	fmt.Println("Currency name for the currency code INR is", currencyCode["INR"])

}

// The above program returns empty string as the currency name for INR. There will be no runtime error when we try to retrieve the value for a key that is not present in the map.
