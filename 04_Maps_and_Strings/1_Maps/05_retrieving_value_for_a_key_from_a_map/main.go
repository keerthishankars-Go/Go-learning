package main  

// we have added some elements to the map, let’s learn how to retrieve them. 

// map[key] is the syntax to retrieve elements of a map.

import (
	"fmt"
)

func main () {
	currencyCode := map[string]string {
			"USD": "US Dollar",
			"GBP": "Pound Sterling",
			"EUR": "Euro",

	}
	currency := "USD"
	currencyName := currencyCode[currency]
	fmt.Println("Currency Name for currency code", currency, "is", currencyName)
			


}