// int8, int16, int32, int64, int are the signed integer data types available in Go
package main

import "fmt"

func main() {
	var a int = 89

	b := 95

	fmt.Println("value of a is:", a, "value of b:", b)
}

// In the above program a is of type int and the type of b is inferred from the value assigned to it (95)
