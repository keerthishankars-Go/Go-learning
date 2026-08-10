package main 

// It is also possible to initialize a map during the declaration itself.

import (
	"fmt"
)

func main () {
	currencyCode := map[string]string {
			"USD" : "US Dollar",
			"GBP": "Pound Sterling",
			"EUR": "Euro",
	}
	currencyCode["INR"] = "Indian Rupee"
	fmt.Println("currencyCode map contents:", currencyCode)
}

// The above program declares currencyCode map and adds 3 items to it during the declaration itself. Later one more element with key INR is added. 

// It’s not necessary that only string types should be keys. All comparable types such as boolean, integer, float, complex, string, … can also be keys. Even user-defined types such as structs can be keys.