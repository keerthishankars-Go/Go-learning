package main  

// The semicolons in the for loop of the above program can also be omitted. This format can be considered as an alternative for while loop. The above program can be rewritten as,

import (
	"fmt"
)

func main () {
	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present. This is similar to while loop.		
		fmt.Printf("%d ",i)
		i += 2

	}
}