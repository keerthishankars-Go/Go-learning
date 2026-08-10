package main 

//We need to stop printing when both i and j are equal i.e when they are equal to 1.

// This is where labels come to our rescue. A label can be used to break from an outer loop. Let’s rewrite the program above using labels.

import (
	"fmt"
)

func main () {
outer: //label
	for i := 0 ; i < 3 ; i++ {
		for j := 1 ; j < 4 ; j++ {
			fmt.Printf("i = %d , j = %d\n", i , j)
			if i == j {
				break outer // breaks loops from outer for loop..
			}
		}
	}
}

// In the program above, we have added a label outer in line no. 8 on the outer for loop and in line no. 13 we break the outer for loop by specifying the label. This program will stop printing when both i and j are equal.