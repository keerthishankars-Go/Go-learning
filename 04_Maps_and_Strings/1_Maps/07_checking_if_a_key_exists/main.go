package main

// In the above section we learned that when a key is not present, the zero value of the type will be returned. This doesn’t help when we want to find out whether the key actually exists in the map.

// For example, we want to know whether a currency code key is present in the currencyCode map. The following syntax is used to find out whether a particular key is present in a map.

// value, ok := map[key]

// ok in the above line of code will be true when the key is present and the value for the key is present in the variable value. If the key is not present, ok will be false and the zero value is returned for value.

import (
	"fmt"
)

func main() {
	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
		//"INR": "Indian Rupee",
	}
	cyCode := "INR"
	if currencyName, ok := currencyCode[cyCode]; ok {
		fmt.Println("Currency name for the currency code", cyCode, "is", currencyName)
		return
	}
	fmt.Println("Currency name for the currency code", cyCode, "not found")
}

// In the above program, in line no. 23, ok will be false since INR key is not present. Hence the program will print,

// Currency name for currency code INR not found
