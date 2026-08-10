package main  

// A map is a inbuilt data type in Go which is used to store key-value pairs. A practical use case for a map is for storing the currency codes and the corresponding currency names.

// A map can be created by passing the data type of key and value to the make function. The following is the syntax to create a new map:

// make(map[type of key]type of value)
// currencyCode := make(map[string]string)

import (
	"fmt"
)

func main () {
	currencyCode := make(map[string]string)
	fmt.Println(currencyCode)
}

// The program above creates a map named currencyCode with string key and string value. The above program will print,

// map[]

// Since we have not added any elements to the map, it’s empty.