package main 

// The zero value of a slice type is nil. 
// A nil slice has length and capacity 0. It is possible to append values to a nil slice using the append function

import (
	"fmt"
)

func main () {
	var names []string // zero value of a slice is nil..
	//if names == nil {
		fmt.Println("slice is nil going to be append")
		names = append(names, "Keerthi", "Sunil", "Chethan")
		fmt.Println("names contents:", names)
	//}
}