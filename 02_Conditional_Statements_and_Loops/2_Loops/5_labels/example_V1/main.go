//Labels can be used to break the outer loop from inside the inner for loop. Let’s understand what I mean by using a simple example.

package main  

import (
	"fmt"
)

func main () {
	for i := 0 ; i < 3 ; i++ {
		for j := 1; j < 4 ; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}