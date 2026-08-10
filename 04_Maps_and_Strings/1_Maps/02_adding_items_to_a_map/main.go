package main 

// The syntax for adding new items to a map is the same as that of arrays. The program below adds some new currency codes and currency names to the currencyCode map.

import (
	"fmt"
)

func main () {
	currencyCode := make(map[string]string)
	currencyCode["USA"] = "US Dollar"
	currencyCode["GBP"] = "Pound Sterling"
	currencyCode["EUR"] = "Euro"
	currencyCode["INR"] = "Indian Rupee"
	fmt.Println("currencyCode map contents:", currencyCode)
}

// As you might have recognized from the above output, the order of the retrieval of values from a map is not guaranteed to be the same as the order in which the elements were added to the map.