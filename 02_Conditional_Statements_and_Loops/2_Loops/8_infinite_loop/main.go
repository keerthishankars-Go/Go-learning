// The syntax for creating an infinite loop is,

// for {
// }

// The following program will keep printing Hello World continuously without terminating.

package main  

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("Hello world")
	}
}