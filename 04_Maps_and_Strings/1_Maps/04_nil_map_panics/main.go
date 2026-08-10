package main  

// The zero value of a map is nil. If you try to add elements to a nil map, a run-time panic will occur. Hence the map has to be initialized before adding elements.

func main () {
	// var currencyCode map[string]string 
	// currencyCode["USD"] = "US Dollar"
}

// In the above program, currencyCode is nil and we are trying to add a new key to a nil map. The program will panic with error

// panic: assignment to entry in nil map